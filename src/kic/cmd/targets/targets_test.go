package targets

import (
	"bytes"
	"kic/boa"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	// "github.com/spf13/cobra"
)

func TestTargets(t *testing.T) {
	if TargetsCmd == nil {
		t.Errorf("TestTargets() failed, got nil")
	}

	rlen := len(TargetsCmd.Commands())
	if rlen != 5 {
		t.Errorf("TestTargets() failed, got %d, wanted: 5", rlen)
	}

	// create a test file
	boa.ShellExecE("mkdir -p autogitops")
	boa.ShellExecE(`echo '{ "targets": [ "test" ] }' > autogitops/autogitops.json`)

	execute(t, TargetsCmd)
	execute(t, TargetsCmd, "list")
	execute(t, TargetsCmd, "add", "foo")
	execute(t, TargetsCmd, "remove", "foo")
	execute(t, TargetsCmd, "clear")

	boa.ShellExecE("rm -rf autogitops")
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
