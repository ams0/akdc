package boa

import (
	"strings"
	"testing"
)

func TestReadConfigValue(t *testing.T) {
	s := ReadConfigValue("keyNotFOund")

	if s != "" {
		t.Errorf("ReadConfigValue(keyNotFound) failed, got %v, want: %v", s, "")
	}
}

func TestReadHostIPs(t *testing.T) {
	s := ReadHostIPs("grepNotFound")

	if s != nil {
		t.Errorf("ReadHostIPs(grepNotFound) failed, got %v, want: %v", s, "")
	}
}
func TestGetBoaCommandPath(t *testing.T) {
	s := GetBoaCommandPath()

	if !strings.HasSuffix(s, "/commands/") {
		t.Errorf("GetBoaCommandPath() failed, got %v, want: %v", s, "_kic/commands/")
	}
}

func TestGetRepoBase(t *testing.T) {
	s := GetRepoBase()

	if s == "" {
		t.Errorf("GetRepoBase() failed, got empty string")
	}
}

func TestReadTextFileFromBoaDir(t *testing.T) {
	// s := ReadTextFileFromBoaDir("root.mod")

	// if s == "" {
	// 	t.Errorf("ReadTextFileFromBoaDir(root.mod) failed, got empty string")
	// }
}

func TestCheckForBadChars(t *testing.T) {
	err := CheckForBadChars("hello|world", "test")

	if err == nil {
		t.Errorf("CheckForBadChars() failed, want: error")
	}
}
