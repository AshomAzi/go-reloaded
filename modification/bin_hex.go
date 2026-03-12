package modification

import (
	"fmt"
	"regexp"
	"strconv"
)

// Regex for finding words followed by (hex) or (bin) tags
var hexBinRegex = regexp.MustCompile(`([a-zA-Z0-9]+)\s*(\S*?\((hex|bin)\)\S*)`)

// ProcessLine takes a single line of text, converts hex/bin tags, and returns the modified line.
func ProcessLine(line string) string {
	return hexBinRegex.ReplaceAllStringFunc(line, func(match string) string {
		subMatches := hexBinRegex.FindStringSubmatch(match)
		if len(subMatches) < 4 {
			return match
		}

		valStr := subMatches[1]
		tagType := subMatches[3] // "hex" or "bin"
		
		var result int64
		var convErr error

		// Determine base based on tag
		base := 16
		if tagType == "bin" {
			base = 2
		}

		result, convErr = strconv.ParseInt(valStr, base, 64)

		if convErr != nil {
			return match // If conversion fails, keep original text
		}

		// Return only the decimal value (removes the tag)
		return fmt.Sprintf("%d", result)
	})
}