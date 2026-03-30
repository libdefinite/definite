package main

import (
	"github.com/spf13/cobra"

	"github.com/libdefinite/definite/internal/ctl"
)

func ctlCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "ctl",
		Short: "Control the server via gRPC",
	}

	c.AddCommand(ctl.StatusCmd(), ctl.ConsoleCmd())
	return c
}
