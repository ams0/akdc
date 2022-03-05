package check

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	// "github.com/spf13/cobra"
)

func TestCheck(t *testing.T) {
	fmt.Println("todo - use the package")

	if CheckCmd == nil {
		t.Errorf("TestCheck() failed, got nil")
	}

	rlen := len(CheckCmd.Commands())
	if rlen != 6 {
		t.Errorf("TestCheck() failed, got %d, wanted: 6", rlen)
	}
	execute(t, CheckCmd)
	execute(t, CheckCmd, "ai-order-accuracy", "--grep", "bad-grep")
	execute(t, CheckCmd, "flux", "--grep", "bad-grep")
	execute(t, CheckCmd, "heartbeat", "--grep", "bad-grep")
	execute(t, CheckCmd, "logs", "--grep", "bad-grep")
	execute(t, CheckCmd, "retries", "--grep", "bad-grep")
	execute(t, CheckCmd, "setup", "--grep", "bad-grep")
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
