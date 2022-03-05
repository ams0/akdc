package boa

import (
	"fmt"
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
	path := GetBoaPath()
	ShellExecE("mkdir -p " + path)
	file := path + "test.boa"
	ShellExecE("echo 'command' > " + file)
	ShellExecE("echo '' >> " + file)
	ShellExecE("echo 'name: test2' >> " + file)
	ShellExecE("echo '' >> " + file)
	ShellExecE("echo 'short: testing2' >> " + file)
	ShellExecE("echo '' >> " + file)

	fmt.Println(len(root.Commands()))

	LoadCommands(root)

	fmt.Println(len(root.Commands()))

	if root == nil {
		t.Errorf("TestLoadCommands() failed, got nil")
	}

	nr := SetNewRoot()

	if nr == nil {
		t.Errorf("TestSetNewRoot() failed, got nil")
	}

	ShellExecE("rm -rf " + path)
}
