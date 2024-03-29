// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

// special version of kic for the kubecon hands-on labs

package kubekic

import (
	"kic/boa"
	"testing"
)

func TestKic(t *testing.T) {
	if KicCmd == nil {
		t.Errorf("KicFleet failed, got nil")
		return
	}

	rlen := len(KicCmd.Commands())
	if rlen != 7 {
		t.Errorf("FleetTest failed, got %d, wanted: 7", rlen)
	}

	boa.ExecCmdNoErrorE(t, KicCmd)
	boa.ExecCmdNoErrorE(t, KicCmd, "check")
	boa.ExecCmdNoErrorE(t, KicCmd, "cluster")
	boa.ExecCmdNoErrorE(t, KicCmd, "check", "bad-param")
	boa.ExecCmdNoErrorE(t, KicCmd, "events")
	boa.ExecCmdNoErrorE(t, KicCmd, "pods")
	boa.ExecCmdNoErrorE(t, KicCmd, "svc")
	boa.ExecCmdNoErrorE(t, KicCmd, "test")
}
