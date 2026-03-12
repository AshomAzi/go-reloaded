package modification

import (
	"regexp"
	"strconv"
	"strings"
)

// Regex to catch tags like (up), (low), (cap) OR (up, 3), (low, 5), etc.
var tagRegex = regexp.MustCompile(`\((up|low|cap)(?:,\s*(\d+))?\)`)

func ConvertCase(line string) string {
	// 1. First, handle your existing Hex/Bin logic (omitted for brevity)
	// line = ConvertHexBin(line)

	// 2. Handle Case Tags
	words := strings.Fields(line)

	for i := 0; i < len(words); i++ {
		match := tagRegex.FindStringSubmatch(words[i])
		if match == nil {
			continue
		}

		tagType := match[1] // up, low, or cap
		count := 1          // Default to 1 if no number is provided
		if match[2] != "" {
			count, _ = strconv.Atoi(match[2])
		}

		// Apply transformation to the previous 'count' words
		for j := 1; j <= count; j++ {
			targetIdx := i - j
			if targetIdx >= 0 {
				words[targetIdx] = applyTransform(words[targetIdx], tagType)
			}
		}

		// Remove the tag from the slice
		words = append(words[:i], words[i+1:]...)
		i-- // Adjust index because we removed an element
	}

	return strings.Join(words, " ")
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
