package hexv

import (
    "fmt"
    "strconv"
    "strings"
)

// ProcessLine takes a raw string and converts "XX (hex)" to decimal
func ProcessLine(line string) string {
    words := strings.Fields(line) // Splits by whitespace
    for i := 0; i < len(words); i++ {
        if words[i] == "(hex)" && i > 0 {
            // Get the word before "(hex)"
            hexVal := words[i-1]
            
            // Convert Hex (Base 16) to Decimal (Base 10)
            decimalVal, err := strconv.ParseInt(hexVal, 16, 64)
            if err == nil {
                // Replace the hex word with the decimal string
                words[i-1] = fmt.Sprintf("%d", decimalVal)
                // Remove the "(hex)" marker from the slice
                words = append(words[:i], words[i+1:]...)
                // Adjust index because we removed an element
                i-- 
            }
        }
    }
    return strings.Join(words, " ")
}
