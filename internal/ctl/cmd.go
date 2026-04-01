package ctl

import (
	"github.com/spf13/cobra"

	"github.com/libdefinite/definite/internal/ctl/console"
)

// ConsoleCmd returns the cobra command for starting the web console.
func ConsoleCmd() *cobra.Command {
	var port int
	cmd := &cobra.Command{
		Use:   "console",
		Short: "Start web console",
		RunE: func(cmd *cobra.Command, args []string) error {
			return console.New(port).Start()
		},
	}
	cmd.Flags().IntVarP(&port, "port", "p", 8765, "console port")
	return cmd
}

// StatusCmd returns the cobra command for fetching cluster status.
func StatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Get cluster status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
