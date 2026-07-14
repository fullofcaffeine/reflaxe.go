package main

import (
	"os"
	"testing"
)

func TestScriptedEntrypoint(t *testing.T) {
	previousArgs := os.Args
	t.Cleanup(func() {
		os.Args = previousArgs
	})
	os.Args = []string{previousArgs[0], "--scripted"}
	main()
}
