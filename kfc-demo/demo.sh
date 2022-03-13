#!/bin/bash
clear
akdc fleet check ai-order-accuracy | sort
echo ""
echo "checking health ..."

x=1
while curl -s "https://west-co-denver-$1.cseretail.com/healthz" | grep -v 404
do
    sleep 1
    x=$(( x + 1 ))

    if [ $x -gt 5 ]
    then
        break
    fi
done

echo ""
echo "App failed: starting failover"
echo ""

cd /workspaces/edge-gitops || exit

git pull > /dev/null

if grep 101 < apps/failover/autogitops/autogitops.json
then
    sed -i "s/101/102/g" apps/failover/autogitops/autogitops.json
else
    sed -i "s/102/101/g" apps/failover/autogitops/autogitops.json
fi

git commit apps/failover -m "testing failover" > /dev/null

git push > /dev/null

cd "$OLDPWD" || exit

akdc fleet sync
