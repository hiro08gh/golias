package cmd

import (
	"github.com/hiro08gh/golias/external"
	"github.com/spf13/cobra"
)

func addCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new",
		Short: "Create alias",
		Args:  cobra.RangeArgs(2, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}
			external.Add(args)

			return nil
		},
	}

	return cmd
}
