package boa

import (
	"path/filepath"
	"testing"
)

func TestReadConfigValue(t *testing.T) {
	s := ReadConfigValue("keyNotFOund")

	if s != "" {
		t.Errorf("ReadConfigValue(keyNotFound) failed, got %v, want: %v", s, "")
	}
}

func TestIo(t *testing.T) {
	if s := GetBinDir(); s == "" {
		t.Errorf("GetBinDir failed")
	}
	if s := GetBinName(); s == "" {
		t.Errorf("GetBinName failed")
	}
	if s := GetBoaPath(); s == "" {
		t.Errorf("GetBoaPath failed")
	}
	if s := GetBoaCommandPath(); s == "" {
		t.Errorf("GetBoaCommandPath failed")
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

func TestReadTextFileFromBoaDir(t *testing.T) {
	// create a temp boa file
	name := "testio.boa"
	path := GetBoaPath()
	ShellExecE("mkdir -p " + path)
	file := filepath.Join(path, name)
	ShellExecE("echo 'command' > " + file)
	ShellExecE("echo '' >> " + file)
	ShellExecE("echo 'name: test2' >> " + file)
	ShellExecE("echo '' >> " + file)
	ShellExecE("echo 'short: testing2' >> " + file)
	ShellExecE("echo '' >> " + file)

	s := ReadTextFileFromBoaDir(name)

	if s == "" {
		t.Errorf("ReadTextFileFromBoaDir(testio.boa) failed, got empty string")
	}

	lines := ReadLinesFromBoaFile(name)

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
