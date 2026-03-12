package modification

import "strings"

// FixArticles is Capitalized so it can be imported
func FixArticles(line string) string {
	words := strings.Fields(line)
	vowels := "aeiouhAEIOUH"

	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" {
			next := words[i+1]
			if len(next) > 0 && strings.ContainsAny(string(next[0]), vowels) {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}
	return strings.Join(words, " ")
}