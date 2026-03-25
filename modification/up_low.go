package modification

import (
	"regexp"
	"strconv"
	"strings"
)

var tagRegex = regexp.MustCompile(`\((up|low|cap)(?:,\s*(\d+))?\)`)

func ConvertCase(line string) string {
	for {
		loc := tagRegex.FindStringIndex(line)
		if loc == nil {
			break
		}

		match := tagRegex.FindStringSubmatch(line[loc[0]:loc[1]])
		tagType := match[1]
		count := 1
		if match[2] != "" {
			count, _ = strconv.Atoi(match[2])
		}

		prefix := line[:loc[0]]
		words := strings.Fields(prefix)

		startIdx := len(words) - count
		if startIdx < 0 {
			startIdx = 0
		}

		for i := startIdx; i < len(words); i++ {
			oldWord := words[i]
			newWord := applyTransform(oldWord, tagType)
			
			lastIndex := strings.LastIndex(prefix, oldWord)
			if lastIndex != -1 {
				prefix = prefix[:lastIndex] + newWord + prefix[lastIndex+len(oldWord):]
			}
		}

		line = prefix + line[loc[1]:]
	}

	return cleanSpaces(line)
}

func applyTransform(word, tagType string) string {
	switch tagType {
	case "up":
		return strings.ToUpper(word)
	case "low":
		return strings.ToLower(word)
	case "cap":
		if len(word) > 0 {
			return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}
	return word
}

func cleanSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}