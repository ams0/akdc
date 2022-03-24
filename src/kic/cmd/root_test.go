package cmd

import (
	"kic/boa"
	"testing"
)

func TestRoot(t *testing.T) {
	if rootCmd == nil {
		t.Errorf("TestRoot() failed, got nil")
	}

	rlen := len(rootCmd.Commands())
	if rlen != 12 {
		t.Errorf("TestRoot() failed, got %d, wanted: 12", rlen)
	}

	boa.ExecCmdNoErrorE(t, rootCmd)
}
