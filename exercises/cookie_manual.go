package exercises

import (
	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// Cookie Problem:
// Bowl 1 contains 30 vanilla cookies and 10 chocolate cookies.
// Bowl 2 contains 20 of each.
// Choose one of the bowls at random and select a cookie at random. The cookie is vanilla.
// What is the probability that it came from Bowl 1?

// CookieManual computes the probability by manually multiplying the priors by the likelihoods
func CookieManual() {

	// prior distribution
	p := prob.NewPmf()
	p.Set(prob.NewPmfElement("Bowl 1", 0.5))
	p.Set(prob.NewPmfElement("Bowl 2", 0.5))

	// observe a vanilla cookie:
	// likelihood of drawing a vanilla cookie from Bowl 1 is 3/4; likelihood for Bowl 2 is 1/2
	p.Mult("Bowl 1", 0.75)
	p.Mult("Bowl 2", 0.5)

	// Renormalize to obtain the posterior distribution
	p.Normalize()
	p.Print()
}
