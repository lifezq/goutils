// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package encode

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"unicode/utf8"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

// ForceDecodeCSV 强制尝试所有可能的编码，优先尝试UTF-8
func ForceDecodeCSV(file io.Reader) ([][]string, string, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, "", err
	}

	// 编码尝试顺序：UTF-8优先，然后是常见中文编码，最后是其他编码
	attempts := []struct {
		name     string
		decoder  encoding.Encoding
		priority int // 优先级，数字越大优先级越高
	}{
		// 最高优先级：UTF-8相关
		{"utf-8", encoding.Nop, 100},    // Nop 表示无操作，直接使用原始数据
		{"utf-8-bom", encoding.Nop, 99}, // 专门处理带BOM的UTF-8

		// 高优先级：常见中文编码
		{"gbk", simplifiedchinese.GBK, 90},
		{"gb18030", simplifiedchinese.GB18030, 85},
		{"gb2312", simplifiedchinese.HZGB2312, 80},

		// 中优先级：繁体中文和其他中文编码
		{"big5", traditionalchinese.Big5, 70},

		// 低优先级：原始检测结果和其他编码
		{"windows-1252", charmap.Windows1252, 50},
		{"iso-8859-1", charmap.ISO8859_1, 40},
	}

	var bestRecords [][]string
	var bestEncoding string
	var bestScore int

	for _, attempt := range attempts {

		var decoded []byte
		var decodeErr error

		// 特殊处理UTF-8
		if attempt.name == "utf-8" {
			decoded = removeBOM(data) // 移除BOM后直接使用
			decodeErr = nil
		} else if attempt.name == "utf-8-bom" {
			// 检查是否有BOM，如果有则保留
			if hasBOM(data) {
				decoded = data
			} else {
				continue // 没有BOM就跳过这个尝试
			}
			decodeErr = nil
		} else {
			// 其他编码正常解码
			decoded, decodeErr = decodeData(data, attempt.decoder.NewDecoder())
		}

		if decodeErr != nil {
			continue
		}

		// 验证解码后的数据是否是有效的UTF-8
		if !utf8.Valid(decoded) {
			continue
		}

		// 尝试解析CSV
		records, parseErr := parseCSVFromBytes(decoded)
		if parseErr != nil {
			continue
		}

		// 计算质量分数（基于优先级和内容质量）
		qualityScore := calculateQualityScore(records, attempt.priority)

		// 如果解析成功且有合理内容，直接返回（UTF-8优先）
		if attempt.name == "utf-8" && hasReasonableContent(records) {
			return records, "utf-8", nil
		}

		// 选择分数最高的结果
		if qualityScore > bestScore {
			bestScore = qualityScore
			bestRecords = records
			bestEncoding = attempt.name
		}
	}

	if bestRecords == nil {
		return nil, "", fmt.Errorf("所有编码尝试都失败，无法解码CSV文件")
	}

	return bestRecords, bestEncoding, nil
}

// removeBOM 移除UTF-8 BOM头
func removeBOM(data []byte) []byte {
	if len(data) >= 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF {
		return data[3:]
	}
	return data
}

// hasBOM 检查是否有UTF-8 BOM头
func hasBOM(data []byte) bool {
	return len(data) >= 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF
}

// decodeData 使用指定编码解码数据
func decodeData(data []byte, decoder transform.Transformer) ([]byte, error) {
	var buf bytes.Buffer
	reader := transform.NewReader(bytes.NewReader(data), decoder)
	_, err := io.Copy(&buf, reader)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// parseCSVFromBytes 从字节数据解析CSV，保留所有行（包括空行）
func parseCSVFromBytes(data []byte) ([][]string, error) {
	// 按行分割数据，保留所有行（包括空行）
	lines := bytes.Split(data, []byte("\n"))

	var records [][]string
	for _, line := range lines {
		// 处理空行
		if len(line) == 0 {
			records = append(records, []string{})
			continue
		}

		// 对非空行进行CSV解析
		csvReader := csv.NewReader(bytes.NewReader(line))
		csvReader.LazyQuotes = true
		csvReader.TrimLeadingSpace = true
		csvReader.FieldsPerRecord = -1 // 不检查字段数量

		record, err := csvReader.Read()
		if err == io.EOF {
			// 对于只包含空白字符的行，添加一个空记录
			if len(record) == 0 {
				records = append(records, []string{})
			}
			continue
		}
		if err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}

// calculateQualityScore 计算解码质量分数
func calculateQualityScore(records [][]string, basePriority int) int {
	if len(records) == 0 {
		return 0
	}

	score := basePriority

	// 内容质量加分项
	chineseCount := countChineseCharacters(records)
	readableCount := countReadableCharacters(records)
	recordCount := len(records)

	// 中文字符比例加分
	if readableCount > 0 {
		chineseRatio := float64(chineseCount) / float64(readableCount)
		if chineseRatio > 0.3 {
			score += int(chineseRatio * 50) // 中文字符比例高，大幅加分
		}
	}

	// 记录数量合理加分
	if recordCount >= 1 && recordCount <= 10000 {
		score += mins(recordCount/10, 50) // 最多加50分
	}

	// 字段一致性加分
	if hasConsistentFields(records) {
		score += 20
	}

	return score
}

// countChineseCharacters 计算中文字符数量
func countChineseCharacters(records [][]string) int {
	count := 0
	for _, record := range records {
		for _, field := range record {
			for _, r := range field {
				// 中文字符的Unicode范围
				if (r >= '\u4e00' && r <= '\u9fff') || // 基本汉字
					(r >= '\u3400' && r <= '\u4dbf') || // 扩展A
					(r >= '\uf900' && r <= '\ufaff') { // 兼容汉字
					count++
				}
			}
		}
	}
	return count
}

// countReadableCharacters 计算可读字符数量
func countReadableCharacters(records [][]string) int {
	count := 0
	for _, record := range records {
		for _, field := range record {
			for _, r := range field {
				if r >= ' ' && r <= '~' { // 基本ASCII可打印字符
					count++
				} else if r >= '\u4e00' && r <= '\u9fff' { // 中文字符
					count++
				}
			}
		}
	}
	return count
}

// hasConsistentFields 检查字段数量是否一致
func hasConsistentFields(records [][]string) bool {
	if len(records) <= 1 {
		return true
	}

	firstFieldCount := len(records[0])
	for i := 1; i < len(records); i++ {
		if len(records[i]) != firstFieldCount {
			return false
		}
	}
	return true
}

// hasReasonableContent 检查是否有合理的内容
func hasReasonableContent(records [][]string) bool {
	if len(records) == 0 {
		return false
	}

	// 检查是否有非空字段
	for _, record := range records {
		for _, field := range record {
			if len(field) > 0 {
				return true
			}
		}
	}
	return false
}

func mins(a, b int) int {
	if a < b {
		return a
	}
	return b
}
