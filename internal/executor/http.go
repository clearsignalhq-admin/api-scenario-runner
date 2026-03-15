package executor

import (
	"context"
	"net/http"
	"time"

	"github.com/clearsignalhq-admin/api-scenario-runner/internal/scenario"
)

type HTTPExecutor struct {
	client *http.Client
}

func NewHTTPExecutor() *HTTPExecutor {
	return &HTTPExecutor{
		client: &http.Client{Timeout: 30 * time.Second},
	}
}

func (e *HTTPExecutor) Execute(ctx context.Context, step scenario.Step) (*StepResult, error) {
	start := time.Now()

	// Stub behavior for skeleton:
	// - if request missing, treat as OK with 204
	// - otherwise do not actually call network yet (keeps skeleton safe & deterministic)
	status := 204
	if step.Request != nil {
		status = 200
	}

	return &StepResult{
		StatusCode: status,
		Headers:    map[string][]string{},
		Body:       nil,
		Duration:   time.Since(start),
	}, nil
}
