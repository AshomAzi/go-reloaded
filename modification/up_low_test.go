package modification

import "testing"

func TestConvertCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello (up)", "HELLO"},
		{"STOP (low)", "stop"},
		{"italy (cap)", "Italy"},
		{"this is a test (up, 2)", "this is A TEST"},
		{"ready set go (cap, 3)", "Ready Set Go"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ConvertCase(tt.input); got != tt.expected {
				t.Errorf("ConvertCase(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}