package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// bucketsCmd represents the buckets command
func newBucketsCommand() *cobra.Command {
	var bucketsCmd = &cobra.Command{
		Use:   "buckets <path>",
		Short: "Print a list of buckets",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return bucketsfunc(args[0])
		},
	}
	return bucketsCmd
}

// Run executes the command.
func bucketsfunc(path string) error {
	// Require database path.
	if path == "" {
		return ErrPathRequired
	} else if _, err := os.Stat(path); os.IsNotExist(err) {
		return ErrFileNotFound
	}

	// Open database.
	db, err := bolt.Open(path, 0600, &bolt.Options{ReadOnly: true})
	if err != nil {
		return err
	}
	defer db.Close()

	// Print buckets.
	return db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			fmt.Fprintln(os.Stdout, string(name))
			return nil
		})
	})
}
