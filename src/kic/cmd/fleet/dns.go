// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"kic/boa"
	"strings"

	"github.com/spf13/cobra"
)

// ListCmd lists the clusters in the fleet
var DnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "DNS Commands",
}

func init() {
	boa.AddScriptCommand(DnsCmd, "add", "Add a DNS entry", fltDnsAddScript())
	boa.AddScriptCommand(DnsCmd, "delete", "Delete a DNS entry", fltDnsDeleteScript())
	boa.AddScriptCommand(DnsCmd, "list", "List DNS entries", fltDnsListScript())
	boa.AddScriptCommand(DnsCmd, "show", "Get a DNS entry by host name", fltDnsShowScript())
}

func runDeleteCmd(cmd *cobra.Command, args []string) error {

	script := fltDnsDeleteScript()

	if len(args) > 0 {
		script = strings.ReplaceAll(script, "$1", args[0])
	}

	boa.ShellExecE(script)

	return nil
}

func fltDnsAddScript() string {

	return `
set -e

host=$1
pip=$2

if [ "$host" = "" ] || [ "$pip" = "" ]
then
  echo "Usage: flt dns add hostName ipAddress"
  exit 0
fi

# get the old IP
old_ip=$(az network dns record-set a list \
--query "[?name=='$host'].{IP:aRecords}" \
--resource-group "$AKDC_DNS_RG" \
--zone-name "$AKDC_SSL" \
-o json | jq -r '.[].IP[].ipv4Address' | tail -n1)

# delete old DNS entry if exists
if [ "$old_ip" != "" ] && [ "$old_ip" != "$pip" ]
then
  echo "Deleting old IP: $old_ip"
  # delete the old DNS entry
  az network dns record-set a remove-record \
  -g "$AKDC_DNS_RG" \
  -z "$AKDC_SSL" \
  -n "$host" \
  -a "$old_ip" -o table
fi

if [ "$old_ip" != "$pip" ]
then
  echo "Adding host: $host"
  # create DNS record
  az network dns record-set a add-record \
  -g "$AKDC_DNS_RG" \
  -z "$AKDC_SSL" \
  -n "$host" \
  -a "$pip" \
  --ttl 10 -o table
fi

`
}

func fltDnsDeleteScript() string {

	return `
set -e

host=$1

if [ "$host" = "" ]
then
    echo "Usage: flt dns delete hostName"
    exit 0
fi

if [ "$AKDC_DNS_RG" = "" ]
then
    echo "AKDC_DNS_RG must be set"
    exit 0
fi

if [ "$AKDC_SSL" = "" ]
then
    echo "AKDC_SSL must be set"
    exit 0
fi

echo "Deleting DNS for $host"

# delete the old DNS entry
az network dns record-set a delete \
-g "$AKDC_DNS_RG" \
-z "$AKDC_SSL" \
-n "$host" \
--yes \
-o table

`
}

func fltDnsListScript() string {

	return `
#!/bin/bash

if ! az account show -o table > /dev/null
then
    echo "Please login to Azure"
    exit 0
fi

az network dns record-set a list -g "$AKDC_DNS_RG" -z "$AKDC_SSL" --query '[].name' -o tsv | sort

`
}

func fltDnsShowScript() string {

	return `
#!/bin/bash

set -e

host=$1

if [ "$host" = "" ]
then
    echo "Usage: flt dns get hostName"
    exit 0
fi

if ! az account show -o table > /dev/null
then
    echo "Please login to Azure"
    exit 0
fi

# delete the old DNS entry
az network dns record-set a show \
-g "$AKDC_DNS_RG" \
-z "$AKDC_SSL" \
-n "$host"

`
}
