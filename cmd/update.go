// Copyright 2016 Marapongo, Inc. All rights reserved.

package cmd

import (
	"github.com/spf13/cobra"
)

func newUpdateCmd() *cobra.Command {
	var detailed bool
	var dryRun bool
	var output string
	var cmd = &cobra.Command{
		Use:   "update [snapshot] [blueprint] [-- [args]]",
		Short: "Update an existing environment and its resources",
		Long: "Update an existing environment and its resources.\n" +
			"\n" +
			"This command updates an existing environment whose state is represented by the\n" +
			"existing snapshot file.  The new desired state is computed by compiling and evaluating\n" +
			"a MuPackage blueprint, and extracting all resource allocations from its MuGL graph.\n" +
			"This is then compared against the existing state to determine what operations must take\n" +
			"place to achieve the desired state.  This command results in a full snapshot of the\n" +
			"environment's new resource state, so that it may be updated incrementally again later.\n" +
			"\n" +
			"By default, the MuPackage blueprint is loaded from the current directory.  Optionally,\n" +
			"a path to a MuPackage elsewhere can be provided as the [blueprint] argument.",
		Run: func(cmd *cobra.Command, args []string) {
			applyExisting(cmd, args, applyOptions{
				Delete:   false,
				Detailed: detailed,
				DryRun:   dryRun,
				Output:   output,
			})
		},
	}

	cmd.PersistentFlags().BoolVar(
		&detailed, "detailed", false,
		"Display detailed output during the application of changes")
	cmd.PersistentFlags().BoolVarP(
		&dryRun, "dry-run", "n", false,
		"Don't actually update resources; just print out the planned updates")
	cmd.PersistentFlags().StringVarP(
		&output, "output", "o", "",
		"Serialize the resulting snapshot to a specific file, instead of overwriting the existing one")

	return cmd
}
