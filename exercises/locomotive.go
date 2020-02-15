package exercises

import (
	"fmt"

	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// Locomotive Problem:
// A railroad numbers its locomotives in order 1..N.
// One day you see a locomotive with the number 60.
// Estimate how many loco- motives the railroad has.

var bounds = []*prob.Bound{
	prob.NewBound(1, 500),
	prob.NewBound(1, 1000),
	prob.NewBound(1, 2000),
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
		hypos := prob.Uniform(bound)
		s := prob.NewSuite(hypos...)
		s.Update(newLocomotiveObservation(60))

		mean, err := s.Mean()
		if err != nil {
			fmt.Printf("Unable to compute mean due to error [%v]", err)
			continue
		}
		fmt.Printf("Upper bound: %d, Posterior mean: %0.2f\n", bound.High, mean)
	}
}

func locomototiveUniformMultipleObservations() {
	fmt.Println("Uniform Multiple Observation")

	for _, bound := range bounds {
		hypos := prob.Uniform(bound)
		s := prob.NewSuite(hypos...)

		obs := []prob.SuiteObservation{
			newLocomotiveObservation(60),
			newLocomotiveObservation(30),
			newLocomotiveObservation(90),
		}
		s.MultiUpdate(obs)

		mean, err := s.Mean()
		if err != nil {
			fmt.Printf("Unable to compute mean due to error [%v]", err)
			continue
		}
		fmt.Printf("Upper bound: %d, Posterior mean: %0.2f\n", bound.High, mean)
	}
}

func locomotivePowerLawSingleObservation() {
	fmt.Println("Power Law Single Observation")

	alpha := 1.0
	for _, bound := range bounds {
		hypos := prob.PowerLaw(bound, alpha)
		s := prob.NewSuite(hypos...)
		s.Update(newLocomotiveObservation(60))

		mean, err := s.Mean()
		if err != nil {
			fmt.Printf("Unable to compute mean due to error [%v]", err)
			continue
		}
		fmt.Printf("Upper bound: %d, Posterior mean: %0.2f\n", bound.High, mean)
	}
}

func locomotivePowerLawMultipleObservation() {
	fmt.Println("Power Law Multiple Observation")

	alpha := 1.0
	for _, bound := range bounds {
		hypos := prob.PowerLaw(bound, alpha)
		s := prob.NewSuite(hypos...)

		obs := []prob.SuiteObservation{
			newLocomotiveObservation(60),
			newLocomotiveObservation(30),
			newLocomotiveObservation(90),
		}
		s.MultiUpdate(obs)

		mean, err := s.Mean()
		if err != nil {
			fmt.Printf("Unable to compute mean due to error [%v]", err)
			continue
		}
		fmt.Printf("Upper bound: %d, Posterior mean: %0.2f ", bound.High, mean)

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
	fmt.Println()
}
