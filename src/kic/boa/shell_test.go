package boa

import (
	"testing"
)

func TestShellExecArgsE(t *testing.T) {
	err := ShellExecArgsE("pwd", nil)

	if err != nil {
		t.Errorf("TestShellExecArgsE() failed, error %v", err)
	}
}

func TestShellExecE(t *testing.T) {
	err := ShellExecE("pwd")

	if err != nil {
		t.Errorf("TestShellExecE() failed, error %v", err)
	}
}

func TestShellExecOut(t *testing.T) {
	_, err := ShellExecOut("pwd", false)

	if err != nil {
		t.Errorf("TestShellExecOut() failed, error %v", err)
	}
}
