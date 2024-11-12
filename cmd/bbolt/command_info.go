package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
) // infoCommand represents the "info" command execution.
type infoCommand struct {
	baseCommand
}

// newInfoCommand returns a infoCommand.
func newInfoCommand() *cobra.Command {
	var infoCmd = &cobra.Command{
		Use:   "info <path>",
		Short: "Info prints basic information about the Bolt database at PATH.",
		Long:  "Info prints basic information about the Bolt database at PATH.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return info(cmd, args[0])
		},
	}
	return infoCmd
}

// Run executes the command.
func info(cmd *cobra.Command, path string) error {

	// Require database path.
	if path == "" {
		return ErrPathRequired
	} else if _, err := os.Stat(path); os.IsNotExist(err) {
		return ErrFileNotFound
	}

	// Open the database.
	db, err := bolt.Open(path, 0600, &bolt.Options{ReadOnly: true})
	if err != nil {
		return err
	}
	defer db.Close()

	// Print basic database info.
	info := db.Info()
	fmt.Fprintf(cmd.OutOrStdout(), "Page Size: %d\n", info.PageSize)

	return nil
}
