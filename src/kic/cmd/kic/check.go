// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package kic

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// ListCmd lists the clusters in the fleet
var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check status on the local cluster",
}

func init() {
	boa.AddScriptCommand(CheckCmd, "all", "Run all status checks on the local cluster", kicCheckAllScript())
	boa.AddScriptCommand(CheckCmd, "heartbeat", "Check heartbeat status on the local cluster", kicCheckHeartbeatScript())
	boa.AddScriptCommand(CheckCmd, "imdb", "Check IMDb reference app status on the local cluster", kicCheckImdbScript())
	boa.AddScriptCommand(CheckCmd, "webv", "Check WebV status on the local cluster", kicCheckWebvScript())
	boa.AddScriptCommand(CheckCmd, "grafana", "Check Grafana status on the local cluster", kicCheckGrafanaScript())
	boa.AddScriptCommand(CheckCmd, "prometheus", "Check Prometheus status on the local cluster", kicCheckPrometheusScript())
}

func kicCheckAllScript() string {

	return `
#!/bin/bash

echo "Checking Heartbeat"
kic check heartbeat

echo "Checking IMDb"
kic check imdb

echo "Checking WebV"
kic check webv

echo "Checking Grafana"
kic check grafana

echo "Checking Prometheus"
kic check prometheus

`
}

func kicCheckHeartbeatScript() string {
	return "http localhost:31080/heartbeat/17"
}

func kicCheckImdbScript() string {
	return "http localhost:30080/version"
}

func kicCheckWebvScript() string {

	return "kubectl exec -it jumpbox -- http webv.imdb.svc.cluster.local:8080/version"
}

func kicCheckGrafanaScript() string {

	return "http localhost:32000/healthz"
}

func kicCheckPrometheusScript() string {

	return "http localhost:30000"
}
