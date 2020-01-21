package prob

type observation interface {
	GetLikelihood(string) float64
}

// Suite is a suite of hypotheses with associated probabilities (a Pmf)
type Suite struct {
	*Pmf
}

// NewSuite creates a new Suite
func NewSuite(hypos ...*PmfElement) *Suite {
	s := &Suite{
		Pmf: NewPmf(),
	}
	for _, hypo := range hypos {
		s.Set(hypo.Name, hypo.Prob)
	}
	s.Normalize()
	return s
}

// Update updates the probabilities based on an observation
func (s *Suite) Update(obs observation) {
	for hypoName := range s.prob {
		like := obs.GetLikelihood(hypoName)
		s.Mult(hypoName, like)
	}
	s.Normalize()
}
