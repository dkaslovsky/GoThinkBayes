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

type observation interface {
	GetLikelihood(string) float64
}

// Suite ...
type Suite struct {
	*Pmf
	//like likelihood
}

// NewSuite ...
func NewSuite(hypos ...*Hypothesis) *Suite {
	s := &Suite{
		Pmf: NewPmf(),
		//like: like,
	}
	for _, hypo := range hypos {
		s.Set(hypo.Name, hypo.Prob)
	}
	s.Normalize()
	return s
}

// Update ...
func (s *Suite) Update(obs observation) {
	for hypoName := range s.prob {
		like := obs.GetLikelihood(hypoName)
		s.Mult(hypoName, like)
	}
	s.Normalize()
}

// MultiUpdate ...
// func (s *Suite) MultiUpdate(obs ...observation) {
// 	for _, ob := range obs {
// 		for hypoName := range s.prob {
// 			like := ob.GetLikelihood(hypoName)
// 			s.Mult(hypoName, like)
// 		}
// 	}
// 	s.Normalize()
// }
