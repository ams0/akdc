package fleet

import (
	"fmt"
	"kic/boa"
	"kic/boa/cfmt"
	"os"
	"strings"
	"testing"
	"time"
)

// set the cluster name
var testCluster = "test-tx-atx-101"

func TestIntegration(t *testing.T) {
	// skip the tests if env var not set
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		fmt.Println("Skipping integration tests")
		return
	}

	// we can't get these from go test
	if os.Getenv("KIC_PATH") == "" || os.Getenv("KIC_NAME") == "" {
		t.Error("please export KIC_PATH and KIC_NAME")
		return
	}

	fmt.Println("Integration test")

	if FleetCmd == nil {
		t.Errorf("TestFleet failed, got nil")
		return
	}

	var (
		s   string
		err error
	)

	cfmt.Info("Creating cluster")
	err = boa.ExecCmdNoErrorE(t, FleetCmd, "create", "--gitops", "--cores", "2", "--ssl", "cseretail.com", "--cluster", testCluster)

	if err != nil {
		cfmt.ErrorE(err)
		return
	}

	cfmt.Info("waiting for setup to complete")
	time.Sleep(60 * time.Second)

	i := 0

	// retry for up to 10 min
	// go test timeout default is 10 min
	for i < 60 {
		s, err = boa.ExecCmdWithResultsE(t, "", "", FleetCmd, "check", "setup")

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("setup:", s)

		if strings.Contains(s, "complete") {
			break
		}

		i++
		time.Sleep(10 * time.Second)
	}

	// VM create failed
	if !strings.Contains(s, "complete") {
		t.Error("vm create failed")
		return
	}

	cfmt.Info("Running checks")
	boa.ExecCmdNoErrorE(t, FleetCmd, "check", "ai-order-accuracy")
	boa.ExecCmdNoErrorE(t, FleetCmd, "check", "flux")
	boa.ExecCmdNoErrorE(t, FleetCmd, "check", "heartbeat")
	boa.ExecCmdNoErrorE(t, FleetCmd, "check", "logs")
	boa.ExecCmdNoErrorE(t, FleetCmd, "check", "retries")

	boa.ExecCmdNoErrorE(t, FleetCmd, "check", "ai-order-accuracy", "--grep", testCluster)
	boa.ExecCmdNoErrorE(t, FleetCmd, "check", "flux", "--grep", testCluster)
	boa.ExecCmdNoErrorE(t, FleetCmd, "check", "heartbeat", "--grep", testCluster)
	boa.ExecCmdNoErrorE(t, FleetCmd, "check", "logs", "--grep", testCluster)
	boa.ExecCmdNoErrorE(t, FleetCmd, "check", "retries", "--grep", testCluster)

	boa.ExecCmdNoErrorE(t, FleetCmd, "exec", "pwd")
	boa.ExecCmdNoErrorE(t, FleetCmd, "list")
	boa.ExecCmdNoErrorE(t, FleetCmd, "patch")
	boa.ExecCmdNoErrorE(t, FleetCmd, "pull")
	boa.ExecCmdNoErrorE(t, FleetCmd, "sync")

	boa.ExecCmdNoErrorE(t, FleetCmd, "exec", "pwd", "--grep", testCluster)
	boa.ExecCmdNoErrorE(t, FleetCmd, "list", "--grep", testCluster)
	boa.ExecCmdNoErrorE(t, FleetCmd, "patch", "--grep", testCluster)
	boa.ExecCmdNoErrorE(t, FleetCmd, "pull", "--grep", testCluster)
	boa.ExecCmdNoErrorE(t, FleetCmd, "sync", "--grep", testCluster)

	boa.ExecCmdNoErrorE(t, FleetCmd, "groups")

	cfmt.Info("Deleting cluster")
	boa.ExecCmdNoErrorE(t, FleetCmd, "delete", testCluster)
	os.Remove("ips")
}

func TestFleet(t *testing.T) {
	if FleetCmd == nil {
		t.Errorf("TestFleet failed, got nil")
		return
	}

	rlen := len(FleetCmd.Commands())
	if rlen != 5 {
		t.Errorf("FleetTest failed, got %d, wanted: 5", rlen)
	}

	boa.ExecCmdNoErrorE(t, FleetCmd)
	boa.ExecCmdNoErrorE(t, FleetCmd, "delete", "test___command", "-h")

	boa.ExecCmdNoErrorE(t, FleetCmd, "create", "--ssl", "testing.com", "--arc", "--do", "-c", "test-cluster")
	boa.ExecCmdNoErrorE(t, FleetCmd, "create", "--dry-run", "--ssl", "testing.com", "--arc", "-c", "test-cluster")
	boa.ExecCmdNoErrorE(t, FleetCmd, "list")

	boa.ExecCmdWithErrorE(t, "exit status 1", FleetCmd, "exec", "pwd", "--grep", "bad-grep")
	boa.ExecCmdWithErrorE(t, "accepts 1 arg", FleetCmd, "ssh")
	boa.ExecCmdWithErrorE(t, "flag needs an argument: --ssl", FleetCmd, "create", "--ssl")

	boa.ShellExecE("rm -f cluster-test-cluster.sh")
}
