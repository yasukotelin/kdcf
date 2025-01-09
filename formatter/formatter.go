package formatter

import (
	"strings"
	"unicode"
)

// formatKotlinData はKotlinデータクラスの構造を整形する
func formatKotlinData(text string) string {
	// 入力が `"` で囲まれている場合に取り除く
	if strings.HasPrefix(text, "\"") && strings.HasSuffix(text, "\"") {
		text = strings.TrimPrefix(strings.TrimSuffix(text, "\""), "\"")
	}

	var result strings.Builder
	var indentLevel int
	indent := func() string {
		return strings.Repeat("    ", indentLevel)
	}

	for i, char := range text {
		switch char {
		case '(':
			// 開き括弧の次に改行とインデントを追加
			result.WriteRune(char)
			result.WriteString("\n")
			indentLevel++
			result.WriteString(indent())
		case ')':
			// 閉じ括弧の前に改行とインデントを調整
			result.WriteString("\n")
			indentLevel--
			result.WriteString(indent())
			result.WriteRune(char)
		case '[':
			// 開き角括弧の次に改行とインデントを追加
			result.WriteRune(char)
			result.WriteString("\n")
			indentLevel++
			result.WriteString(indent())
		case ']':
			// 閉じ角括弧の前に改行とインデントを調整
			result.WriteString("\n")
			indentLevel--
			result.WriteString(indent())
			result.WriteRune(char)
		case ',':
			// カンマの次に改行とインデントを追加
			result.WriteRune(char)
			result.WriteString("\n")
			result.WriteString(indent())
		case '=':
			// 等号の前後にスペースを追加
			result.WriteString(" = ")
		default:
			// 他の文字はそのまま追加
			if !unicode.IsSpace(char) || (i > 0 && text[i-1] != ',') {
				result.WriteRune(char)
			}
		}
	}

	return result.String()
}
