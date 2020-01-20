package exercises

import (
	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// MontyHall runs the Monty Hall problem:
// We pick door A, then Monty picks to open door B showing no car
// Should we stick with door A or switch to door C?
func MontyHall() {
	doorA := prob.NewHypothesis("door A", 1./3)
	doorB := prob.NewHypothesis("door B", 1./3)
	doorC := prob.NewHypothesis("door C", 1./3)

	// likelihood function representing "Monty chooses door B and there is no car there"
	// the observation is the door that Monty opens
	// the hypothesis is the door the car is behind
	likelihood := func(obs string, hypoName string) float64 {
		if hypoName == obs {
			// we only observe a door that Monty shows which cannot contain the car
			return 0
		}
		if hypoName == doorA.Name {
			// under the hypothesis that the car is behind A, Monty can choose B or C
			// and the probability that the car is not behind B is 1
			return 0.5
		}
		// if the car is behind door C, Monty must open door B and the car cannot not be there
		return 1
	}

	monty := prob.NewSuite(likelihood, doorA, doorB, doorC)
	monty.Update(doorB.Name)

	monty.Print()
}
