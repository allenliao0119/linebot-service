package helper

import (
	"regexp"
	"strings"
)

func ConvertMarkdownToLineText(text string) string {
	// Remove #### heading symbols, keep the text
	text = regexp.MustCompile(`#{1,6}\s*`).ReplaceAllString(text, "")

	// Convert **bold** to plain text (optional: add symbols for emphasis)
	text = regexp.MustCompile(`\*\*([^*]+)\*\*`).ReplaceAllString(text, "【$1】")

	// Convert - list items to • or numbers
	text = regexp.MustCompile(`(?m)^-\s+`).ReplaceAllString(text, "• ")

	return strings.TrimSpace(text)
}
