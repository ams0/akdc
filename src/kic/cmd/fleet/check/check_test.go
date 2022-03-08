package check

import (
	"kic/boa"
	"testing"
)

func TestCheck(t *testing.T) {
	if CheckCmd == nil {
		t.Errorf("TestCheck failed, got nil")
		return
	}

	rlen := len(CheckCmd.Commands())
	if rlen != 6 {
		t.Errorf("TestCheck len(Commands) failed, got %d, wanted: 6", rlen)
	}

	boa.ExecCmdNoErrorE(t, CheckCmd)
	boa.ExecCmdWithErrorE(t, "exit status 1", CheckCmd, "ai-order-accuracy", "--grep", "bad-grep")
	boa.ExecCmdWithErrorE(t, "exit status 1", CheckCmd, "flux", "--grep", "bad-grep")
	boa.ExecCmdWithErrorE(t, "exit status 1", CheckCmd, "heartbeat", "--grep", "bad-grep")
	boa.ExecCmdWithErrorE(t, "exit status 1", CheckCmd, "logs", "--grep", "bad-grep")
	boa.ExecCmdWithErrorE(t, "exit status 1", CheckCmd, "retries", "--grep", "bad-grep")
	boa.ExecCmdWithErrorE(t, "exit status 1", CheckCmd, "setup", "--grep", "bad-grep")
}
