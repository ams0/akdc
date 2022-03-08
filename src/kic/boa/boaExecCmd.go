// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package boa

import (
	"bytes"
	"io"
	"kic/boa/cfmt"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func ExecCmdNoErrorE(t *testing.T, c *cobra.Command, args ...string) error {
	t.Helper()

	arg := ""
	if len(args) > 0 {
		arg = args[0]
	}

	cfmt.Info(t.Name(), c.Name(), arg)

	c.SetArgs(args)
	err := c.Execute()

	if err != nil {
		t.Error("Unexpected error:", err)
	}

	return err
}

func ExecCmdWithErrorE(t *testing.T, errMatch string, c *cobra.Command, args ...string) error {
	t.Helper()

	c.SetArgs(args)
	err := c.Execute()

	if err == nil {
		t.Error("Expected error", errMatch)
	} else {
		if !strings.Contains(err.Error(), errMatch) {
			t.Errorf("Expected error %s; got error %s", errMatch, err)
		}
	}

	return err
}

// execute command and return result and error
// warning: this is not thread safe
func ExecCmdWithResultsE(t *testing.T, resultMatch string, errMatch string, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	var (
		old *os.File
		r   *os.File
		w   *os.File
		buf *bytes.Buffer
	)

	// redirect stdout so we can capture
	old = os.Stdout
	r, w, _ = os.Pipe()
	os.Stdout = w

	buf = new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)

	c.SetArgs(args)
	err := c.Execute()

	// reset stdout
	w.Close()
	os.Stdout = old
	io.Copy(buf, r)
	result := strings.TrimSpace(buf.String())

	if !strings.Contains(result, resultMatch) {
		t.Error("Result does not match:", resultMatch)
	}

	if err != nil {
		if !strings.Contains(err.Error(), errMatch) {
			t.Errorf("Expected error %s; got error %s", errMatch, err)
		}
	} else if errMatch != "" {
		t.Errorf("Expected error %s, got nil", errMatch)
	}

	return result, err
}
