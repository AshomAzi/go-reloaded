package modification

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Punctuation(line string) string {
	words := strings.Fields(line)

	// 1. Math and Case Logic (Left-to-Right)
	for i := 0; i < len(words); i++ {
		w := words[i]
		if w == "(hex)" || w == "(bin)" {
			base := 16
			if w == "(bin)" { base = 2 }
			val, _ := strconv.ParseInt(words[i-1], base, 64)
			words[i-1] = fmt.Sprintf("%d", val)
			words = append(words[:i], words[i+1:]...)
			i--
		} else if m := regexp.MustCompile(`\((up|low|cap)(?:,\s*(\d+))?\)`).FindStringSubmatch(w); m != nil {
			count, _ := strconv.Atoi(m[2]); if count == 0 { count = 1 }
			for j := 1; j <= count && i-j >= 0; j++ {
				words[i-j] = transform(words[i-j], m[1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	res := strings.Join(words, " ")

	// 2. Standard Punctuation (.,!?:;)
	// Pull punctuation to the left
	res = regexp.MustCompile(`\s+([.,!?:;]+)`).ReplaceAllString(res, "$1")
	// Ensure space to the right
	res = regexp.MustCompile(`([.,!?:;]+)([a-zA-Z0-9])`).ReplaceAllString(res, "$1 $2")

	// 3. Single Quote Logic 
	// This Regex finds a quote, any amount of space, content, any amount of space, and the closing quote
	// It replaces it with 'content' (no internal spaces)
	res = regexp.MustCompile(`'\s*(.*?)\s*'`).ReplaceAllString(res, "'$1'")

	return strings.TrimSpace(res)
}

func transform(w, tag string) string {
	if tag == "up" { return strings.ToUpper(w) }
	if tag == "low" { return strings.ToLower(w) }
	if len(w) > 0 { return strings.ToUpper(w[:1]) + strings.ToLower(w[1:]) }
	return w
}