package executor

import (
	"context"
	"time"

	"github.com/clearsignalhq-admin/api-scenario-runner/internal/scenario"
)

type StepResult struct {
	StatusCode int
	Headers    map[string][]string
	Body       []byte
	Duration   time.Duration
}

type Executor interface {
	Execute(ctx context.Context, step scenario.Step) (*StepResult, error)
}
