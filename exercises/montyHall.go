package exercises

import (
	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// Monty Hall problem:
// We pick door A, then Monty picks to open door B showing no car
// Should we stick with door A or switch to door C?

// prior distribution for location of car (uniform)
var (
	doorA = prob.NewPmfElement("door A", 1)
	doorB = prob.NewPmfElement("door B", 1)
	doorC = prob.NewPmfElement("door C", 1)
)

// an observation is the name of a revealed door
type doorObservation struct {
	name string
}

// Getlikelihood is the likelihood function for "Monty chooses door B and there is no car there"
func (o doorObservation) GetLikelihood(hypoName string) float64 {
	if hypoName == o.name {
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

// MontyHall runs the Monty Hall problem
func MontyHall() {
	monty := prob.NewSuite(doorA, doorB, doorC)

	obs := doorObservation{name: "door B"}
	monty.Update(obs)

	monty.Print()
}
