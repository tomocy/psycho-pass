package enforcer

import (
	"fmt"

	"github.com/tomocy/psycho-pass/target"
)

type Enforcer interface {
	Prompt()
	Enforce(target target.Target)
}

type Paralyzer struct{}

func (p *Paralyzer) Prompt() {
	fmt.Println("Enforcement mode: Non-Lethal Paralyzer. Please aim calmly and subdue the target.")
}

func (p *Paralyzer) Enforce(target target.Target) {}

type Eliminator struct{}

func (e *Eliminator) Prompt() {
	fmt.Println("Enforcement mode: Lethal Eliminator. Please aim carefully and eliminate the target")
}

func (e *Eliminator) Enforce(target target.Target) {}

type Decomposer struct{}

func (d *Decomposer) Prompt() {
	fmt.Println("Enforcement mode: Destroy Decomposer. Target will be completely annihilated. Please proceed with maximum caution.")
}

func (d *Decomposer) Enforce(target target.Target) {}
