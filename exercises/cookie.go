package exercises

import (
	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// Cookie Problem:
// Bowl 1 contains 30 vanilla cookies and 10 chocolate cookies.
// Bowl 2 contains 20 of each.
// Choose one of the bowls at random and select a cookie at random. The cookie is vanilla.
// What is the probability that it came from Bowl 1?

// CookieByHand computes the probability by manually multiplying the priors by the likelihoods
func CookieByHand() {

	// prior distribution
	p := prob.NewPmf()
	p.Set("Bowl 1", 0.5)
	p.Set("Bowl 2", 0.5)

	// observe a vanilla cookie:
	// likelihood of drawing a vanilla cookie from Bowl 1 is 3/4; likelihood for Bowl 2 is 1/2
	p.Mult("Bowl 1", 0.75)
	p.Mult("Bowl 2", 0.5)

	// Renormalize to obtain the posterior distribution
	p.Normalize()
	p.Print()
}

// prior distribution (uniform) for hypotheses
var bowl1 = prob.NewPmfElement("Bowl 1", 1)
var bowl2 = prob.NewPmfElement("Bowl 2", 1)

type cookieBowl map[string]float64

// define cookie bowls by their distribution of flavors
var cookieMixes = map[string]cookieBowl{
	bowl1.Name: cookieBowl{"chocolate": 0.25, "vanilla": 0.75},
	bowl2.Name: cookieBowl{"chocolate": 0.5, "vanilla": 0.5},
}

// an observation is the name (flavor) of cookie observed
type cookieObservation struct {
	Name string
}

// Getlikelihood is the likelihood function for the Cookie problem
func (o cookieObservation) GetLikelihood(hypoName string) float64 {
	bowlMix, ok := cookieMixes[hypoName]
	if !ok {
		return 0
	}

	like, ok := bowlMix[o.Name]
	if !ok {
		return 0
	}
	return like
}

// Cookie computes the probability (after many other observations) using a suite of hypotheses
func Cookie() {
	c := prob.NewSuite(bowl1, bowl2)
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
