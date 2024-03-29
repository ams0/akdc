package test

import (
	"fmt"
	"kic/boa"
	"testing"
)

func TestTest(t *testing.T) {
	if TestCmd == nil {
		t.Errorf("TestTest failed, got nil")
		return
	}

	rlen := len(TestCmd.Commands())
	if rlen != 2 {
		t.Errorf("TestTest failed, got %d, wanted: 2", rlen)
	}

	boa.ExecCmdNoErrorE(t, TestCmd)
	// todo - fix bug
	// boa.ExecCmdNoErrorE(t, TestCmd, "integration", "--dry-run", "-l", "10", "--verbose", "--region", "test", "--zone", "test", "--tag", "test", "--log-format", "tsv", "--max-errors", "1", "--summary", "Tsv")
	boa.ExecCmdNoErrorE(t, TestCmd, "load", "--dry-run", "-f", "bad-file.json", "-l", "10", "--verbose", "--region", "test", "--zone", "test", "--tag", "test", "--log-format", "tsv", "--delay-start", "1", "--duration", "10", "--random")

	fmt.Println(boa.GetBinDir(), boa.GetBinName())
}
