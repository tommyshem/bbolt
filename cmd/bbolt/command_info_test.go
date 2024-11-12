package main_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	main "go.etcd.io/bbolt/cmd/bbolt"
	"go.etcd.io/bbolt/internal/btesting"
)

// Ensure the "info" command can print information about a database.
//
// This test case verifies that the "info" command can print information about a
// database. It checks that the output of the command matches the expected result.
func TestInfoCommand_Run(t *testing.T) {
	// Create a test database.
	db := btesting.MustCreateDB(t)
	db.Close()

	defer requireDBNoChange(t, dbData(t, db.Path()), db.Path())

	// Run the info command.
	t.Log("Running info command")
	rootCmd := main.NewRootCommand()
	_, actualOutput, err := executeCommand(rootCmd, "info", db.Path())
	require.NoError(t, err)

	// check results
	t.Log("Verify result")
	expected := "Page Size: 4096\n"
	require.EqualValues(t, expected, actualOutput)
}
