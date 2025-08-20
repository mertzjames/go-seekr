package cmd

import (
	"slices"
	"testing"
)

func Test_checkIfText(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected bool
	}{
		{"Text file", "./testdata/python/config.py", true},
		{"Binary file", "./testdata/example.bin", false},
		{"JavaScript file", "./testdata/javascript/database.js", true},
		{"JSON file", "./testdata/json/settings.json", true},
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

func Test_uniqueStr(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"No duplicates", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"With duplicates", []string{"a", "b", "a", "c"}, []string{"a", "b", "c"}},
		{"All duplicates", []string{"a", "a", "a"}, []string{"a"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := unique(tt.input)
			if !slices.Equal(got, tt.expected) {
				t.Errorf("unique(%v) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func Test_uniqueInt(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"No duplicates", []int{1, 2, 3}, []int{1, 2, 3}},
		{"With duplicates", []int{1, 2, 1, 3}, []int{1, 2, 3}},
		{"All duplicates", []int{1, 1, 1}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := unique(tt.input)
			if !slices.Equal(got, tt.expected) {
				t.Errorf("uniqueInt(%v) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func Test_uniqueSlices(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{"No duplicates", [][]int{{1, 2}, {3, 4}}, [][]int{{1, 2}, {3, 4}}},
		{"With duplicates", [][]int{{1, 2}, {2, 3}, {1, 2}}, [][]int{{1, 2}, {2, 3}}},
		{"All duplicates", [][]int{{1, 2}, {1, 2}, {1, 2}}, [][]int{{1, 2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniqueSlices(tt.input)
			for i := range got {
				if !slices.Equal(got[i], tt.expected[i]) {
					t.Errorf("uniqueSlices(%v) = %v; want %v", tt.input, got, tt.expected)
				}
			}

		})
	}
}

func Test_extractVars(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  []string
		delimeter DelimeterOptions
	}{
		{"Single line", "MY_SECRET", []string{"MY_SECRET"}, DelimeterNewline},
		{"Multiple lines", "MY_SECRET\nANOTHER_SECRET", []string{"MY_SECRET", "ANOTHER_SECRET"}, DelimeterNewline},
		{"Commented Line", "#MY_SECRET", []string{}, DelimeterNewline},
		{"Empty Line", "\n", []string{}, DelimeterNewline},
		{"Whitespace Line", "   ", []string{}, DelimeterNewline},
		{"Regex Pattern", "^([a-zA-Z0-9_\\-\\.]+)@([a-zA-Z0-9_\\-\\.]+)\\.([a-zA-Z]{2,5})$", []string{"^([a-zA-Z0-9_\\-\\.]+)@([a-zA-Z0-9_\\-\\.]+)\\.([a-zA-Z]{2,5})$"}, DelimeterNewline},
		{"Mixed Content", "MY_SECRET\n# Commented Line\nANOTHER_SECRET\n([a-zA-Z0-9_\\-\\.]+)", []string{"MY_SECRET", "ANOTHER_SECRET", "([a-zA-Z0-9_\\-\\.]+)"}, DelimeterNewline},
		{"Comma Delimeter", "MY_SECRET,ANOTHER_SECRET", []string{"MY_SECRET", "ANOTHER_SECRET"}, DelimeterComma},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractVars(tt.input, tt.delimeter)
			if !slices.Equal(got, tt.expected) {
				t.Errorf("extractContent(%q) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}
