package boa

import (
	"path/filepath"
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
	s, err := ReadHostIPs("grepNotFound")

	if err == nil {
		t.Errorf("Expected error")
	}

	if s != nil {
		t.Errorf("ReadHostIPs(grepNotFound) failed, got %v, want: %v", s, "")
	}
}
func TestGetBoaCommandPath(t *testing.T) {
	s := GetBoaCommandPath()

	if !strings.HasSuffix(s, "/commands") {
		t.Errorf("GetBoaCommandPath() failed, got %v, want: %v", s, "/commands")
	}
}

func TestGetRepoBase(t *testing.T) {
	s := GetRepoBase()

	if s == "" {
		t.Errorf("GetRepoBase() failed, got empty string")
	}
}

func TestReadTextFileFromBoaDir(t *testing.T) {
	// create a temp boa file
	path := GetBoaPath()
	ShellExecE("mkdir -p " + path)
	file := filepath.Join(path, "testio.boa")
	ShellExecE("echo 'command' > " + file)
	ShellExecE("echo '' >> " + file)
	ShellExecE("echo 'name: test2' >> " + file)
	ShellExecE("echo '' >> " + file)
	ShellExecE("echo 'short: testing2' >> " + file)
	ShellExecE("echo '' >> " + file)

	s := ReadTextFileFromBoaDir("testio.boa")

	if s == "" {
		t.Errorf("ReadTextFileFromBoaDir(testio.boa) failed, got empty string")
	}

	lines := ReadLinesFromFile(file)

	if len(lines) < 1 {
		t.Errorf("ReadLinesFromFile(testio.boa) failed, got empty array")
	}

	ShellExecE("rm -f " + file)
}

func TestCheckForBadChars(t *testing.T) {
	err := CheckForBadChars("hello|world", "test")

	if err == nil {
		t.Errorf("CheckForBadChars() failed, want: error")
	}
}
