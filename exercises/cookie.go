package exercises

import (
	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// The Cookie Problem:
// Bowl 1 contains 30 vanilla cookies and 10 chocolate cookies.
// Bowl 2 contains 20 of each.
// Choose one of the bowls at random and select a cookie at random. The cookie is vanilla.
// What is the probability that it came from Bowl 1?

type bowl map[string]float64

// CookieManual computes the probability by manually multiplying the priors by the likelihoods
func CookieManual() {

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
var bowl1 = prob.NewHypothesis("Bowl 1", 0.5)
var bowl2 = prob.NewHypothesis("Bowl 2", 0.5)

type cookieObservation struct {
	Name string
}

func (o cookieObservation) GetLikelihood(hypoName string) float64 {
	mixes := map[string]bowl{
		bowl1.Name: bowl{"chocolate": 0.25, "vanilla": 0.75},
		bowl2.Name: bowl{"chocolate": 0.5, "vanilla": 0.5},
	}

	hypoMix, ok := mixes[hypoName]
	if !ok {
		return 0
	}

	like, ok := hypoMix[o.Name]
	if !ok {
		return 0
	}
	return like
}

// CookieSuite computes the probability (after many other observations) using a suite of hypotheses
func CookieSuite() {
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
	//c.MultiUpdate(observations)
	for _, obs := range observations {
		c.Update(obs)
	}
	c.Print()
}
