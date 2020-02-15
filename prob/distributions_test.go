package prob

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCredibleIntervalPercentiles(t *testing.T) {
	tests := map[string]struct {
		l             float64
		expectedLower float64
		expectedUpper float64
	}{
		"len = 100": {
			l:             100,
			expectedLower: 0,
			expectedUpper: 1,
		},
		"len = 1": {
			l:             1,
			expectedLower: 0.495,
			expectedUpper: 0.505,
		},
		"len has no decimal": {
			l:             90,
			expectedLower: 0.05,
			expectedUpper: 0.95,
		},
		"len has decimal": {
			l:             90.2,
			expectedLower: 0.049,
			expectedUpper: 0.951,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			lower, upper := getCredibleIntervalPercentiles(test.l)

			if test.expectedLower == 0 {
				assert.Equal(t, 0.0, lower)
			} else {
				assert.InEpsilon(t, test.expectedLower, lower, float64EqualTol)
			}
			assert.InEpsilon(t, test.expectedUpper, upper, float64EqualTol)
		})
	}
}
