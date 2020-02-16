package prob

import (
	"math/rand"
	"time"
)

// SuiteObservation is the interface that must be satisfied to update probabilities
type SuiteObservation interface {
	GetLikelihood(float64) float64
}

// Suite is a suite of hypotheses with associated probabilities (a Pmf)
type Suite struct {
	*Pmf
}

// NewSuite creates a new Suite
func NewSuite(hypos ...*PmfElement) *Suite {
	s := &Suite{NewPmf()}
	for _, hypo := range hypos {
		s.Set(hypo)
	}
	s.Normalize()
	return s
}

// Update updates the probabilities based on an observation
func (s *Suite) Update(ob SuiteObservation) {
	for hypoName := range s.prob {
		like := ob.GetLikelihood(hypoName)
		s.Mult(hypoName, like)
	}
	s.Normalize()
}

// UpdateSet updates the probabilities based on multiple observations
func (s *Suite) UpdateSet(obs []SuiteObservation) {
	// iterate elements of obs in random order for numerical stability: avoids long runs
	// of one observation that push the probability of the others to values very close to zero
	rand.Seed(time.Now().UnixNano())
	for _, i := range rand.Perm(len(obs)) {
		ob := obs[i]
		for hypoName := range s.prob {
			like := ob.GetLikelihood(hypoName)
			s.Mult(hypoName, like)
		}
	}
	s.Normalize()
}

// NamedSuiteObservation is the interface that must be satisfied to update probabilities
type NamedSuiteObservation interface {
	GetLikelihood(string) float64
}

// NamedSuite is a suite of hypotheses with associated probabilities (a Pmf)
type NamedSuite struct {
	*NamedPmf
}

// NewNamedSuite creates a new Suite
func NewNamedSuite(hypos ...*NamedPmfElement) *NamedSuite {
	s := &NamedSuite{NewNamedPmf()}
	for _, hypo := range hypos {
		s.Set(hypo)
	}
	s.Normalize()
	return s
}

// Update updates the probabilities based on an observation
func (s *NamedSuite) Update(ob NamedSuiteObservation) {
	for hypoName := range s.nameToIdx {
		like := ob.GetLikelihood(hypoName)
		s.Mult(hypoName, like)
	}
	s.Normalize()
}

// UpdateSet updates the probabilities based on multiple observations
func (s *NamedSuite) UpdateSet(obs []NamedSuiteObservation) {
	// iterate elements of obs in random order for numerical stability: avoids long runs
	// of one observation that push the probability of the others to values very close to zero
	rand.Seed(time.Now().UnixNano())
	for _, i := range rand.Perm(len(obs)) {
		ob := obs[i]
		for hypoName := range s.nameToIdx {
			like := ob.GetLikelihood(hypoName)
			s.Mult(hypoName, like)
		}
	}
	s.Normalize()
}
