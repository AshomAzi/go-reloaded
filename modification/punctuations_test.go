package modification

import "testing"

func TestPunctuation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Close gap", "Hello , world !", "Hello, world!"},
		{"Multiple marks", "Wait ...", "Wait..."},
		{"Space after", "one,two", "one, two"},
		{"Quotes", "I am ' happy '", "I am 'happy'"},
		{"Combined", "word ,word ' quote '", "word, word 'quote'"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Punctuation(tt.input); got != tt.expected {
				t.Errorf("Punctuation() = %q, want %q", got, tt.expected)
			}
		})
	}
}