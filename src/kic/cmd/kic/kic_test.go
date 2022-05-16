package kic

import (
	"kic/boa"
	"testing"
)

func TestKic(t *testing.T) {
	KicCmd = LoadCommands(KicCmd)

	if KicCmd == nil {
		t.Errorf("KicFleet failed, got nil")
		return
	}

	rlen := len(KicCmd.Commands())
	if rlen != 9 {
		t.Errorf("FleetTest failed, got %d, wanted: 9", rlen)
	}

	boa.ExecCmdNoErrorE(t, KicCmd)
	// boa.ExecCmdWithErrorE(t, "Use \"kic [command] --help\" for more information about a command.", KicCmd, "check", "bad-param")
	boa.ExecCmdNoErrorE(t, KicCmd, "check")
	boa.ExecCmdNoErrorE(t, KicCmd, "check", "bad-param")
	boa.ExecCmdNoErrorE(t, KicCmd, "cluster")
	boa.ExecCmdNoErrorE(t, KicCmd, "env")
	boa.ExecCmdNoErrorE(t, KicCmd, "events")
	boa.ExecCmdNoErrorE(t, KicCmd, "pods")
	boa.ExecCmdNoErrorE(t, KicCmd, "svc")
	boa.ExecCmdNoErrorE(t, KicCmd, "targets")
	boa.ExecCmdNoErrorE(t, KicCmd, "test")
}
