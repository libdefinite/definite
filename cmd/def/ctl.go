package main

import (
	"github.com/libdefinite/definite/internal/ctl"
	"github.com/spf13/cobra"
)

func ctlCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "ctl",
		Short: "Control the server via gRPC",
	}

	c.AddCommand(ctl.FetchCtlCmd(), ctl.WebCtlCmd())
	return c
}
