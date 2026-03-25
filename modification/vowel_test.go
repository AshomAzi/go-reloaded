package modification

import "testing"

func TestFixArticles(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Standard Vowels
		{"A to An (vowel)", "a apple", "an apple"},
		{"An to A (consonant)", "an cat", "a cat"},
		
		// Silent H (Should use 'an')
		{"Silent H - hour", "a hour", "an hour"},
		{"Silent H - honest", "a honest", "an honest"},
		{"Silent H - heir", "A heir", "An heir"},
		
		// Voiced H (Should use 'a')
		{"Voiced H - house", "an house", "a house"},
		{"Voiced H - happy", "an happy", "a happy"},
		{"Voiced H - helicopter", "An helicopter", "A helicopter"},
		
		// Edge Cases
		{"Punctuation attached", "a hour.", "an hour."},
		{"Already correct", "an apple", "an apple"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FixArticles(tt.input); got != tt.expected {
				t.Errorf("FixArticles() = %q, want %q", got, tt.expected)
			}
		})
	}
}