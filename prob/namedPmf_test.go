package prob

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			p := NewNamedPmf()
			for _, elem := range test.elements {
				p.Set(elem)
			}

			for _, elem := range test.elements {
				assert.Contains(t, p.nameToIdx, elem.Name)
				assert.Contains(t, p.pmf.prob, p.nameToIdx[elem.Name])
				assert.Equal(t, elem.Prob, p.pmf.prob[p.nameToIdx[elem.Name]])
			}
		})
	}
}

func TestNamedMult(t *testing.T) {
	tests := map[string]struct {
		elements    []*NamedPmfElement
		elem        string
		multVal     float64
		expectedSum float64
	}{
		"element not in Pmf": {
			elements: []*NamedPmfElement{
				NewNamedPmfElement("a", 0.5),
				NewNamedPmfElement("b", 0.5),
			},
			elem:        "c",
			multVal:     0.5,
			expectedSum: 1,
		},
		"element in Pmf": {
			elements: []*NamedPmfElement{
				NewNamedPmfElement("a", 0.5),
				NewNamedPmfElement("b", 0.5),
			},
			elem:        "a",
			multVal:     0.5,
			expectedSum: 0.75,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := NewNamedPmf()
			for _, element := range test.elements {
				p.Set(element)
			}

			var origProb float64
			idx, found := p.nameToIdx[test.elem]
			origProb, foundInProb := p.pmf.prob[idx]

			p.Mult(test.elem, test.multVal)
			// test probability of specified element correctly multiplied
			if found {
				assert.True(t, foundInProb)
				assert.Equal(t, origProb*test.multVal, p.pmf.prob[idx])
			}
			// test other probabilities are unchanged
			for _, element := range test.elements {
				if element.Name == test.elem {
					continue
				}
				assert.Equal(t, element.Prob, p.pmf.prob[p.nameToIdx[element.Name]])
			}
			assert.Equal(t, test.expectedSum, p.pmf.sum)
		})
	}
}

func TestNamedProb(t *testing.T) {
	tests := map[string]struct {
		elements     []*NamedPmfElement
		elem         string
		expectedProb float64
	}{
		"elememt in Pmf": {
			elements: []*NamedPmfElement{
				NewNamedPmfElement("a", 0.25),
				NewNamedPmfElement("b", 0.75),
			},
			elem:         "a",
			expectedProb: 0.25,
		},
		"elememt not in Pmf": {
			elements: []*NamedPmfElement{
				NewNamedPmfElement("a", 0.25),
				NewNamedPmfElement("b", 0.75),
			},
			elem:         "c",
			expectedProb: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := NewNamedPmf()
			for _, element := range test.elements {
				p.Set(element)
			}

			prob := p.Prob(test.elem)
			assert.Equal(t, test.expectedProb, prob)
		})
	}
}

func TestNewNamedSuite(t *testing.T) {
	tests := map[string]struct {
		elements []*NamedPmfElement
	}{
		"single element": {
			elements: []*NamedPmfElement{
				NewNamedPmfElement("a", 100),
			},
		},
		"multiple elements": {
			elements: []*NamedPmfElement{
				NewNamedPmfElement("a", 100),
				NewNamedPmfElement("b", 200),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewNamedSuite(test.elements...)
			for _, elem := range test.elements {
				assert.Contains(t, s.nameToIdx, elem.Name)
				assert.Contains(t, s.pmf.prob, s.nameToIdx[elem.Name])
			}
			assert.Equal(t, 1.0, s.pmf.sum)
		})
	}
}

type namedSuiteTestObservation struct {
	name string
}

func (o *namedSuiteTestObservation) GetLikelihood(hypo string) float64 {
	if hypo == "a" {
		if o.name == "foo" {
			return 0.25
		}
		if o.name == "bar" {
			return 0.75
		}
	}
	if hypo == "b" {
		if o.name == "foo" {
			return 0.4
		}
		if o.name == "bar" {
			return 0.6
		}
	}
	return 0
}

var namedSuiteUpdateHypos = []*NamedPmfElement{
	NewNamedPmfElement("a", 1),
	NewNamedPmfElement("b", 1),
}

func TestNamedSuiteUpdate(t *testing.T) {
	t.Run("named suite update", func(t *testing.T) {

		ob := &namedSuiteTestObservation{"foo"}
		expectedPosterior := map[string]float64{
			"a": 0.25 / 0.65,
			"b": 0.4 / 0.65,
		}

		s := NewNamedSuite(namedSuiteUpdateHypos...)
		s.Update(ob)
		for elem, prob := range expectedPosterior {
			if prob == 0 {
				assert.Equal(t, 0.0, s.Prob(elem))
			} else {
				assert.InEpsilon(t, prob, s.Prob(elem), float64EqualTol)
			}
		}
	})
}

func TestNamedSuiteMultiUpdate(t *testing.T) {
	t.Run("named suite multiupdate", func(t *testing.T) {

		obs := []NamedSuiteObservation{
			&namedSuiteTestObservation{"foo"},
			&namedSuiteTestObservation{"bar"},
		}

		expectedPosterior := map[string]float64{
			"a": 0.1875 / 0.4275,
			"b": 0.24 / 0.4275,
		}

		s := NewNamedSuite(namedSuiteUpdateHypos...)
		s.MultiUpdate(obs)
		for elem, prob := range expectedPosterior {
			if prob == 0 {
				assert.Equal(t, 0.0, s.Prob(elem))
			} else {
				assert.InEpsilon(t, prob, s.Prob(elem), float64EqualTol)
			}
		}
	})
}
