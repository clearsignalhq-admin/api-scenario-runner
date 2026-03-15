package report

import "time"

type ScenarioResult struct {
	ScenarioName string
	StartedAt    time.Time
	FinishedAt   time.Time
	Success      bool

	StepResults []StepResult
	RuleResults []RuleResult
}

type StepResult struct {
	StepName    string
	Ok          bool
	StatusCode  int
	Duration    time.Duration
	Error       string
}

type RuleResult struct {
	Name string
	Pass bool
	Info string
}
