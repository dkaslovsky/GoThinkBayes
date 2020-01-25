package exercises

import (
	"fmt"
	"math"

	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// Locomotive Problem:
// A railroad numbers its locomotives in order 1..N.
// One day you see a locomotive with the number 60.
// Estimate how many loco- motives the railroad has.

type bound struct {
	low  int
	high int
}

var bounds = []bound{
	bound{
		low:  1,
		high: 500,
	},
	bound{
		low:  1,
		high: 1000,
	},
	bound{
		low:  1,
		high: 2000,
	},
}

func uniform(b bound) (elems []*prob.PmfElement) {
	for i := b.low; i <= b.high; i++ {
		elems = append(elems, prob.NewPmfElement(float64(i), 1))
	}
	return elems
}

func powerLaw(b bound, alpha float64) (elems []*prob.PmfElement) {
	a := -alpha
	for i := b.low; i <= b.high; i++ {
		n := float64(i)
		elems = append(elems, prob.NewPmfElement(n, math.Pow(n, a)))
	}
	return elems
}

// locomotive observations (likelihood function) are the same as that of the dice problem
type locomotiveObservation struct {
	*diceObservation
}

func newLocomotiveObservation(val int) *locomotiveObservation {
	return &locomotiveObservation{
		&diceObservation{float64(val)},
	}
}

func locomototiveUniformSingleObservation() {
	fmt.Println("Uniform Single Observation")

	for _, bound := range bounds {
		hypos := uniform(bound)
		l := prob.NewSuite(hypos...)
		l.Update(newLocomotiveObservation(60))

		fmt.Printf("Upper bound: %d, Posterior mean: %0.2f\n", bound.high, l.Mean())
	}
}

func locomototiveUniformMultipleObservations() {
	fmt.Println("Uniform Multiple Observation")

	for _, bound := range bounds {
		hypos := uniform(bound)
		l := prob.NewSuite(hypos...)

		obs := []prob.SuiteObservation{
			newLocomotiveObservation(60),
			newLocomotiveObservation(30),
			newLocomotiveObservation(90),
		}
		l.MultiUpdate(obs)

		fmt.Printf("Upper bound: %d, Posterior mean: %0.2f\n", bound.high, l.Mean())
	}
}

func locomotivePowerLawSingleObservation() {
	fmt.Println("Power Law Single Observation")

	alpha := 1.0
	for _, bound := range bounds {
		hypos := powerLaw(bound, alpha)
		l := prob.NewSuite(hypos...)
		l.Update(newLocomotiveObservation(60))

		fmt.Printf("Upper bound: %d, Posterior mean: %0.2f\n", bound.high, l.Mean())
	}
}

func locomotivePowerLawMultipleObservation() {
	fmt.Println("Power Law Multiple Observation")

	alpha := 1.0
	for _, bound := range bounds {
		hypos := powerLaw(bound, alpha)
		l := prob.NewSuite(hypos...)

		obs := []prob.SuiteObservation{
			newLocomotiveObservation(60),
			newLocomotiveObservation(30),
			newLocomotiveObservation(90),
		}
		l.MultiUpdate(obs)

		fmt.Printf("Upper bound: %d, Posterior mean: %0.2f ", bound.high, l.Mean())

		cdf := l.MakeCdf()
		lower := cdf.Percentile(0.05)
		upper := cdf.Percentile(0.95)
		fmt.Printf("[90%% Credible Interval: (%f, %f)]\n", lower, upper)
	}

}

// Locomotive runs the locomotive problem
func Locomotive() {
	locomototiveUniformSingleObservation()
	locomototiveUniformMultipleObservations()
	locomotivePowerLawSingleObservation()
	locomotivePowerLawMultipleObservation()
}
