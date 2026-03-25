package modification

import (
	"fmt"
	"regexp"
	"strconv"
)

// Regex for finding words followed by (hex) or (bin) tags
var hexBin = regexp.MustCompile(`([a-zA-Z0-9]+)\s*(\S*?\((hex|bin)\)\S*)`)

// 1. The "Worker" function: Handles the logic for ONE match
func replaceMatch(line string) string {
    subMatches := hexBin.FindStringSubmatch(line)
    if len(subMatches) < 4 {
        return line
    }

    valStr := subMatches[1]
    tagType := subMatches[3]
    base := 16
	
    if tagType == "bin" {
        base = 2
    } else {
        
    }

    result, err := strconv.ParseInt(valStr, base, 64)
    if err != nil {
        return line
    }

    return fmt.Sprintf("%d", result)
}

// 2. The "Entry" function: What you call from main.go
func ProcessLine(line string) string {
    // We just "plug in" the worker function here
    return hexBin.ReplaceAllStringFunc(line, replaceMatch)
}