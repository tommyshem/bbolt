package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"

	"go.etcd.io/bbolt/version"
)

func newVersionCommand() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "print the current version of bbolt",
		Long:  "print the current version of bbolt",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), "bbolt Version: ", version.Version)
			fmt.Fprintln(cmd.OutOrStdout(), "Go Version: ", runtime.Version())
			fmt.Fprintln(cmd.OutOrStdout(), "Go OS/Arch: ", runtime.GOOS, "/", runtime.GOARCH)
		},
	}

	return versionCmd
}
