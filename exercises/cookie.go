package exercises

import (
	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// Cookie Problem:
// Bowl 1 contains 30 vanilla cookies and 10 chocolate cookies.
// Bowl 2 contains 20 of each.
// Choose one of the bowls at random and select a cookie at random. The cookie is vanilla.
// What is the probability that it came from Bowl 1?

type cookieBowl map[string]float64

type cookieHypothesis struct {
	Hypo *prob.PmfElement
	Bowl cookieBowl
}

// prior distribution (uniform) for hypotheses
// define cookie bowls by their distribution of flavors
var (
	bowl1 = cookieHypothesis{
		Hypo: prob.NewPmfElement("Bowl 1", 1),
		Bowl: cookieBowl{
			"chocolate": 0.25,
			"vanilla":   0.75,
		},
	}
	bowl2 = cookieHypothesis{
		Hypo: prob.NewPmfElement("Bowl 2", 1),
		Bowl: cookieBowl{
			"chocolate": 0.5,
			"vanilla":   0.5,
		},
	}
)

var cookieHypos = map[string]cookieHypothesis{
	bowl1.Hypo.Name: bowl1,
	bowl2.Hypo.Name: bowl2,
}

// an observation is the name (flavor) of cookie observed
type cookieObservation struct {
	Name string
}

// Getlikelihood is the likelihood function for the Cookie problem
func (o cookieObservation) GetLikelihood(hypoName string) float64 {
	hypo, ok := cookieHypos[hypoName]
	if !ok {
		return 0
	}

	like, ok := hypo.Bowl[o.Name]
	if !ok {
		return 0
	}
	return like
}

// Cookie computes the probability (after many other observations) using a suite of hypotheses
func Cookie() {
	c := prob.NewSuite(bowl1.Hypo, bowl2.Hypo)
	observations := []cookieObservation{
		cookieObservation{Name: "vanilla"},
		cookieObservation{Name: "chocolate"},
		cookieObservation{Name: "vanilla"},
		cookieObservation{Name: "chocolate"},
		cookieObservation{Name: "chocolate"},
		cookieObservation{Name: "chocolate"},
		cookieObservation{Name: "vanilla"},
		cookieObservation{Name: "chocolate"},
	}

	for _, obs := range observations {
		c.Update(obs)
	}

	c.Print()
}
