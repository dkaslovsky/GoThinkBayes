package exercises

import (
	"fmt"

	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// Locomotive Problem:
// A railroad numbers its locomotives in order 1..N.
// One day you see a locomotive with the number 60.
// Estimate how many loco- motives the railroad has.

func uniform(low, high int) (elems []*prob.NumericPmfElement) {
	for i := low; i <= high; i++ {
		elems = append(elems, prob.NewNumericPmfElement(float64(i), 1))
	}
	return elems
}

// Locomotive runs the locomotive problem
func Locomotive() {
	uniformBounds := []struct {
		low  int
		high int
	}{
		{
			low:  1,
			high: 500,
		},
		{
			low:  1,
			high: 1000,
		},
		{
			low:  1,
			high: 2000,
		},
	}

	for _, uBound := range uniformBounds {
		hypos := uniform(uBound.low, uBound.high)
		l := prob.NewNumericSuite(hypos...)

		// observations (likelihood function) are the same as that of the dice problem
		obs := []prob.NumericSuiteObservation{
			&diceObservation{60},
			&diceObservation{30},
			&diceObservation{90},
		}
		l.MultiUpdate(obs)

		fmt.Printf("Upper bound: %d, Posterior mean: %0.2f\n", uBound.high, l.Mean())
	}
}
