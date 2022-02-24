# description to potentially read at runtime
kicModule
	check
	Check status on the local k3d cluster

kicCommand
	flux
	Check Flux status on the local cluster
	flux get kustomization

kicCommand
	heartbeat
	Check heartbeat service on the local cluster
	http https://$AKDC_FQDN/heartbeat/17

kicCommand
	logs
	Check Flux status on the local cluster
	cat /var/log/cloud-init-output.log

kicCommand
	ngsa
	Check NGSA status on the local cluster
	kubectl get po -n ngsa | grep ngsa

kicCommand
	setup
	Check setup status on the local cluster
	cat ~/status

kicCommand
	webv
	Check WebV status on the local cluster
	kubectl get po -n ngsa | grep webv
