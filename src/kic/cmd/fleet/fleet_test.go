package fleet

import (
	"bytes"
	"fmt"
	"kic/boa"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestFleet(t *testing.T) {
	fmt.Println("todo - use the package")

	if FleetCmd == nil {
		t.Errorf("TestFleet() failed, got nil")
	}

	rlen := len(FleetCmd.Commands())
	if rlen != 10 {
		t.Errorf("TestFleet() failed, got %d, wanted: 10", rlen)
	}

	execute(t, FleetCmd)
	execute(t, FleetCmd, "create", "--ssl", "cseretail.com", "--arc")
	execute(t, FleetCmd, "create", "--ssl")
	execute(t, FleetCmd, "create", "--ssl", "cseretail.com", "--arc", "--do", "-c", "test-cluster")
	boa.ShellExecE("rm -f cluster-test-cluster.sh")
	execute(t, FleetCmd, "delete")
	execute(t, FleetCmd, "exec", "pwd", "--grep", "bad-grep")
	execute(t, FleetCmd, "groups")
	execute(t, FleetCmd, "pull", "--grep", "bad-grep")
	execute(t, FleetCmd, "ssh")
	execute(t, FleetCmd, "sync", "--grep", "bad-grep")
	execute(t, FleetCmd, "token", "--grep", "bad-grep")
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
