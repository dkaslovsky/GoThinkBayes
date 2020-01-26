package prob

import (
	"sort"
)

// Cdf is a cumulative distribution function
type Cdf struct {
	elemsToIdx map[float64]int
	idxToelems map[int]float64
	prob       []float64
}

// NewCdf creates a new Cdf
func NewCdf(p map[float64]float64) *Cdf {
	elems := map[float64]int{}
	prob := []float64{}

	cumsum := 0.0
	for i, key := range sortKeys(p) {
		elems[key] = i
		cumsum += p[key]
		prob = append(prob, cumsum)
	}

	return &Cdf{
		elemsToIdx: elems,
		idxToelems: reverseMap(elems),
		prob:       prob,
	}
}

// Percentile computes the specified percentile of the distribution
func (c *Cdf) Percentile(p float64) float64 {
	if p < 0 {
		return c.idxToelems[0]
	}
	if p > 1 {
		return c.idxToelems[len(c.idxToelems)-1]
	}
	i := sort.Search(len(c.prob), func(i int) bool {
		return c.prob[i] >= p
	})
	return c.idxToelems[i]
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
