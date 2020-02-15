package prob

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCredibleIntervalPercentiles(t *testing.T) {
	tests := map[string]struct {
		len           float64
		expectedLower float64
		expectedUpper float64
	}{
		"len = 100": {
			len:           100,
			expectedLower: 0,
			expectedUpper: 1,
		},
		"len = 1": {
			len:           1,
			expectedLower: 0.495,
			expectedUpper: 0.505,
		},
		"len has no decimal": {
			len:           90,
			expectedLower: 0.05,
			expectedUpper: 0.95,
		},
		"len has decimal": {
			len:           90.2,
			expectedLower: 0.049,
			expectedUpper: 0.951,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			lower, upper := getCredibleIntervalPercentiles(test.len)

			if test.expectedLower == 0 {
				assert.Equal(t, 0.0, lower)
			} else {
				assert.InEpsilon(t, test.expectedLower, lower, float64EqualTol)
			}
			assert.InEpsilon(t, test.expectedUpper, upper, float64EqualTol)
		})
	}
}
