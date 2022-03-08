package boa

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestBoa(t *testing.T) {
	var root = &cobra.Command{
		Use:   "kic",
		Short: "Testing",
	}

	var foo = &cobra.Command{
		Use:   "foo",
		Short: "Testing",
	}

	root.AddCommand(foo)

	c := GetCommandByUse(root, "foo")

	if c == nil {
		t.Errorf("TestGetCommandByUse() failed, got nil")
	}

	// create a temp boa file
	LoadCommands(root)

	if root == nil {
		t.Errorf("TestLoadCommands() failed, got nil")
	}

	nr := SetNewRoot()

	if nr == nil {
		t.Errorf("TestSetNewRoot() failed, got nil")
	}
}
