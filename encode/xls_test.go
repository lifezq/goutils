// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package encode

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func buildTestWorkbook() []byte {
	var wb bytes.Buffer
	wb.Write(u16b(0x0809))
	wb.Write(u16b(16))
	wb.Write(u16b(0x0600))
	wb.Write(u16b(0x0005))
	wb.Write(make([]byte, 12))

	wb.Write(u16b(0x0085))
	wb.Write(u16b(8 + 1 + 8))
	sheetOffsetPlaceholder := wb.Len()
	wb.Write(u32b(0))
	wb.Write(u16b(0))
	wb.WriteByte(byte(len("Sheet1")))
	wb.WriteString("Sheet1")
	wb.Write(make([]byte, 8))

	var sst bytes.Buffer
	sst.Write(u32b(2))
	sst.Write(u32b(2))
	nihao := []uint16{'你', '好'}
	sst.Write(u16b(uint16(len(nihao))))
	sst.WriteByte(0x00)
	for _, ch := range nihao {
		sst.Write(u16b(ch))
	}
	world := []uint16{'w', 'o', 'r', 'l', 'd'}
	sst.Write(u16b(uint16(len(world))))
	sst.WriteByte(0x00)
	for _, ch := range world {
		sst.Write(u16b(ch))
	}
	wb.Write(u16b(0x00FC))
	wb.Write(u16b(uint16(sst.Len())))
	wb.Write(sst.Bytes())

	wb.Write(u16b(0x000A))
	wb.Write(u16b(0))

	sheetOffset := wb.Len()
	binary.LittleEndian.PutUint32(wb.Bytes()[sheetOffsetPlaceholder:sheetOffsetPlaceholder+4], uint32(sheetOffset))

	wb.Write(u16b(0x0809))
	wb.Write(u16b(16))
	wb.Write(u16b(0x0600))
	wb.Write(u16b(0x0010))
	wb.Write(make([]byte, 12))

	wb.Write(u16b(0x00FD))
	wb.Write(u16b(10))
	wb.Write(u16b(0))
	wb.Write(u16b(0))
	wb.Write(u16b(0))
	wb.Write(u32b(0))
	wb.Write(u16b(0x00FD))
	wb.Write(u16b(10))
	wb.Write(u16b(0))
	wb.Write(u16b(1))
	wb.Write(u16b(0))
	wb.Write(u32b(1))

	wb.Write(u16b(0x000A))
	wb.Write(u16b(0))

	return wb.Bytes()
}

func buildCFBFWithWorkbook(wb []byte) []byte {
	sectorSize := 512
	wbSectors := (len(wb) + sectorSize - 1) / sectorSize

	h := make([]byte, 512)
	copy(h[0:], []byte{0xD0, 0xCF, 0x11, 0xE0, 0xA1, 0xB1, 0x1A, 0x00})
	binary.LittleEndian.PutUint16(h[0x1C:0x1E], 0xFFFE)
	binary.LittleEndian.PutUint16(h[0x1E:0x20], 9)
	binary.LittleEndian.PutUint16(h[0x20:0x22], 6)
	binary.LittleEndian.PutUint32(h[0x2C:0x30], 1)
	binary.LittleEndian.PutUint32(h[0x30:0x34], 1)
	binary.LittleEndian.PutUint32(h[0x3C:0x40], 0xFFFFFFFF)
	binary.LittleEndian.PutUint32(h[0x40:0x44], 0)
	binary.LittleEndian.PutUint32(h[0x44:0x48], 0xFFFFFFFF)
	binary.LittleEndian.PutUint32(h[0x48:0x4C], 0)
	binary.LittleEndian.PutUint32(h[0x4C:0x50], 0)

	fat := make([]byte, sectorSize)
	entries := sectorSize / 4
	for i := 0; i < entries; i++ {
		binary.LittleEndian.PutUint32(fat[4*i:4*i+4], 0xFFFFFFFF)
	}
	binary.LittleEndian.PutUint32(fat[0:4], 0xFFFFFFFD)
	binary.LittleEndian.PutUint32(fat[4:8], 0xFFFFFFFE)
	for s := 2; s < 2+wbSectors-1; s++ {
		binary.LittleEndian.PutUint32(fat[4*s:4*s+4], uint32(s+1))
	}
	if wbSectors > 0 {
		last := 2 + wbSectors - 1
		binary.LittleEndian.PutUint32(fat[4*last:4*last+4], 0xFFFFFFFE)
	}

	dir := make([]byte, sectorSize)
	writeDirEntry(dir[0:128], "Root Entry", 5, 0, 0)
	writeDirEntry(dir[128:256], "Workbook", 2, 2, uint64(wbSectors*sectorSize))

	var wbBuf bytes.Buffer
	wbBuf.Grow(wbSectors * sectorSize)
	wbBuf.Write(wb)
	if pad := wbSectors*sectorSize - len(wb); pad > 0 {
		wbBuf.Write(make([]byte, pad))
	}

	var out bytes.Buffer
	out.Write(h)
	out.Write(fat)
	out.Write(dir)
	out.Write(wbBuf.Bytes())
	return out.Bytes()
}

func writeDirEntry(dst []byte, name string, typ byte, startSector uint32, size uint64) {
	for i := range dst {
		dst[i] = 0
	}
	u := []uint16{}
	for _, r := range name {
		u = append(u, uint16(r))
	}
	maxChars := (64 - 1)
	if len(u) > maxChars {
		u = u[:maxChars]
	}
	off := 0x00
	for i := 0; i < len(u); i++ {
		binary.LittleEndian.PutUint16(dst[off+2*i:off+2*i+2], u[i])
	}
	nlen := uint16(len(u)*2 + 2)
	binary.LittleEndian.PutUint16(dst[0x40:0x42], nlen)
	dst[0x42] = typ
	binary.LittleEndian.PutUint32(dst[0x74:0x78], startSector)
	binary.LittleEndian.PutUint64(dst[0x78:0x80], size)
}

func u16b(v uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, v)
	return b
}
func u32b(v uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, v)
	return b
}

func TestForceDecodeXLS(t *testing.T) {
	wb := buildTestWorkbook()
	container := buildCFBFWithWorkbook(wb)
	records, enc, err := ForceDecodeXLS(bytes.NewReader(container))
	if err != nil {
		t.Fatalf("ForceDecodeXLS error: %v", err)
	}
	if enc != "utf-16le" {
		t.Fatalf("expected utf-16le, got %s", enc)
	}
	if len(records) == 0 || len(records[0]) < 2 {
		t.Fatalf("unexpected records shape: %#v", records)
	}
	if records[0][0] != "你好" || records[0][1] != "world" {
		t.Fatalf("unexpected cell values: %#v", records[0])
	}
}
