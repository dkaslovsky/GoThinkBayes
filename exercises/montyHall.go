package exercises

import (
	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// prior distribution
var doorA = prob.NewHypothesis("door A", 1./3)
var doorB = prob.NewHypothesis("door B", 1./3)
var doorC = prob.NewHypothesis("door C", 1./3)

type doorObservation struct {
	Name string
}

// Getlikelihood is the likelihood function for "Monty chooses door B and there is no car there"
// the hypothesis is the door the car is behind
func (o doorObservation) GetLikelihood(hypoName string) float64 {
	if hypoName == o.Name {
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

// MontyHall runs the Monty Hall problem:
// We pick door A, then Monty picks to open door B showing no car
// Should we stick with door A or switch to door C?
func MontyHall() {
	monty := prob.NewSuite(doorA, doorB, doorC)

	obs := doorObservation{Name: "door B"}
	monty.Update(obs)

	monty.Print()
}
