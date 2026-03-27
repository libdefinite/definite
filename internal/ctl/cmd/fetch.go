package cmd

import (
	"github.com/spf13/cobra"
)

// FetchCtlCmd returns the cobra command for fetching server status.
func FetchCtlCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Get server status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
