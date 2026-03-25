package modification

import "testing"

func TestProcessLine(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1E (hex)", "30"},
		{"10 (bin)", "2"},
		{"f (hex)", "15"},
		{"111 (bin)", "7"},
		{"invalid (hex)", "invalid (hex)"}, // Should return original if parsing fails
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ProcessLine(tt.input); got != tt.expected {
				t.Errorf("ProcessLine(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}