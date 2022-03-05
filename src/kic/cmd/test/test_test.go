package test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	// "github.com/spf13/cobra"
)

func TestTest(t *testing.T) {
	fmt.Println("todo - use the package")

	if TestCmd == nil {
		t.Errorf("TestTest() failed, got nil")
	}

	rlen := len(TestCmd.Commands())
	if rlen != 2 {
		t.Errorf("TestTest() failed, got %d, wanted: 2", rlen)
	}
	execute(t, TestCmd)
	execute(t, TestCmd, "integration", "-f", "bad-file.json", "-l", "10", "--verbose", "--region", "test", "--zone", "test", "--tag", "test", "--log-format", "tsv", "--dry-run", "--max-errors", "1", "--summary", "Tsv")
	execute(t, TestCmd, "load", "-f", "bad-file.json", "-l", "10", "--verbose", "--region", "test", "--zone", "test", "--tag", "test", "--log-format", "tsv", "--dry-run", "--delay-start", "1", "--duration", "10", "--random")
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
