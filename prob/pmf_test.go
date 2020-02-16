package prob

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// float64EqualTol is the tolerance at which we consider float64s equal
const float64EqualTol = 1e-9

func setupPmf(elems []*PmfElement) *Pmf {
	p := NewPmf()
	for _, elem := range elems {
		p.Set(elem)
	}
	return p
}

func getSum(m map[float64]float64) float64 {
	sum := 0.0
	for _, p := range m {
		sum += p
	}
	return sum
}

func TestNewPmf(t *testing.T) {
	t.Run("new Pmf", func(t *testing.T) {
		p := NewPmf()

		assert.Empty(t, p.prob)
	})
}

func TestSet(t *testing.T) {
	tests := map[string]struct {
		elements []*PmfElement
	}{
		"single element": {
			elements: []*PmfElement{NewPmfElement(1, 1)},
		},
		"multiple elements": {
			elements: []*PmfElement{
				NewPmfElement(1, 1),
				NewPmfElement(2.1, 10.5),
				NewPmfElement(3, 1.6),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := setupPmf(test.elements)

			for _, elem := range test.elements {
				require.Contains(t, p.prob, elem.Val)
				assert.Equal(t, elem.Prob, p.prob[elem.Val])
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	tests := map[string]struct {
		elements     []*PmfElement
		expectedProb map[float64]float64
		expectedSum  float64
	}{
		"empty Pmf": {
			elements:     []*PmfElement{},
			expectedProb: map[float64]float64{},
			expectedSum:  0,
		},
		"Pmf with single element": {
			elements: []*PmfElement{
				NewPmfElement(1, 100),
			},
			expectedProb: map[float64]float64{1: 1},
			expectedSum:  1,
		},
		"Pmf with multiple elements, uniform": {
			elements: []*PmfElement{
				NewPmfElement(1, 1),
				NewPmfElement(2, 1),
				NewPmfElement(3, 1),
				NewPmfElement(4, 1),
			},
			expectedProb: map[float64]float64{
				1: 0.25,
				2: 0.25,
				3: 0.25,
				4: 0.25,
			},
			expectedSum: 1,
		},
		"Pmf with multiple elements, nonuniform": {
			elements: []*PmfElement{
				NewPmfElement(1, 1),
				NewPmfElement(2, 5),
				NewPmfElement(3, 1),
				NewPmfElement(4, 1),
			},
			expectedProb: map[float64]float64{
				1: 0.125,
				2: 0.625,
				3: 0.125,
				4: 0.125,
			},
			expectedSum: 1,
		},
		"Pmf with multiple elements and sum 0": {
			elements: []*PmfElement{
				NewPmfElement(1, 0),
				NewPmfElement(2, 0),
				NewPmfElement(3, 0),
				NewPmfElement(4, 0),
			},
			expectedProb: map[float64]float64{
				1: 0,
				2: 0,
				3: 0,
				4: 0,
			},
			expectedSum: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := setupPmf(test.elements)

			p.Normalize()

			for elem, prob := range p.prob {
				assert.Equal(t, test.expectedProb[elem], prob)
			}
			if test.expectedSum == 0 {
				assert.Equal(t, test.expectedSum, getSum(p.prob))
			} else {
				assert.InEpsilon(t, test.expectedSum, getSum(p.prob), float64EqualTol)
			}
		})
	}
}

func TestMult(t *testing.T) {
	tests := map[string]struct {
		elements    []*PmfElement
		val         float64
		multFactor  float64
		expectedSum float64
	}{
		"element not in Pmf": {
			elements: []*PmfElement{
				NewPmfElement(1, 0.5),
				NewPmfElement(2, 0.5),
			},
			val:        3,
			multFactor: 0.5,
		},
		"element in Pmf": {
			elements: []*PmfElement{
				NewPmfElement(1, 0.5),
				NewPmfElement(2, 0.5),
			},
			val:        1,
			multFactor: 0.5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := setupPmf(test.elements)

			// store original probability before mutating
			origProb, found := p.prob[test.val]

			p.Mult(test.val, test.multFactor)

			// test probability of specified element correctly multiplied
			if found {
				assert.Equal(t, origProb*test.multFactor, p.prob[test.val])
			}
			// test other probabilities are unchanged
			for _, element := range test.elements {
				if element.Val == test.val {
					continue
				}
				assert.Equal(t, element.Prob, p.prob[element.Val])
			}
		})
	}
}

func TestProb(t *testing.T) {
	tests := map[string]struct {
		elements     []*PmfElement
		val          float64
		expectedProb float64
	}{
		"elememt in Pmf": {
			elements: []*PmfElement{
				NewPmfElement(1, 0.25),
				NewPmfElement(2, 0.75),
			},
			val:          1,
			expectedProb: 0.25,
		},
		"elememt not in Pmf": {
			elements: []*PmfElement{
				NewPmfElement(1, 0.25),
				NewPmfElement(2, 0.75),
			},
			val:          3,
			expectedProb: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := setupPmf(test.elements)

			prob := p.Prob(test.val)

			assert.Equal(t, test.expectedProb, prob)
		})
	}
}

func TestMean(t *testing.T) {
	tests := map[string]struct {
		elements     []*PmfElement
		expectedMean float64
		shouldErr    bool
	}{
		"empty Pmf": {
			elements:     []*PmfElement{},
			expectedMean: 0,
			shouldErr:    true,
		},
		"single elememt in Pmf": {
			elements: []*PmfElement{
				NewPmfElement(1, 1),
			},
			expectedMean: 1,
			shouldErr:    false,
		},
		"multiple elememts in Pmf": {
			elements: []*PmfElement{
				NewPmfElement(1, 0.25),
				NewPmfElement(2, 0.75),
			},
			expectedMean: 1.75,
			shouldErr:    false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := setupPmf(test.elements)

			m, err := p.Mean()

			if test.shouldErr {
				require.NotNil(t, err)
				return
			}
			assert.Equal(t, test.expectedMean, m)
		})
	}
}

func TestPmfPercentile(t *testing.T) {
	tests := map[string]struct {
		pmf        *Pmf
		percentile float64
		expected   float64
		shouldErr  bool
	}{
		"empty pmf": {
			pmf:        NewPmf(),
			percentile: 0.5,
			expected:   0,
			shouldErr:  true,
		},
		"percentile less than 0": {
			pmf: &Pmf{
				prob: map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
			},
			percentile: -0.5,
			expected:   0,
			shouldErr:  true,
		},
		"percentile greater than 1": {
			pmf: &Pmf{
				prob: map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
			},
			percentile: 5,
			expected:   0,
			shouldErr:  true,
		},
		"unnormalized pmf": {
			pmf: &Pmf{
				prob: map[float64]float64{1: 0.02, 2: 0.03, 3: 0.04, 4: 0.01},
			},
			percentile: 0.5,
			expected:   0,
			shouldErr:  true,
		},
		"unnormalized pmf with sum 1": {
			pmf: &Pmf{
				prob: map[float64]float64{1: 0.02, 2: 0.03, 3: 0.04, 4: 0.01},
			},
			percentile: 0.5,
			expected:   0,
			shouldErr:  true,
		},
		"percentile 0": {
			pmf: &Pmf{
				prob: map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
			},
			percentile: 0,
			expected:   1,
			shouldErr:  false,
		},
		"percentile 1": {
			pmf: &Pmf{
				prob: map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
			},
			percentile: 1,
			expected:   4,
			shouldErr:  false,
		},
		"percentile 0.5": {
			pmf: &Pmf{
				prob: map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
			},
			percentile: 0.5,
			expected:   2,
			shouldErr:  false,
		},
		"percentile 0.51": {
			pmf: &Pmf{
				prob: map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
			},
			percentile: 0.51,
			expected:   3,
			shouldErr:  false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := test.pmf.Percentile(test.percentile)

			if test.shouldErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			assert.Equal(t, test.expected, res)
		})
	}
}

func TestMaximumLikelihood(t *testing.T) {
	tests := map[string]struct {
		elements  []*PmfElement
		expected  float64
		shouldErr bool
	}{
		"empty Pmf": {
			elements:  []*PmfElement{},
			expected:  0,
			shouldErr: true,
		},
		"single elememt in Pmf": {
			elements: []*PmfElement{
				NewPmfElement(15, 1),
			},
			expected:  15,
			shouldErr: false,
		},
		"multiple elememts in Pmf": {
			elements: []*PmfElement{
				NewPmfElement(1, 0.1),
				NewPmfElement(2, 0.7),
				NewPmfElement(3, 0.2),
			},
			expected:  2,
			shouldErr: false,
		},
		"multiple elememts in Pmf with zero probabilities": {
			elements: []*PmfElement{
				NewPmfElement(1, 0),
				NewPmfElement(2, 0),
				NewPmfElement(3, 0),
			},
			expected:  0,
			shouldErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := setupPmf(test.elements)

			mle, err := p.MaximumLikelihood()

			if test.shouldErr {
				require.NotNil(t, err)
				return
			}
			assert.Equal(t, test.expected, mle)
		})
	}
}
