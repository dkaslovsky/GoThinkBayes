package prob

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestNewBeta(t *testing.T) {
	tests := map[string]struct {
		alpha     float64
		beta      float64
		shouldErr bool
	}{
		"alpha is negative": {
			alpha:     -1,
			beta:      1,
			shouldErr: true,
		},
		"alpha is zero": {
			alpha:     0,
			beta:      1,
			shouldErr: true,
		},
		"beta is negative": {
			alpha:     1,
			beta:      -1,
			shouldErr: true,
		},
		"beta is zero": {
			alpha:     1,
			beta:      0,
			shouldErr: true,
		},
		"alpha and beta are positive": {
			alpha:     10,
			beta:      5,
			shouldErr: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			b, err := NewBeta(test.alpha, test.beta)

			if test.shouldErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			assert.Equal(t, test.alpha, b.alpha)
			assert.Equal(t, test.beta, b.beta)
		})
	}
}

func TestBetaUpdate(t *testing.T) {
	t.Run("Beta Update", func(t *testing.T) {
		b := &Beta{alpha: 10.0, beta: 5.0}
		b.Update(13.0, 0.1)

		assert.Equal(t, 23.0, b.alpha)
		assert.Equal(t, 5.1, b.beta)
	})
}

func TestBetaMean(t *testing.T) {
	t.Run("Beta Mean", func(t *testing.T) {
		b := &Beta{alpha: 1.0, beta: 3.0}
		m := b.Mean()

		assert.Equal(t, 0.25, m)
	})
}

func TestBetaEvalPdf(t *testing.T) {
	tests := map[string]struct {
		b            *Beta
		val          float64
		expectedProb float64
	}{
		"Pr(0)": {
			b:            &Beta{12.123, 4.321},
			val:          0.0,
			expectedProb: 0.0,
		},
		"Pr(0.5)": {
			b:            &Beta{2, 2},
			val:          0.5,
			expectedProb: 0.25,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			pr := test.b.EvalPdf(test.val)

			assert.Equal(t, test.expectedProb, pr)
		})
	}
}

func TestBetaMakePmf(t *testing.T) {
	t.Run("Beta MakePmf", func(t *testing.T) {
		nSteps := 6
		b := &Beta{2, 2}
		expectedElems := map[float64]float64{
			0.0: 0.0,
			0.2: 0.16,
			0.4: 0.24,
			0.6: 0.24,
			0.8: 0.16,
			1.0: 0.0,
		}

		p := b.MakePmf(nSteps)

		assert.Equal(t, len(expectedElems), len(p.prob))

		for val, pr := range p.prob {
			require.Contains(t, expectedElems, val)
			expectedPr := expectedElems[val]
			if expectedPr == 0 {
				assert.Equal(t, 0.0, pr)
			} else {
				assert.InEpsilon(t, expectedPr, pr, float64EqualTol)
			}
		}
	})
}
