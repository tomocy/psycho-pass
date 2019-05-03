package sibyl

import (
	"math/rand"

	"github.com/tomocy/psycho-pass/enforcer"
	"github.com/tomocy/psycho-pass/target"
)

var Sibyl = new(sibyl)

type sibyl struct{}

func (s *sibyl) GiveOracle(target target.Target) enforcer.Enforcer {
	cc, err := s.calculateCrimeCoefficient(target)
	if err != nil {
		return new(enforcer.Decomposer)
	}

	if 300 <= cc {
		return new(enforcer.Eliminator)
	}

	return new(enforcer.Eliminator)
}

func (s *sibyl) calculateCrimeCoefficient(target target.Target) (uint, error) {
	cc := uint(rand.Intn(600))
	return cc, nil
}
