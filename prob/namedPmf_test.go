package prob

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupNamedPmf(elems []*NamedPmfElement) *NamedPmf {
	p := NewNamedPmf()
	for _, elem := range elems {
		p.Set(elem)
	}
	return p
}

func TestNewNamedPmf(t *testing.T) {
	t.Run("new Pmf", func(t *testing.T) {
		p := NewNamedPmf()

		assert.Empty(t, p.nameToIdx)
		assert.Equal(t, 0.0, p.nextIdx)
	})
}

func TestNamedSet(t *testing.T) {
	tests := map[string]struct {
		elements []*NamedPmfElement
	}{
		"single element": {
			elements: []*NamedPmfElement{NewNamedPmfElement("a", 1)},
		},
		"multiple elements": {
			elements: []*NamedPmfElement{
				NewNamedPmfElement("a", 1),
				NewNamedPmfElement("b", 10.5),
				NewNamedPmfElement("c", 1.6),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := setupNamedPmf(test.elements)

			for _, elem := range test.elements {
				require.Contains(t, p.nameToIdx, elem.Name)
				require.Contains(t, p.pmf.prob, p.nameToIdx[elem.Name])
				assert.Equal(t, elem.Prob, p.pmf.prob[p.nameToIdx[elem.Name]])
			}
		})
	}
}

func TestNamedMult(t *testing.T) {
	tests := map[string]struct {
		elements   []*NamedPmfElement
		val        string
		multFactor float64
	}{
		"element not in Pmf": {
			elements: []*NamedPmfElement{
				NewNamedPmfElement("a", 0.5),
				NewNamedPmfElement("b", 0.5),
			},
			val:        "c",
			multFactor: 0.5,
		},
		"element in Pmf": {
			elements: []*NamedPmfElement{
				NewNamedPmfElement("a", 0.5),
				NewNamedPmfElement("b", 0.5),
			},
			val:        "a",
			multFactor: 0.5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := setupNamedPmf(test.elements)

			idx, found := p.nameToIdx[test.val]
			origProb, foundInProb := p.pmf.prob[idx]

			p.Mult(test.val, test.multFactor)

			// test probability of specified element correctly multiplied
			if found {
				require.True(t, foundInProb)
				assert.Equal(t, origProb*test.multFactor, p.pmf.prob[idx])
			}
			// test other probabilities are unchanged
			for _, element := range test.elements {
				if element.Name == test.val {
					continue
				}
				assert.Equal(t, element.Prob, p.pmf.prob[p.nameToIdx[element.Name]])
			}
		})
	}
}

func TestNamedProb(t *testing.T) {
	tests := map[string]struct {
		elements     []*NamedPmfElement
		val          string
		expectedProb float64
	}{
		"elememt in Pmf": {
			elements: []*NamedPmfElement{
				NewNamedPmfElement("a", 0.25),
				NewNamedPmfElement("b", 0.75),
			},
			val:          "a",
			expectedProb: 0.25,
		},
		"elememt not in Pmf": {
			elements: []*NamedPmfElement{
				NewNamedPmfElement("a", 0.25),
				NewNamedPmfElement("b", 0.75),
			},
			val:          "c",
			expectedProb: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := setupNamedPmf(test.elements)

			prob := p.Prob(test.val)

			assert.Equal(t, test.expectedProb, prob)
		})
	}
}

func TestNamedMaximumLikelihood(t *testing.T) {
	tests := map[string]struct {
		elements  []*NamedPmfElement
		expected  string
		shouldErr bool
	}{
		"empty Pmf": {
			elements:  []*NamedPmfElement{},
			expected:  "",
			shouldErr: true,
		},
		"single elememt in Pmf": {
			elements: []*NamedPmfElement{
				NewNamedPmfElement("a", 1),
			},
			expected:  "a",
			shouldErr: false,
		},
		"multiple elememts in Pmf": {
			elements: []*NamedPmfElement{
				NewNamedPmfElement("a", 0.1),
				NewNamedPmfElement("b", 0.7),
				NewNamedPmfElement("c", 0.2),
			},
			expected:  "b",
			shouldErr: false,
		},
		"multiple elememts in Pmf with zero probabilities": {
			elements: []*NamedPmfElement{
				NewNamedPmfElement("a", 0),
				NewNamedPmfElement("b", 0),
				NewNamedPmfElement("c", 0),
			},
			expected:  "",
			shouldErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := setupNamedPmf(test.elements)

			mle, err := p.MaximumLikelihood()

			if test.shouldErr {
				require.NotNil(t, err)
				return
			}
			assert.Equal(t, test.expected, mle)
		})
	}
}
