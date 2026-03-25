package modification

import "strings"

func FixArticles(line string) string {
	words := strings.Fields(line)
	vowels := "aeiouAEIOU" // Removed 'h' from the general vowel list
	
	// Define the specific words where 'h' is silent
	silentH := map[string]bool{
		"hour":   true,
		"honest": true,
		"honor":  true,
		"honour": true,
		"heir":   true,
		"herb":   true, // American English usage
	}

	for i := 0; i < len(words)-1; i++ {
		currentWord := words[i]
		lowerCurrent := strings.ToLower(currentWord)
		nextWord := words[i+1]

		if len(nextWord) == 0 {
			continue
		}

		// Normalize the next word to lowercase and strip punctuation for checking
		nextWordClean := strings.ToLower(strings.Trim(nextWord, ".,!?:;"))
		firstLetterNext := string(nextWord[0])
		
		// Logic to determine if we need "an"
		// 1. Starts with a standard vowel (a, e, i, o, u)
		// 2. Or it is one of our special silent 'h' words
		needsAn := strings.ContainsAny(firstLetterNext, vowels) || silentH[nextWordClean]

		if lowerCurrent == "a" && needsAn {
			if currentWord == "A" {
				words[i] = "An"
			} else {
				words[i] = "an"
			}
		} else if lowerCurrent == "an" && !needsAn {
			if currentWord == "An" {
				words[i] = "A"
			} else {
				words[i] = "a"
			}
		}
	}
	return strings.Join(words, " ")
}