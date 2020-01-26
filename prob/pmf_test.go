package prob

import (
	"testing"

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
				assert.Contains(t, p.prob, elem.Val)
				assert.Equal(t, elem.Prob, p.prob[elem.Val])
			}
			assert.Equal(t, test.expectedSum, p.sum)
		})
	}
}

func TestNormalize(t *testing.T) {
	tests := map[string]struct {
		elements     []*PmfElement
		sum          float64
		expectedProb map[float64]float64
		expectedSum  float64
	}{
		"empty Pmf": {
			elements:     []*PmfElement{},
			sum:          0,
			expectedProb: map[float64]float64{},
			expectedSum:  0,
		},
		"Pmf with single element": {
			elements: []*PmfElement{
				NewPmfElement(1, 100),
			},
			sum:          100,
			expectedProb: map[float64]float64{1: 1},
			expectedSum:  1,
		},
		"Pmf with multiple elements": {
			elements: []*PmfElement{
				NewPmfElement(1, 1),
				NewPmfElement(2, 1),
				NewPmfElement(3, 1),
				NewPmfElement(4, 1),
			},
			sum: 4,
			expectedProb: map[float64]float64{
				1: 0.25,
				2: 0.25,
				3: 0.25,
				4: 0.25,
			},
			expectedSum: 1,
		},
		"Pmf with multiple elements and sum 0": {
			elements: []*PmfElement{
				NewPmfElement(1, 1),
				NewPmfElement(2, 1),
				NewPmfElement(3, 1),
				NewPmfElement(4, 1),
			},
			sum: 0,
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
			p.sum = test.sum

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
			origProb, found := p.prob[test.elem]

			p.Mult(test.elem, test.multVal)
			// test probability of specified element correctly multiplied
			if found {
				assert.Equal(t, origProb*test.multVal, p.prob[test.elem])
			}
			// test other probabilities are unchanged
			for _, element := range test.elements {
				if element.Val == test.elem {
					continue
				}
				assert.Equal(t, element.Prob, p.prob[element.Val])
			}
			assert.Equal(t, test.expectedSum, p.sum)
		})
	}
}
