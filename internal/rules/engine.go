package rules

import (
	"context"

	"github.com/clearsignalhq-admin/api-scenario-runner/internal/report"
	"github.com/clearsignalhq-admin/api-scenario-runner/internal/scenario"
)

type Engine interface {
	Evaluate(ctx context.Context, s *scenario.Scenario, res *report.ScenarioResult) []report.RuleResult
}

type StubEngine struct{}

func NewEngine() Engine {
	return &StubEngine{}
}

func (e *StubEngine) Evaluate(ctx context.Context, s *scenario.Scenario, res *report.ScenarioResult) []report.RuleResult {
	out := make([]report.RuleResult, 0, len(s.Rules))
	for _, r := range s.Rules {
		out = append(out, report.RuleResult{Name: r.Name, Pass: true})
	}
	return out
}
