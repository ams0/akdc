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
	if rlen != 6 {
		t.Errorf("TestRoot() failed, got %d, wanted: 6", rlen)
	}

	boa.ExecCmdNoErrorE(t, rootCmd)
}
