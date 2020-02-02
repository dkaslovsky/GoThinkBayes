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
		s := prob.NewSuite(hypos...)
		s.Update(newLocomotiveObservation(60))

		fmt.Printf("Upper bound: %d, Posterior mean: %0.2f\n", bound.high, s.Mean())
	}
}

func locomototiveUniformMultipleObservations() {
	fmt.Println("Uniform Multiple Observation")

	for _, bound := range bounds {
		hypos := uniform(bound)
		s := prob.NewSuite(hypos...)

		obs := []prob.SuiteObservation{
			newLocomotiveObservation(60),
			newLocomotiveObservation(30),
			newLocomotiveObservation(90),
		}
		s.MultiUpdate(obs)

		fmt.Printf("Upper bound: %d, Posterior mean: %0.2f\n", bound.high, s.Mean())
	}
}

func locomotivePowerLawSingleObservation() {
	fmt.Println("Power Law Single Observation")

	alpha := 1.0
	for _, bound := range bounds {
		hypos := powerLaw(bound, alpha)
		s := prob.NewSuite(hypos...)
		s.Update(newLocomotiveObservation(60))

		fmt.Printf("Upper bound: %d, Posterior mean: %0.2f\n", bound.high, s.Mean())
	}
}

func locomotivePowerLawMultipleObservation() {
	fmt.Println("Power Law Multiple Observation")

	alpha := 1.0
	for _, bound := range bounds {
		hypos := powerLaw(bound, alpha)
		s := prob.NewSuite(hypos...)

		obs := []prob.SuiteObservation{
			newLocomotiveObservation(60),
			newLocomotiveObservation(30),
			newLocomotiveObservation(90),
		}
		s.MultiUpdate(obs)

		fmt.Printf("Upper bound: %d, Posterior mean: %0.2f ", bound.high, s.Mean())

		// compute 90% credible interval from cdf
		cdf, err := s.MakeCdf()
		if err != nil {
			fmt.Printf("[Could not compute 90%% Credible Interval due to error [%v]]\n", err)
			continue
		}
		lower, err := cdf.Percentile(0.05)
		if err != nil {
			fmt.Printf("[Could not compute 90%% Credible Interval due to error [%v]]\n", err)
			continue
		}
		upper, err := cdf.Percentile(0.95)
		if err != nil {
			fmt.Printf("[Could not compute 90%% Credible Interval due to error [%v]]\n", err)
			continue
		}
		fmt.Printf("[90%% Credible Interval: (%0.2f, %0.2f)]\n", lower, upper)
	}

}

// Locomotive runs the locomotive problem
func Locomotive() {
	locomototiveUniformSingleObservation()
	locomototiveUniformMultipleObservations()
	locomotivePowerLawSingleObservation()
	locomotivePowerLawMultipleObservation()
}
