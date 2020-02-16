package prob

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			assert.Equal(t, 1.0, getSum(s.prob))
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

func TestSuiteUpdateSet(t *testing.T) {
	t.Run("suite UpdateSet", func(t *testing.T) {

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

		s.UpdateSet(obs)

		for elem, prob := range expectedPosterior {
			if prob == 0 {
				assert.Equal(t, 0.0, s.prob[elem])
			} else {
				assert.InEpsilon(t, prob, s.prob[elem], float64EqualTol)
			}
		}
	})
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
			assert.Equal(t, 1.0, getSum(s.pmf.prob))
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

func TestNamedSuiteUpdateSet(t *testing.T) {
	t.Run("named suite UpdateSet", func(t *testing.T) {

		obs := []NamedSuiteObservation{
			&namedSuiteTestObservation{"foo"},
			&namedSuiteTestObservation{"bar"},
		}

		expectedPosterior := map[string]float64{
			"a": 0.1875 / 0.4275,
			"b": 0.24 / 0.4275,
		}

		s := NewNamedSuite(namedSuiteUpdateHypos...)

		s.UpdateSet(obs)

		for elem, prob := range expectedPosterior {
			if prob == 0 {
				assert.Equal(t, 0.0, s.Prob(elem))
			} else {
				assert.InEpsilon(t, prob, s.Prob(elem), float64EqualTol)
			}
		}
	})
}
