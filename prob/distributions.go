package prob

import (
	"fmt"
	"math"
)

// TODO: consider noninteger bounds (floating point precision problem for keys underlying map...)

// Bound contains the bounds of a distribution
type Bound struct {
	Low  int
	High int
}

// NewBound constructs a new Bound
func NewBound(low, high int) *Bound {
	return &Bound{
		Low:  low,
		High: high,
	}
}

// Uniform generates PmfElements representing a uniform distribution
func Uniform(b *Bound) (elems []*PmfElement) {
	for i := b.Low; i <= b.High; i++ {
		elems = append(elems, NewPmfElement(float64(i), 1))
	}
	return elems
}

// Triangle generates PmfElements representing a triangle distribution
func Triangle(b *Bound) (elems []*PmfElement) {
	l := b.High - b.Low + 1 // interval length
	mdpt := l / 2
	isEven := math.Mod(float64(l), 2) == 0

	var prob int
	for i := 0; i < l; i++ {
		switch {
		case i < mdpt:
			prob = i
		case i > mdpt:
			prob = l - i - 1
		case i == mdpt && isEven:
			prob = i - 1
		default: // i == mdpt && !isEven:
			prob = i
		}
		elems = append(elems, NewPmfElement(float64(b.Low+i), float64(prob)))
	}
	return elems
}

// PowerLaw generates PmfElements representing a power law distribution
func PowerLaw(b *Bound, alpha float64) (elems []*PmfElement) {
	a := -alpha
	for i := b.Low; i <= b.High; i++ {
		n := float64(i)
		elems = append(elems, NewPmfElement(n, math.Pow(n, a)))
	}
	return elems
}

type percentileGetter interface {
	Percentile(float64) (float64, error)
}

// CredibleInterval computes the credible interval of specified length
func CredibleInterval(p percentileGetter, l float64) (lower float64, upper float64, err error) {
	if l <= 0 || l > 100 {
		return lower, upper, fmt.Errorf("cannot compute CI of length [%f]", l)
	}

	lowerP, upperP := getCredibleIntervalPercentiles(l)

	lower, err = p.Percentile(lowerP)
	if err != nil {
		return lower, upper, fmt.Errorf(
			"error computing credible interval of length [%f]: %v", l, err,
		)
	}
	upper, err = p.Percentile(upperP)
	if err != nil {
		return lower, upper, fmt.Errorf(
			"error computing credible interval of length [%f]: %v", l, err,
		)
	}
	return lower, upper, nil
}

func getCredibleIntervalPercentiles(l float64) (lower float64, upper float64) {
	lowerP := (100.0 - l) / 2.0 / 100
	upperP := 1.0 - lowerP
	return lowerP, upperP
}

// Beta is a Beta distribution
type Beta struct {
	alpha float64
	beta  float64
}

// NewBeta creates a new Beta
func NewBeta(alpha float64, beta float64) (b *Beta, err error) {
	if alpha <= 0 || beta <= 0 {
		return b, fmt.Errorf("cannot initialize Beta distribution with non-positive parameter(s)")
	}
	b = &Beta{
		alpha: alpha,
		beta:  beta,
	}
	return b, nil
}

// Update updates the parameters of a Beta distribution
func (b *Beta) Update(nHeads float64, nTails float64) {
	b.alpha += nHeads
	b.beta += nTails
}

// Mean computes the mean of a Beta distribution
func (b *Beta) Mean() float64 {
	return b.alpha / (b.alpha + b.beta)
}

// EvalPdf computes the probability associated with a value for a Beta distribution
func (b *Beta) EvalPdf(x float64) float64 {
	return math.Pow(x, b.alpha-1) * math.Pow(1-x, b.beta-1)
}

// MakePmf returns a Pmf representing a discretized Beta distribution
func (b *Beta) MakePmf(nPoints int) *Pmf {
	p := NewPmf()
	denom := float64(nPoints - 1)

	var i float64
	for i = 0; i <= denom; i++ {
		val := i / denom
		p.Set(NewPmfElement(val, b.EvalPdf(val)))
	}
	return p
}
