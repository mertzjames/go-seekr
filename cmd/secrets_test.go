package cmd

import (
	"testing"
)

func TestCheckIfText(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected bool
	}{
		{"Text file", "./tests/secrets/config.py", true},
		{"Binary file", "./tests/secrets/example.bin", false},
		{"JavaScript file", "./tests/secrets/database.js", true},
		{"JSON file", "./tests/secrets/settings.json", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkIfText(tt.filePath)
			if got != tt.expected {
				t.Errorf("checkIfText(%q) = %v; want %v", tt.filePath, got, tt.expected)
			}
		})
	}
}
