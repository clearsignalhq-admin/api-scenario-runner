package cmd

import (
	"github.com/clearsignalhq-admin/api-scenario-runner/internal/engine"
	"github.com/clearsignalhq-admin/api-scenario-runner/internal/report"
	"github.com/clearsignalhq-admin/api-scenario-runner/internal/scenario"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <file>",
	Short: "Run a scenario",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		s, err := scenario.LoadFile(path)
		if err != nil {
			return err
		}

		r := engine.NewRunner()
		res, err := r.Run(cmd.Context(), s)
		if err != nil {
			return err
		}

		reporter := report.NewConsoleReporter()
		reporter.Print(res)
		return nil
	},
}
