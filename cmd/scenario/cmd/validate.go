package cmd

import (
	"fmt"

	"github.com/clearsignalhq-admin/api-scenario-runner/internal/scenario"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate <file>",
	Short: "Validate scenario YAML",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		_, err := scenario.LoadFile(path)
		if err != nil {
			return err
		}
		fmt.Println("OK")
		return nil
	},
}
