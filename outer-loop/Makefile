.PHONY: help all create delete bootstrap flux check test load-test jumpbox reconcile

help :
	@echo "Usage:"
	@echo "   make all              - create and bootstrap a cluster and deploy the apps"
	@echo "   make create           - create a k3d cluster"
	@echo "   make delete           - delete the k3d cluster"
	@echo "   make bootstrap        - deploy monitoring and logging"
	@echo "   make flux             - deploy Flux for GitOps"
	@echo "   make reconcile        - force Flux reconcile"
	@echo "   make check            - check the endpoints with curl"
	@echo "   make test             - run a WebValidate test"
	@echo "   make load-test        - run a 60 second WebValidate test"
	@echo "   make jumpbox          - deploy a 'jumpbox' pod"

all : create jumpbox bootstrap get-pods

delete :
	# delete the cluster (if exists)
	@# this will fail harmlessly if the cluster does not exist
	@k3d cluster delete

create : delete
	@# create the cluster and wait for ready
	@# this will fail harmlessly if the cluster exists
	@# default cluster name is k3d

	@k3d cluster create --registry-use k3d-registry.localhost:5500 --config ../deploy/k3d.yaml

	# wait for cluster to be ready
	@kubectl wait node --for condition=ready --all --timeout=30s
	@sleep 10
	@kubectl wait pod -l k8s-app=kube-dns -n kube-system --for condition=ready --timeout 30s

bootstrap :
	@# continue on most errors

	# deploy prometheus, grafana, and fluent bit
	@kubectl apply -f ../deploy/bootstrap
	@kubectl apply -R -f ../deploy/bootstrap

flux :
	@# deploy Flux for GitOps
	# Run: ./setup-flux.sh Region State City Number

reconcile :
	@flux reconcile source git GitOps

check :
	# curl all of the endpoints
	# ngsa
	@-curl localhost:30080/version && echo "" && echo ""

	# WebV
	@-curl localhost:30088/version && echo "" && echo ""

	# Prometheus
	@-curl localhost:30000

	# Grafana
	@-curl localhost:32000

test :
	# use WebValidate to run a test
	@cd ../webv && webv --verbose --summary tsv --server http://localhost:30080 --files baseline.json baseline.json baseline.json
	@echo
	# the 400 and 404 results are expected
	# Errors and ValidationErrorCount should both be 0

load-test :
	# use WebValidate to run a 60 second test
	@cd ../webv && webv --verbose --server http://localhost:30080 --files benchmark.json --run-loop --sleep 100 --duration 60

jumpbox :
	# start a jumpbox pod
	@-kubectl delete pod jumpbox --ignore-not-found=true

	@kubectl run jumpbox --image=ghcr.io/cse-labs/jumpbox --restart=Always
	@kubectl wait pod jumpbox --for condition=ready --timeout=30s

get-pods :
	@kubectl get pods -A
