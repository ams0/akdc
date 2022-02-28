#!/bin/bash

curl https://central-mo-kc-101.cseretail.com/heartbeat/17; echo "  central-mo-kc-101" &
curl https://central-mo-kc-102.cseretail.com/heartbeat/17; echo "  central-mo-kc-102" &
curl https://central-tx-austin-101.cseretail.com/heartbeat/17; echo "  central-tx-austin-101" &
curl https://central-tx-austin-102.cseretail.com/heartbeat/17; echo "  central-tx-austin-102" &
curl https://east-ga-atlanta-101.cseretail.com/heartbeat/17; echo "  east-ga-atlanta-101" &
curl https://east-ga-atlanta-102.cseretail.com/heartbeat/17; echo "  east-ga-atlanta-102" &
curl https://east-nc-raleigh-101.cseretail.com/heartbeat/17; echo "  east-nc-raleigh-101" &
curl https://east-nc-raleigh-102.cseretail.com/heartbeat/17; echo "  east-nc-raleigh-102" &
curl https://west-ca-sd-101.cseretail.com/heartbeat/17; echo "  west-ca-sd-101" &
curl https://west-ca-sd-102.cseretail.com/heartbeat/17; echo "  west-ca-sd-102" &
curl https://west-wa-seattle-101.cseretail.com/heartbeat/17; echo "  west-wa-seattle-101" &
curl https://west-wa-seattle-102.cseretail.com/heartbeat/17; echo "  west-wa-seattle-102" &

curl https://corp-monitoring.cseretail.com/heartbeat/17; echo "  corp-monitoring" &

curl https://west-wa-redmond-101.cseretail.com/heartbeat/17; echo "  west-wa-redmond-101" &
curl https://west-wa-redmond-102.cseretail.com/heartbeat/17; echo "  west-wa-redmond-102" &
