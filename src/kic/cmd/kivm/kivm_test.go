package kivm

import (
	"kic/boa"
	"testing"
)

func TestKic(t *testing.T) {
	KivmCmd = LoadCommands(KivmCmd)

	if KivmCmd == nil {
		t.Errorf("KicFleet failed, got nil")
		return
	}

	rlen := len(KivmCmd.Commands())
	if rlen != 10 {
		t.Errorf("FleetTest failed, got %d, wanted: 10", rlen)
	}

	boa.ExecCmdNoErrorE(t, KivmCmd)
	// boa.ExecCmdWithErrorE(t, "Use \"kic [command] --help\" for more information about a command.", KivmCmd, "check", "bad-param")
	boa.ExecCmdNoErrorE(t, KivmCmd, "check", "bad-param")
	boa.ExecCmdNoErrorE(t, KivmCmd, "events")
	boa.ExecCmdNoErrorE(t, KivmCmd, "pods")
	boa.ExecCmdNoErrorE(t, KivmCmd, "svc")
}
