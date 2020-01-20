package exercises

import (
	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// MontyHall ...
func MontyHall() {
	doorA := prob.NewHypothesis("door A", 1./3)
	doorB := prob.NewHypothesis("door B", 1./3)
	doorC := prob.NewHypothesis("door C", 1./3)

	likelihood := func(obs string, hypoName string) float64 {
		if hypoName == obs {
			return 0
		}
		if hypoName == doorA.Name {
			return 0.5
		}
		return 1
	}

	monty := prob.NewSuite(likelihood, doorA, doorB, doorC)
	monty.Update(doorB.Name)

	monty.Print()
}
