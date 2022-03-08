package targets

import (
	"kic/boa"
	"testing"
)

func TestTargets(t *testing.T) {
	if TargetsCmd == nil {
		t.Errorf("TestTargets() failed, got nil")
	}

	rlen := len(TargetsCmd.Commands())
	if rlen != 5 {
		t.Errorf("TestTargets() failed, got %d, wanted: 5", rlen)
	}

	boa.ExecCmdWithErrorE(t, "", TargetsCmd, "list")
	boa.ExecCmdWithErrorE(t, "", TargetsCmd, "clear")

	// create a test file
	boa.ShellExecE("mkdir -p autogitops")
	boa.ShellExecE(`echo '{ "targets": [ "test" ] }' > autogitops/autogitops.json`)

	boa.ExecCmdNoErrorE(t, TargetsCmd)
	boa.ExecCmdNoErrorE(t, TargetsCmd, "list")
	boa.ExecCmdNoErrorE(t, TargetsCmd, "add", "foo")
	// todo - fix this test
	// boa.ExecCmdNoErrorE(t, TargetsCmd, "remove", "foo")
	boa.ExecCmdNoErrorE(t, TargetsCmd, "clear")
	// do not run push!

	// remove test file
	boa.ShellExecE("rm -rf autogitops")
}
