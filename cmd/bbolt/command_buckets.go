package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// newBucketsCommand creates a new command that prints a list of buckets in the given Bolt database.
//
// The path to a Bolt database must be specified as an argument.
func newBucketsCommand() *cobra.Command {
	var bucketsCmd = &cobra.Command{
		Use:   "buckets <path>",
		Short: "Print a list of buckets",
		Long:  "Print a list of buckets in the given Bolt database\nThe path to a Bolt database must be specified as an argument",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return printBucketsList(cmd, args[0])
		},
	}
	return bucketsCmd
}

// printBucketsList prints a list of buckets in the given Bolt database.
func printBucketsList(cmd *cobra.Command, path string) error {
	// Required database path.
	if path == "" {
		return ErrPathRequired
		// Verify if the specified database file exists.
	} else if _, err := os.Stat(path); os.IsNotExist(err) {
		return ErrFileNotFound
	}

	// Open database.
	db, err := bolt.Open(path, 0600, &bolt.Options{ReadOnly: true})
	if err != nil {
		return err
	}
	defer db.Close()

	// Print the list of buckets in the database.
	return db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			fmt.Fprintln(cmd.OutOrStdout(), string(name))
			return nil
		})
	})
}
