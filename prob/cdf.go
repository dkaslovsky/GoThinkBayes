package prob

import (
	"fmt"
	"sort"
)

type Cdf struct {
	elemsToIdx map[float64]int
	idxToelems map[int]float64
	prob       []float64
}

func NewCdf(p map[float64]float64) *Cdf {
	elems := make(map[float64]int)
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

func (c *Cdf) Percentile(p float64) (val float64, err error) {
	if p < 0 || p > 1 {
		return val, fmt.Errorf("invalid percentile [%f]", p)
	}
	// TODO: use bisection?
	for i, prob := range c.prob {
		if prob >= p {
			return c.idxToelems[i], nil
		}
	}
	return val, fmt.Errorf("unable to compute percentile [%f]", p)
}

func sortKeys(p map[float64]float64) []float64 {
	keys := make([]float64, 0, len(p))
	for elem := range p {
		keys = append(keys, elem)
	}
	sort.Float64s(keys)
	return keys
}

func reverseMap(m map[float64]int) map[int]float64 {
	revM := make(map[int]float64)
	for key, val := range m {
		revM[val] = key
	}
	return revM
}
