package prob

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortKeys(t *testing.T) {
	tests := map[string]struct {
		input    map[float64]float64
		expected []float64
	}{
		"empty input": {
			input:    map[float64]float64{},
			expected: []float64{},
		},
		"input with single element": {
			input:    map[float64]float64{1.1: 0},
			expected: []float64{1.1},
		},
		"input with multiple elements": {
			input:    map[float64]float64{1.2: 100, 1.1: 0, 2: 50.5},
			expected: []float64{1.1, 1.2, 2},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			keys := sortKeys(test.input)
			assert.Equal(t, test.expected, keys)
		})
	}
}

func TestReverseMap(t *testing.T) {
	tests := map[string]struct {
		input    map[float64]int
		expected map[int]float64
	}{
		"empty input": {
			input:    map[float64]int{},
			expected: map[int]float64{},
		},
		"input with single element": {
			input:    map[float64]int{1.1: 100},
			expected: map[int]float64{100: 1.1},
		},
		"input with multiple elements": {
			input:    map[float64]int{1.1: 100, 2: 1},
			expected: map[int]float64{100: 1.1, 1: 2},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			revMap := reverseMap(test.input)
			assert.Equal(t, test.expected, revMap)
		})
	}
}
