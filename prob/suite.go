package prob

// Hypothesis ...
type Hypothesis struct {
	Name string
	Prob float64
}

// NewHypothesis ...
func NewHypothesis(name string, prob float64) *Hypothesis {
	return &Hypothesis{
		Name: name,
		Prob: prob,
	}
}

type likelihood func(obs string, hypoName string) float64

// Suite ...
type Suite struct {
	*Pmf
	like likelihood
}

// NewSuite ...
func NewSuite(hypos []*Hypothesis, like likelihood) *Suite {
	s := &Suite{
		Pmf:  NewPmf(),
		like: like,
	}
	for _, hypo := range hypos {
		s.Set(hypo.Name, hypo.Prob)
	}
	s.Normalize()
	return s
}

// Update ...
func (s *Suite) Update(obs string) {
	for hypoName := range s.prob {
		like := s.like(obs, hypoName)
		s.Mult(hypoName, like)
	}
	s.Normalize()
}

// MultiUpdate ...
func (s *Suite) MultiUpdate(obs []string) {
	for _, ob := range obs {
		for hypoName := range s.prob {
			like := s.like(ob, hypoName)
			s.Mult(hypoName, like)
		}
	}
	s.Normalize()
}
