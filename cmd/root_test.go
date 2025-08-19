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
