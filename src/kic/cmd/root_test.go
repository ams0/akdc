package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	// "github.com/spf13/cobra"
)

func TestRoot(t *testing.T) {
	fmt.Println("todo - use the package")

	if rootCmd == nil {
		t.Errorf("TestRoot() failed, got nil")
	}

	rlen := len(rootCmd.Commands())
	if rlen != 3 {
		t.Errorf("TestRoot() failed, got %d, wanted: 3", rlen)
	}

	out, err := execute(t, rootCmd)

	fmt.Println(err == nil, len(out))

}

func execute(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()

	return strings.TrimSpace(buf.String()), err
}
