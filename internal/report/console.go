package report

import (
	"fmt"
)

type ConsoleReporter struct{}

func NewConsoleReporter() *ConsoleReporter {
	return &ConsoleReporter{}
}

func (r *ConsoleReporter) Print(res *ScenarioResult) {
	fmt.Printf("Running scenario: %s\n\n", res.ScenarioName)

	for _, sr := range res.StepResults {
		if sr.Ok {
			fmt.Printf("Step %s .......... OK (%d)\n", sr.StepName, sr.StatusCode)
			continue
		}
		fmt.Printf("Step %s .......... FAIL (%s)\n", sr.StepName, sr.Error)
	}

	if len(res.RuleResults) > 0 {
		fmt.Println()
	}
	for _, rr := range res.RuleResults {
		status := "PASS"
		if !rr.Pass {
			status = "FAIL"
		}
		fmt.Printf("Rule %s .. %s\n", rr.Name, status)
	}

	fmt.Println()
	if res.Success {
		fmt.Println("Scenario result: SUCCESS")
	} else {
		fmt.Println("Scenario result: FAILURE")
	}
}
