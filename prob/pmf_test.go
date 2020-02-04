package prob

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestNewPmf(t *testing.T) {
	t.Run("new Pmf", func(t *testing.T) {
		p := NewPmf()
		assert.Empty(t, p.prob)
		assert.Equal(t, 0.0, p.sum)
	})
}

func TestSet(t *testing.T) {
	tests := map[string]struct {
		elements    []*PmfElement
		expectedSum float64
	}{
		"single element": {
			elements:    []*PmfElement{NewPmfElement(1, 1)},
			expectedSum: 1,
		},
		"multiple elements": {
			elements: []*PmfElement{
				NewPmfElement(1, 1),
				NewPmfElement(2.1, 10.5),
				NewPmfElement(3, 1.6),
			},
			expectedSum: 13.1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := NewPmf()
			for _, elem := range test.elements {
				p.Set(elem)
			}

			for _, elem := range test.elements {
				require.Contains(t, p.prob, elem.Val)
				assert.Equal(t, elem.Prob, p.prob[elem.Val])
			}
			assert.Equal(t, test.expectedSum, p.sum)
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
			p := NewPmf()
			for _, element := range test.elements {
				p.Set(element)
			}

			p.Normalize()
			for elem, prob := range p.prob {
				assert.Equal(t, test.expectedProb[elem], prob)
			}
			assert.Equal(t, test.expectedSum, p.sum)
		})
	}
}

func TestMult(t *testing.T) {
	tests := map[string]struct {
		elements    []*PmfElement
		elem        float64
		multVal     float64
		expectedSum float64
	}{
		"element not in Pmf": {
			elements: []*PmfElement{
				NewPmfElement(1, 0.5),
				NewPmfElement(2, 0.5),
			},
			elem:        3,
			multVal:     0.5,
			expectedSum: 1,
		},
		"element in Pmf": {
			elements: []*PmfElement{
				NewPmfElement(1, 0.5),
				NewPmfElement(2, 0.5),
			},
			elem:        1,
			multVal:     0.5,
			expectedSum: 0.75,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := NewPmf()
			for _, element := range test.elements {
				p.Set(element)
			}
			// store original probability before mutating
			origProb, found := p.prob[test.elem]

			p.Mult(test.elem, test.multVal)

			for _, element := range test.elements {
				// test probability of specified element correctly multiplied
				if element.Val == test.elem && found {
					assert.Equal(t, origProb*test.multVal, p.prob[test.elem])
					continue
				}
				// test other probabilities are unchanged
				assert.Equal(t, element.Prob, p.prob[element.Val])
			}
			assert.Equal(t, test.expectedSum, p.sum)
		})
	}
}

func TestProb(t *testing.T) {
	tests := map[string]struct {
		elements     []*PmfElement
		elem         float64
		expectedProb float64
	}{
		"elememt in Pmf": {
			elements: []*PmfElement{
				NewPmfElement(1, 0.25),
				NewPmfElement(2, 0.75),
			},
			elem:         1,
			expectedProb: 0.25,
		},
		"elememt not in Pmf": {
			elements: []*PmfElement{
				NewPmfElement(1, 0.25),
				NewPmfElement(2, 0.75),
			},
			elem:         3,
			expectedProb: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := NewPmf()
			for _, element := range test.elements {
				p.Set(element)
			}

			prob := p.Prob(test.elem)
			assert.Equal(t, test.expectedProb, prob)
		})
	}
}

func TestMean(t *testing.T) {
	tests := map[string]struct {
		elements     []*PmfElement
		expectedMean float64
	}{
		"single elememt in Pmf": {
			elements: []*PmfElement{
				NewPmfElement(1, 1),
			},
			expectedMean: 1,
		},
		"multiple elememts in Pmf": {
			elements: []*PmfElement{
				NewPmfElement(1, 0.25),
				NewPmfElement(2, 0.75),
			},
			expectedMean: 1.75,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := NewPmf()
			for _, element := range test.elements {
				p.Set(element)
			}

			val := p.Mean()
			assert.Equal(t, test.expectedMean, val)
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
				sum:  1,
			},
			percentile: -0.5,
			expected:   0,
			shouldErr:  true,
		},
		"percentile greater than 1": {
			pmf: &Pmf{
				prob: map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
				sum:  1,
			},
			percentile: 5,
			expected:   0,
			shouldErr:  true,
		},
		"unnormalized pmf": {
			pmf: &Pmf{
				prob: map[float64]float64{1: 0.02, 2: 0.03, 3: 0.04, 4: 0.01},
				sum:  0.1,
			},
			percentile: 0.5,
			expected:   0,
			shouldErr:  true,
		},
		"unnormalized pmf with sum = 1": {
			pmf: &Pmf{
				prob: map[float64]float64{1: 0.02, 2: 0.03, 3: 0.04, 4: 0.01},
				sum:  1,
			},
			percentile: 0.5,
			expected:   0,
			shouldErr:  true,
		},
		"percentile 0": {
			pmf: &Pmf{
				prob: map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
				sum:  1,
			},
			percentile: 0,
			expected:   1,
			shouldErr:  false,
		},
		"percentile 1": {
			pmf: &Pmf{
				prob: map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
				sum:  1,
			},
			percentile: 1,
			expected:   4,
			shouldErr:  false,
		},
		"percentile 0.5": {
			pmf: &Pmf{
				prob: map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
				sum:  1,
			},
			percentile: 0.5,
			expected:   2,
			shouldErr:  false,
		},
		"percentile 0.51": {
			pmf: &Pmf{
				prob: map[float64]float64{1: 0.2, 2: 0.3, 3: 0.4, 4: 0.1},
				sum:  1,
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

func TestNewSuite(t *testing.T) {
	tests := map[string]struct {
		elements []*PmfElement
	}{
		"single element": {
			elements: []*PmfElement{
				NewPmfElement(1, 100),
			},
		},
		"multiple elements": {
			elements: []*PmfElement{
				NewPmfElement(1, 100),
				NewPmfElement(2, 200),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewSuite(test.elements...)
			for _, elem := range test.elements {
				assert.Contains(t, s.prob, elem.Val)
			}
			assert.Equal(t, 1.0, s.sum)
		})
	}
}

type suiteTestObservation struct {
	val float64
}

func (o *suiteTestObservation) GetLikelihood(hypo float64) float64 {
	if hypo < o.val {
		return 0
	}
	return 1 / hypo
}

var suiteUpdateHypos = []*PmfElement{
	NewPmfElement(2, 1),
	NewPmfElement(3, 1),
	NewPmfElement(4, 1),
	NewPmfElement(5, 1),
}

func TestSuiteUpdate(t *testing.T) {
	t.Run("suite update", func(t *testing.T) {

		ob := &suiteTestObservation{4}
		expectedPosterior := map[float64]float64{
			2: 0.0,
			3: 0.0,
			4: 0.25 / 0.45,
			5: 0.2 / 0.45,
		}

		s := NewSuite(suiteUpdateHypos...)
		s.Update(ob)
		for elem, prob := range expectedPosterior {
			if prob == 0 {
				assert.Equal(t, 0.0, s.prob[elem])
			} else {
				assert.InEpsilon(t, prob, s.prob[elem], float64EqualTol)
			}
		}
	})
}

func TestSuiteMultiUpdate(t *testing.T) {
	t.Run("suite multiupdate", func(t *testing.T) {

		obs := []SuiteObservation{
			&suiteTestObservation{4},
			&suiteTestObservation{4},
		}
		expectedPosterior := map[float64]float64{
			2: 0.0,
			3: 0.0,
			4: 0.0625 / 0.1025,
			5: 0.04 / 0.1025,
		}

		s := NewSuite(suiteUpdateHypos...)
		s.MultiUpdate(obs)
		for elem, prob := range expectedPosterior {
			if prob == 0 {
				assert.Equal(t, 0.0, s.prob[elem])
			} else {
				assert.InEpsilon(t, prob, s.prob[elem], float64EqualTol)
			}
		}
	})
}
