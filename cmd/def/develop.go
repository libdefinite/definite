package main

import "github.com/spf13/cobra"

func webConsoleCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "web-console",
		Short: "start management web console",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
