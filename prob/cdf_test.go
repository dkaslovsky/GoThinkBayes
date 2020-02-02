package prob

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCdfPercentile(t *testing.T) {
	tests := map[string]struct {
		prob       map[float64]float64
		percentile float64
		expected   float64
		shouldErr  bool
	}{
		"percentile less than 0": {
			prob:       map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
			percentile: -0.5,
			expected:   0,
			shouldErr:  true,
		},
		"percentile greater than 1": {
			prob:       map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
			percentile: 5,
			expected:   0,
			shouldErr:  true,
		},
		"percentile 0": {
			prob:       map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
			percentile: 0,
			expected:   1,
			shouldErr:  false,
		},
		"percentile 1": {
			prob:       map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
			percentile: 1,
			expected:   4,
			shouldErr:  false,
		},
		"percentile 0.5": {
			prob:       map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
			percentile: 0.5,
			expected:   2,
			shouldErr:  false,
		},
		"percentile 0.51": {
			prob:       map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
			percentile: 0.51,
			expected:   3,
			shouldErr:  false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			c, err := NewCdf(test.prob)
			require.Nil(t, err)

			res, err := c.Percentile(test.percentile)
			if test.shouldErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			assert.Equal(t, test.expected, res)
		})
	}
}

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
