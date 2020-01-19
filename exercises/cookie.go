package exercises

import (
	"fmt"

	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// Bowl 1 contains 30 vanilla cookies and 10 chocolate cookies.
// Bowl 2 contains 20 of each.
// Choose one of the bowls at random and select a cookie at random. The cookie is vanilla.
// What is the probability that it came from Bowl 1?

type bowl map[string]float64

type cookieProblem struct {
	*prob.Pmf
	Mixes map[string]bowl
}

func newCookieProblem(mixes map[string]bowl) *cookieProblem {
	c := &cookieProblem{
		Pmf:   prob.NewPmf(),
		Mixes: mixes,
	}
	for hypo := range mixes {
		c.Set(hypo, 1)
	}
	c.Normalize()
	return c
}

func (c *cookieProblem) update(obs string) {
	for hypo := range c.Mixes {
		like := c.likelihood(obs, hypo)
		c.Mult(hypo, like)
	}
	c.Normalize()
}

func (c *cookieProblem) likelihood(obs string, hypo string) float64 {
	mix := c.Mixes[hypo]
	like := mix[obs]
	return like
}

// CookieManual ...
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
	fmt.Println("posterior")
	p.Print()
}

// CookieEncapsulated ...
func CookieEncapsulated() {

	mixes := map[string]bowl{
		"Bowl 1": bowl{"chocolate": 0.25, "vanilla": 0.75},
		"Bowl 2": bowl{"chocolate": 0.5, "vanilla": 0.5},
	}

	c := newCookieProblem(mixes)

	observations := []string{
		"vanilla",
		"chocolate",
		"vanilla",
		"chocolate",
		"chocolate",
		"chocolate",
		"vanilla",
		"chocolate",
	}
	for _, obs := range observations {
		c.update(obs)
	}

	fmt.Println("posterior")
	c.Print()
}

// CookieSuite ...
func CookieSuite() {

	bowl1 := prob.NewHypothesis("Bowl 1", 0.5)
	bowl2 := prob.NewHypothesis("Bowl 2", 0.5)
	hypos := []*prob.Hypothesis{bowl1, bowl2}

	mixes := map[string]bowl{
		"Bowl 1": bowl{"chocolate": 0.25, "vanilla": 0.75},
		"Bowl 2": bowl{"chocolate": 0.5, "vanilla": 0.5},
	}

	likelihood := func(obs string, hypoName string) float64 {
		hypoMix, ok := mixes[hypoName]
		if !ok {
			return 0
		}

		like, ok := hypoMix[obs]
		if !ok {
			return 0
		}
		return like

	}

	c := prob.NewSuite(hypos, likelihood)
	observations := []string{
		"vanilla",
		"chocolate",
		"vanilla",
		"chocolate",
		"chocolate",
		"chocolate",
		"vanilla",
		"chocolate",
	}
	for _, obs := range observations {
		c.Update(obs)
	}

	fmt.Println("posterior")
	c.Print()
}
