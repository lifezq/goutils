// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package encode

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
	"unicode/utf16"
)

func ForceDecodeXLS(r io.Reader) ([][]string, string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, "", err
	}
	wb, err := extractWorkbookStream(data)
	if err != nil {
		return nil, "", err
	}
	rows, enc, err := parseBIFF8Workbook(wb)
	if err != nil {
		return nil, "", err
	}
	return rows, enc, nil
}

const (
	cfbfMagic            = 0xE011CFD0
	oleFree              = 0xFFFFFFFF
	oleEndOfChain        = 0xFFFFFFFE
	oleFatSect           = 0xFFFFFFFD
	oleDifSect           = 0xFFFFFFFC
	dirEntSize           = 128
	headerSize           = 512
	defaultSectorShift   = 9
	defaultMiniSectorLog = 6
	miniStreamCutoff     = 4096
)

type oleHeader struct {
	sectorSize   int
	miniSize     int
	fatCount     uint32
	dirStart     uint32
	miniFatStart uint32
	miniFatCount uint32
	difatStart   uint32
	difatCount   uint32
	difatHeader  [109]uint32
}

func readHeader(b []byte) (*oleHeader, error) {
	if len(b) < headerSize {
		return nil, fmt.Errorf("invalid CFBF header size")
	}
	sig := binary.LittleEndian.Uint32(b[0:4])
	if sig != cfbfMagic {
		return nil, fmt.Errorf("not an XLS OLE container")
	}
	if binary.LittleEndian.Uint16(b[0x1C:0x1E]) != 0xFFFE {
		return nil, fmt.Errorf("invalid byte order")
	}
	sectorShift := binary.LittleEndian.Uint16(b[0x1E:0x20])
	miniShift := binary.LittleEndian.Uint16(b[0x20:0x22])
	h := &oleHeader{
		sectorSize:   int(1) << sectorShift,
		miniSize:     int(1) << miniShift,
		fatCount:     binary.LittleEndian.Uint32(b[0x2C:0x30]),
		dirStart:     binary.LittleEndian.Uint32(b[0x30:0x34]),
		miniFatStart: binary.LittleEndian.Uint32(b[0x3C:0x40]),
		miniFatCount: binary.LittleEndian.Uint32(b[0x40:0x44]),
		difatStart:   binary.LittleEndian.Uint32(b[0x44:0x48]),
		difatCount:   binary.LittleEndian.Uint32(b[0x48:0x4C]),
	}
	for i := 0; i < 109; i++ {
		h.difatHeader[i] = binary.LittleEndian.Uint32(b[0x4C+4*i : 0x4C+4*(i+1)])
	}
	return h, nil
}

func sectorOffset(h *oleHeader, sid uint32) int {
	return headerSize + int(sid)*h.sectorSize
}

func readFATSectorIDs(b []byte, h *oleHeader) ([]uint32, error) {
	var fats []uint32
	for _, v := range h.difatHeader {
		if v == oleFree {
			continue
		}
		fats = append(fats, v)
	}
	if uint32(len(fats)) < h.fatCount {
		if len(fats) == 0 {
			return nil, fmt.Errorf("no FAT sector listed")
		}
	}
	return fats, nil
}

func readFAT(b []byte, h *oleHeader) ([]uint32, error) {
	fatSectors, err := readFATSectorIDs(b, h)
	if err != nil {
		return nil, err
	}
	entriesPer := h.sectorSize / 4
	var fat []uint32
	for _, sid := range fatSectors {
		off := sectorOffset(h, sid)
		if off+h.sectorSize > len(b) {
			return nil, fmt.Errorf("FAT sector out of range")
		}
		for i := 0; i < entriesPer; i++ {
			val := binary.LittleEndian.Uint32(b[off+4*i : off+4*(i+1)])
			fat = append(fat, val)
		}
	}
	return fat, nil
}

func walkChain(fat []uint32, start uint32) ([]uint32, error) {
	var chain []uint32
	s := start
	limit := len(fat)
	for {
		if int(s) < 0 || int(s) >= limit {
			return nil, fmt.Errorf("sector chain out of range")
		}
		chain = append(chain, s)
		next := fat[s]
		if next == oleEndOfChain {
			break
		}
		if next == oleFree || next == oleFatSect || next == oleDifSect {
			return nil, fmt.Errorf("invalid sector chain")
		}
		s = next
		if len(chain) > 1_000_000 {
			return nil, fmt.Errorf("sector chain too long")
		}
	}
	return chain, nil
}

type dirEntry struct {
	name        string
	entryType   byte
	startSector uint32
	streamSize  uint64
}

func readDirectory(b []byte, h *oleHeader, fat []uint32) ([]dirEntry, error) {
	chain, err := walkChain(fat, h.dirStart)
	if err != nil {
		return nil, err
	}
	var dirBytes []byte
	for _, sid := range chain {
		off := sectorOffset(h, sid)
		end := off + h.sectorSize
		if end > len(b) {
			return nil, fmt.Errorf("directory sector out of range")
		}
		dirBytes = append(dirBytes, b[off:end]...)
	}
	if len(dirBytes)%dirEntSize != 0 {
		return nil, fmt.Errorf("invalid directory size")
	}
	var entries []dirEntry
	for i := 0; i < len(dirBytes); i += dirEntSize {
		ent := dirBytes[i : i+dirEntSize]
		nlen := int(binary.LittleEndian.Uint16(ent[0x40:0x42]))
		// name length in bytes including null terminator
		name := ""
		if nlen >= 2 && 0x00+nlen-2 <= len(ent) {
			raw := ent[0x00 : 0x00+nlen-2]
			u16 := make([]uint16, len(raw)/2)
			for j := 0; j < len(u16); j++ {
				u16[j] = binary.LittleEndian.Uint16(raw[2*j : 2*j+2])
			}
			runes := utf16.Decode(u16)
			name = string(runes)
		}
		entryType := ent[0x42]
		start := binary.LittleEndian.Uint32(ent[0x74:0x78])
		size := binary.LittleEndian.Uint64(ent[0x78:0x80])
		entries = append(entries, dirEntry{
			name:        name,
			entryType:   entryType,
			startSector: start,
			streamSize:  size,
		})
	}
	return entries, nil
}

func extractStream(b []byte, h *oleHeader, fat []uint32, start uint32, size uint64) ([]byte, error) {
	chain, err := walkChain(fat, start)
	if err != nil {
		return nil, err
	}
	var out bytes.Buffer
	for _, sid := range chain {
		off := sectorOffset(h, sid)
		end := off + h.sectorSize
		if end > len(b) {
			return nil, fmt.Errorf("stream sector out of range")
		}
		out.Write(b[off:end])
	}
	buf := out.Bytes()
	if size <= uint64(len(buf)) {
		return buf[:size], nil
	}
	return buf, nil
}

func extractWorkbookStream(b []byte) ([]byte, error) {
	h, err := readHeader(b)
	if err != nil {
		return nil, err
	}
	fat, err := readFAT(b, h)
	if err != nil {
		return nil, err
	}
	entries, err := readDirectory(b, h, fat)
	if err != nil {
		return nil, err
	}
	var wb *dirEntry
	for i := range entries {
		if entries[i].entryType == 2 && (entries[i].name == "Workbook" || entries[i].name == "Book") {
			wb = &entries[i]
			break
		}
	}
	if wb == nil {
		return nil, fmt.Errorf("Workbook stream not found")
	}
	return extractStream(b, h, fat, wb.startSector, wb.streamSize)
}

// ---------------------------
// BIFF8 解析（极简实现）
// ---------------------------

const (
	biffRecBOF        = 0x0809
	biffRecEOF        = 0x000A
	biffRecBoundSheet = 0x0085
	biffRecSST        = 0x00FC
	biffRecLabelSST   = 0x00FD
	biffRecNumber     = 0x0203
	biffRecRSTRING    = 0x00D6
	biffRecLABEL      = 0x0204
)

type sstString struct {
	text string
	enc  string
}

var xlsLogger func(string, ...any)

func init() {
	if os.Getenv("XLS_DEBUG") == "1" {
		xlsLogger = func(format string, args ...any) {
			fmt.Printf("[xls] "+format+"\n", args...)
		}
	}
}

type sstReader struct {
	segs [][]byte
	si   int
	off  int
}

func (sr *sstReader) has() bool {
	return sr.si < len(sr.segs)
}
func (sr *sstReader) ensure() error {
	if !sr.has() {
		return fmt.Errorf("sst stream exhausted")
	}
	return nil
}
func (sr *sstReader) readByte() (byte, error) {
	if err := sr.ensure(); err != nil {
		return 0, err
	}
	seg := sr.segs[sr.si]
	if sr.off >= len(seg) {
		sr.si++
		sr.off = 0
		if err := sr.ensure(); err != nil {
			return 0, err
		}
		seg = sr.segs[sr.si]
	}
	b := seg[sr.off]
	sr.off++
	return b, nil
}
func (sr *sstReader) readN(n int) ([]byte, error) {
	var out []byte
	for n > 0 {
		if err := sr.ensure(); err != nil {
			return nil, err
		}
		seg := sr.segs[sr.si]
		remain := len(seg) - sr.off
		if remain <= 0 {
			sr.si++
			sr.off = 0
			continue
		}
		take := remain
		if take > n {
			take = n
		}
		out = append(out, seg[sr.off:sr.off+take]...)
		sr.off += take
		n -= take
	}
	return out, nil
}

func parseStreamUnicodeString(sr *sstReader) (sstString, error) {
	hdr, err := sr.readN(3)
	if err != nil {
		return sstString{}, err
	}
	cch := binary.LittleEndian.Uint16(hdr[0:2])
	flags := hdr[2]
	var rt uint16
	var extSize uint32
	if flags&0x08 != 0 {
		buf, err := sr.readN(2)
		if err != nil {
			return sstString{}, err
		}
		rt = binary.LittleEndian.Uint16(buf)
	}
	if flags&0x04 != 0 {
		buf, err := sr.readN(4)
		if err != nil {
			return sstString{}, err
		}
		extSize = binary.LittleEndian.Uint32(buf)
	}
	var text string
	var encUsed string
	if flags&0x01 != 0 {
		total := int(cch) * 2
		u := make([]uint16, 0, int(cch))
		for total > 0 {
			if sr.off >= len(sr.segs[sr.si]) {
				sr.si++
				sr.off = 0
			}
			if sr.si >= len(sr.segs) {
				return sstString{}, fmt.Errorf("string continuation missing")
			}
			chunk, err := sr.readN(2)
			if err != nil {
				return sstString{}, err
			}
			u = append(u, binary.LittleEndian.Uint16(chunk))
			total -= 2
			if total > 0 && sr.off >= len(sr.segs[sr.si]) {
				if sr.si+1 < len(sr.segs) {
					sr.si++
					sr.off = 0
					if sr.si < len(sr.segs) && len(sr.segs[sr.si]) > 0 {
						flag, _ := sr.readByte()
						if flag&0x01 == 0 {
							remain := int(cch) - len(u)
							raw, err := sr.readN(remain)
							if err != nil {
								return sstString{}, err
							}
							runes := make([]rune, len(raw))
							for i, by := range raw {
								runes[i] = rune(by)
							}
							text = string(utf16.Decode(u)) + string(runes)
							encUsed = "utf-16le"
							goto tail
						}
					}
				}
			}
		}
		text = string(utf16.Decode(u))
		encUsed = "utf-16le"
	} else {
		total := int(cch)
		raw := make([]byte, 0, total)
		for total > 0 {
			if sr.off >= len(sr.segs[sr.si]) {
				sr.si++
				sr.off = 0
			}
			if sr.si >= len(sr.segs) {
				return sstString{}, fmt.Errorf("string continuation missing")
			}
			seg := sr.segs[sr.si]
			toTake := len(seg) - sr.off
			if toTake == 0 {
				continue
			}
			if toTake > total {
				toTake = total
			}
			raw = append(raw, seg[sr.off:sr.off+toTake]...)
			sr.off += toTake
			total -= toTake
			if total > 0 && sr.off >= len(seg) {
				if sr.si+1 < len(sr.segs) {
					sr.si++
					sr.off = 0
					if sr.si < len(sr.segs) && len(sr.segs[sr.si]) > 0 {
						flag, err := sr.readByte()
						if err != nil {
							return sstString{}, err
						}
						if flag&0x01 != 0 {
							rem := total * 2
							u := make([]uint16, 0, rem/2)
							for rem > 0 {
								chunk, err := sr.readN(2)
								if err != nil {
									return sstString{}, err
								}
								u = append(u, binary.LittleEndian.Uint16(chunk))
								rem -= 2
							}
							runes := make([]rune, len(raw))
							for i, by := range raw {
								runes[i] = rune(by)
							}
							text = string(runes) + string(utf16.Decode(u))
							encUsed = "utf-16le"
							goto tail
						}
					}
				}
			}
		}
		runes := make([]rune, len(raw))
		for i, by := range raw {
			runes[i] = rune(by)
		}
		text = string(runes)
		encUsed = "ansi"
	}
tail:
	if flags&0x08 != 0 {
		if _, err := sr.readN(int(rt) * 4); err != nil {
			return sstString{}, err
		}
	}
	if flags&0x04 != 0 {
		if _, err := sr.readN(int(extSize)); err != nil {
			return sstString{}, err
		}
	}
	return sstString{text: text, enc: encUsed}, nil
}

func u16(b []byte, off int) uint16 {
	return binary.LittleEndian.Uint16(b[off : off+2])
}
func u32(b []byte, off int) uint32 {
	return binary.LittleEndian.Uint32(b[off : off+4])
}
func f64(b []byte, off int) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(b[off : off+8]))
}

func parseUnicodeString(b []byte) (sstString, int, error) {
	if len(b) < 3 {
		return sstString{}, 0, fmt.Errorf("invalid string header")
	}
	cch := u16(b, 0)
	flags := b[2]
	off := 3
	var rt uint16
	var extSize uint32
	if flags&0x08 != 0 {
		if len(b) < off+2 {
			return sstString{}, 0, fmt.Errorf("invalid rich header")
		}
		rt = u16(b, off)
		off += 2
	}
	if flags&0x04 != 0 {
		if len(b) < off+4 {
			return sstString{}, 0, fmt.Errorf("invalid ext header")
		}
		extSize = u32(b, off)
		off += 4
	}
	var text string
	var encUsed string
	if flags&0x01 == 0 {
		if len(b) < off+int(cch) {
			return sstString{}, 0, fmt.Errorf("string length out of range")
		}
		raw := b[off : off+int(cch)]
		runes := make([]rune, len(raw))
		for i, by := range raw {
			runes[i] = rune(by)
		}
		text = string(runes)
		encUsed = "ansi"
		off += int(cch)
	} else {
		byteLen := int(cch) * 2
		if len(b) < off+byteLen {
			return sstString{}, 0, fmt.Errorf("string length out of range")
		}
		u := make([]uint16, cch)
		for i := 0; i < int(cch); i++ {
			u[i] = u16(b, off+2*i)
		}
		text = string(utf16.Decode(u))
		encUsed = "utf-16le"
		off += byteLen
	}
	// skip trailing formatting runs and phonetic/ext data
	if flags&0x08 != 0 {
		trail := int(rt) * 4
		if len(b) < off+trail {
			return sstString{}, 0, fmt.Errorf("rich runs out of range")
		}
		off += trail
	}
	if flags&0x04 != 0 {
		if len(b) < off+int(extSize) {
			return sstString{}, 0, fmt.Errorf("ext data out of range")
		}
		off += int(extSize)
	}
	return sstString{text: text, enc: encUsed}, off, nil
}

func parseBIFF8Workbook(b []byte) ([][]string, string, error) {
	var sst []sstString
	var sheetOffset int = -1
	var encoding string = "utf-16le"

	for pos := 0; pos+4 <= len(b); {
		rid := u16(b, pos)
		rlen := int(u16(b, pos+2))
		pos += 4
		if pos+rlen > len(b) {
			return nil, "", fmt.Errorf("record length out of range")
		}
		payload := b[pos : pos+rlen]
		pos += rlen
		switch rid {
		case biffRecBoundSheet:
			if rlen >= 8 {
				off := int(u32(payload, 0))
				if sheetOffset < 0 {
					sheetOffset = off
				}
				if xlsLogger != nil {
					xlsLogger("BoundSheet offset=%d", off)
				}
			}
		case biffRecSST:
			if rlen >= 8 {
				total := int(u32(payload, 0))
				unique := int(u32(payload, 4))
				var segs [][]byte
				if 8 < rlen {
					segs = append(segs, payload[8:rlen])
				}
				for pos < len(b) {
					nrid := u16(b, pos)
					nlen := int(u16(b, pos+2))
					if nrid != 0x003C {
						break
					}
					pos += 4
					if pos+nlen > len(b) {
						break
					}
					segs = append(segs, b[pos:pos+nlen])
					pos += nlen
				}
				sr := &sstReader{segs: segs}
				count := unique
				if count <= 0 {
					count = total
				}
				if xlsLogger != nil {
					xlsLogger("SST total=%d unique=%d segments=%d", total, unique, len(segs))
				}
				for i := 0; i < count && sr.has(); i++ {
					s, err := parseStreamUnicodeString(sr)
					if err != nil {
						break
					}
					if s.enc == "utf-16le" {
						encoding = "utf-16le"
					}
					sst = append(sst, s)
					if xlsLogger != nil && i < 5 {
						xlsLogger("SST[%d] enc=%s text_sample=%s", i, s.enc, s.text)
					}
				}
			}
		}
		if rid == biffRecEOF && sheetOffset >= 0 {
			break
		}
	}
	if sheetOffset < 0 {
		return nil, "", fmt.Errorf("worksheet offset not found")
	}
	if sheetOffset > len(b) {
		return nil, "", fmt.Errorf("worksheet offset out of range")
	}
	var cells = map[uint16]map[uint16]string{}
	for pos := sheetOffset; pos+4 <= len(b); {
		rid := u16(b, pos)
		rlen := int(u16(b, pos+2))
		pos += 4
		if pos+rlen > len(b) {
			break
		}
		payload := b[pos : pos+rlen]
		pos += rlen
		switch rid {
		case biffRecLabelSST:
			if rlen >= 10 {
				row := u16(payload, 0)
				col := u16(payload, 2)
				idx := u32(payload, 6)
				if int(idx) < len(sst) {
					if _, ok := cells[row]; !ok {
						cells[row] = map[uint16]string{}
					}
					cells[row][col] = sst[idx].text
					if xlsLogger != nil {
						xlsLogger("LabelSST r=%d c=%d sst=%d val=%s", row, col, idx, sst[idx].text)
					}
				}
			}
		case biffRecRSTRING:
			if rlen >= 6 {
				row := u16(payload, 0)
				col := u16(payload, 2)
				text, used, err := parseUnicodeString(payload[6:])
				if err == nil && used > 0 {
					if _, ok := cells[row]; !ok {
						cells[row] = map[uint16]string{}
					}
					cells[row][col] = text.text
					if xlsLogger != nil {
						xlsLogger("RSTRING r=%d c=%d val=%s", row, col, text.text)
					}
				}
			}
		case biffRecLABEL:
			if rlen >= 8 {
				row := u16(payload, 0)
				col := u16(payload, 2)
				cch := int(u16(payload, 6))
				if 8+cch <= rlen {
					raw := payload[8 : 8+cch]
					runes := make([]rune, len(raw))
					for i, by := range raw {
						runes[i] = rune(by)
					}
					if _, ok := cells[row]; !ok {
						cells[row] = map[uint16]string{}
					}
					cells[row][col] = string(runes)
					if xlsLogger != nil {
						xlsLogger("LABEL r=%d c=%d val=%s", row, col, cells[row][col])
					}
				}
			}
		case biffRecNumber:
			if rlen >= 14 {
				row := u16(payload, 0)
				col := u16(payload, 2)
				val := f64(payload, 6)
				if _, ok := cells[row]; !ok {
					cells[row] = map[uint16]string{}
				}
				s := fmt.Sprintf("%.15g", val)
				cells[row][col] = s
				if xlsLogger != nil {
					xlsLogger("NUMBER r=%d c=%d val=%s", row, col, s)
				}
			}
		case biffRecEOF:
			goto BUILD
		}
	}

BUILD:
	var maxRow uint16
	var maxCol uint16
	for r, cols := range cells {
		if r > maxRow {
			maxRow = r
		}
		for c := range cols {
			if c > maxCol {
				maxCol = c
			}
		}
	}
	var rows [][]string
	for r := uint16(0); r <= maxRow; r++ {
		line := make([]string, int(maxCol)+1)
		if cols, ok := cells[r]; ok {
			for c, v := range cols {
				line[int(c)] = v
			}
		}
		end := len(line)
		for end > 0 && line[end-1] == "" {
			end--
		}
		rows = append(rows, line[:end])
	}
	return rows, encoding, nil
}
