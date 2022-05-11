package kivm

import (
	"kic/boa"
	"testing"
)

func TestKic(t *testing.T) {
	AddCommands()

	if KivmCommand == nil {
		t.Errorf("KicFleet failed, got nil")
		return
	}

	rlen := len(KivmCommand.Commands())
	if rlen != 10 {
		t.Errorf("FleetTest failed, got %d, wanted: 10", rlen)
	}

	boa.ExecCmdNoErrorE(t, KivmCommand)
	// boa.ExecCmdWithErrorE(t, "Use \"kic [command] --help\" for more information about a command.", KivmCommand, "check", "bad-param")
	boa.ExecCmdNoErrorE(t, KivmCommand, "check", "bad-param")
	boa.ExecCmdNoErrorE(t, KivmCommand, "events")
	boa.ExecCmdNoErrorE(t, KivmCommand, "pods")
	boa.ExecCmdNoErrorE(t, KivmCommand, "svc")
}
