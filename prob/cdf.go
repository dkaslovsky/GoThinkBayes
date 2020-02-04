package prob

import (
	"fmt"
	"sort"
)

// Cdf is a cumulative distribution function
type Cdf struct {
	elemsToIdx map[float64]int
	idxToelems map[int]float64
	prob       []float64
}

// NewCdf creates a new Cdf
func NewCdf(p map[float64]float64) (c *Cdf, err error) {
	if len(p) == 0 {
		return c, fmt.Errorf("cannot compute cdf from empty input map")
	}

	elems := map[float64]int{}
	prob := []float64{}

	cumsum := 0.0
	for i, key := range sortKeys(p) {
		elems[key] = i
		cumsum += p[key]
		prob = append(prob, cumsum)
	}

	// cumulative sum of input probabilities must be 1 for a proper cdf
	if !almostEqual(cumsum, 1.0, float64EqualTol) {
		return c, fmt.Errorf("cumulative sum of probabilities [%f] not equal to 1.0", cumsum)
	}

	c = &Cdf{
		elemsToIdx: elems,
		idxToelems: reverseMap(elems),
		prob:       prob,
	}
	return c, nil
}

// Percentile computes the specified percentile of the distribution
func (c *Cdf) Percentile(p float64) (float64, error) {
	if p < 0 || p > 1 {
		return 0, fmt.Errorf("percentile [%f] is outside of required range [0, 1]", p)
	}
	i := sort.Search(len(c.prob), func(i int) bool {
		return c.prob[i] >= p
	})
	return c.idxToelems[i], nil
}

func sortKeys(p map[float64]float64) []float64 {
	keys := make([]float64, 0, len(p))
	for elem := range p {
		keys = append(keys, elem)
	}
	sort.Float64s(keys)
	return keys
}

// result of reverseMap is unique only if values of input map are unique;
// intended use is to reverse a map constructed with increasing and therefore unique values
func reverseMap(m map[float64]int) map[int]float64 {
	revM := map[int]float64{}
	for key, val := range m {
		revM[val] = key
	}
	return revM
}
