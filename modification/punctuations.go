package modification

import (
	"regexp"
	"strings"
)

func Punctuation(line string) string {
	// 1. Handle the "Close the gap" rule
	// This finds any whitespace followed by one or more punctuation marks (.,!?:;)
	// It pulls the entire group (like ...) to the previous word.
	// Result: "thinking ... You" -> "thinking... You"
	res := regexp.MustCompile(`\s+([.,!?:;]+)`).ReplaceAllString(line, "$1")

	// 2. Handle the "Space apart from the next one" rule
	// This ensures that if a punctuation group is touching a character on the right, 
	// a space is injected. 
	// Result: "there,and" -> "there, and"
	res = regexp.MustCompile(`([.,!?:;]+)([^.,!?:;\s])`).ReplaceAllString(res, "$1 $2")

	// 3. Cleanup Single Quotes
	// Removes internal spaces: ' word ' -> 'word'
	res = regexp.MustCompile(`'\s*(.*?)\s*'`).ReplaceAllString(res, "'$1'")

	// 4. Final Polish
	// Standardizes all spacing to prevent double-spaces from previous steps.
	return strings.Join(strings.Fields(res), " ")
}