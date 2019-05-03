package dominator

import (
	"github.com/tomocy/psycho-pass/sibyl"
	"github.com/tomocy/psycho-pass/target"
)

type Dominator struct{}

func (d *Dominator) Enforce(target target.Target) {
	enforcer := sibyl.Sibyl.GiveOracle(target)
	enforcer.Prompt()
	enforcer.Enforce(target)
}
