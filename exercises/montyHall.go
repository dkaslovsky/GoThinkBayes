package exercises

import (
	"fmt"

	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// MontyHall ...
func MontyHall() {
	doorA := prob.NewHypothesis("door A", 1./3)
	doorB := prob.NewHypothesis("door B", 1./3)
	doorC := prob.NewHypothesis("door C", 1./3)
	hypos := []*prob.Hypothesis{doorA, doorB, doorC}

	likelihood := func(obs string, hypoName string) float64 {
		if hypoName == obs {
			return 0
		}
		if hypoName == doorA.Name {
			return 0.5
		}
		return 1
	}

	monty := prob.NewSuite(hypos, likelihood)
	monty.Update(doorB.Name)

	fmt.Println("posterior")
	monty.Print()
}
