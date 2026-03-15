package engine

import (
	"context"
	"time"

	"github.com/clearsignalhq-admin/api-scenario-runner/internal/executor"
	"github.com/clearsignalhq-admin/api-scenario-runner/internal/report"
	"github.com/clearsignalhq-admin/api-scenario-runner/internal/rules"
	"github.com/clearsignalhq-admin/api-scenario-runner/internal/scenario"
)

type Runner struct {
	executor executor.Executor
	rules    rules.Engine
}

func NewRunner() *Runner {
	return &Runner{
		executor: executor.NewHTTPExecutor(),
		rules:    rules.NewEngine(),
	}
}

func (r *Runner) Run(ctx context.Context, s *scenario.Scenario) (*report.ScenarioResult, error) {
	start := time.Now()
	res := &report.ScenarioResult{
		ScenarioName: s.Name,
		StartedAt:    start,
	}

	for _, step := range s.Steps {
		repeat := step.Repeat
		if repeat <= 0 {
			repeat = 1
		}
		for i := 0; i < repeat; i++ {
			sr, err := r.executor.Execute(ctx, step)
			if err != nil {
				res.StepResults = append(res.StepResults, report.StepResult{StepName: step.Name, Ok: false, Error: err.Error()})
				res.Success = false
				res.FinishedAt = time.Now()
				return res, nil
			}
			res.StepResults = append(res.StepResults, report.StepResult{StepName: step.Name, Ok: true, StatusCode: sr.StatusCode, Duration: sr.Duration})
		}
	}

	ruleResults := r.rules.Evaluate(ctx, s, res)
	res.RuleResults = ruleResults

	success := true
	for _, rr := range ruleResults {
		if !rr.Pass {
			success = false
			break
		}
	}
	res.Success = success
	res.FinishedAt = time.Now()
	return res, nil
}
