.PHONY: help all create delete bootstrap flux deploy check clean app webv test load-test jumpbox

help :
	@echo "Usage:"
	@echo "   make all              - create and bootstrap a cluster and deploy the apps"
	@echo "   make create           - create a k3d cluster"
	@echo "   make delete           - delete the k3d cluster"
	@echo "   make bootstrap        - deploy monitoring and logging"
	@echo "   make flux             - deploy Flux for GitOps"
	@echo "   make deploy           - deploy the apps to the cluster"
	@echo "   make check            - check the endpoints with curl"
	@echo "   make test             - run a WebValidate test"
	@echo "   make load-test        - run a 60 second WebValidate test"
	@echo "   make clean            - delete the apps from the cluster"
	@echo "   make app              - build and deploy a local app docker image"
	@echo "   make webv             - build and deploy a local WebV docker image"
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

	@k3d cluster create --registry-use k3d-registry.localhost:5500 --config k3d.yaml

	# wait for cluster to be ready
	@kubectl wait node --for condition=ready --all --timeout=30s
	@sleep 10
	@kubectl wait pod -l k8s-app=kube-dns -n kube-system --for condition=ready --timeout 30s

bootstrap :
	@# continue on most errors

	# deploy prometheus, grafana, and fluent bit
	@kubectl apply -f bootstrap
	@kubectl apply -R -f bootstrap

flux :
	@# deploy Flux for GitOps
	# Run: ./setup-flux.sh Region State City Number

deploy :
	# deploy the apps

	@# continue on most errors
	@kubectl apply -f ngsa-memory

	# deploy WebV after the app starts
	@kubectl wait pod ngsa-memory --for condition=ready --timeout=30s
	@kubectl apply -f webvalidate
	@kubectl wait pod webv --for condition=ready --timeout=30s

	# display pod status
	@kubectl get po

check :
	# curl all of the endpoints
	# Prometheus
	@curl localhost:30000

	# Grafana
	@curl localhost:32000

clean :
	# delete the deployment
	@# continue on error
	-kubectl delete -f webvalidate
	-kubectl delete -f ngsa-memory
	-kubectl delete pod jumpbox --ignore-not-found=true
	-kubectl delete ns monitoring --ignore-not-found=true
	-kubectl delete ns logging --ignore-not-found=true

	# show namespaces and running pods
	@kubectl get ns && kubectl get po -A

app :
	# build push the local image
	docker build ../ngsa-app -t k3d-registry.localhost:5500/ngsa-app:local
	docker push k3d-registry.localhost:5500/ngsa-app:local

	# delete WebV
	-kubectl delete -f webvalidate --ignore-not-found=true

	# delete/deploy the app
	-kubectl delete -f ngsa-memory --ignore-not-found=true
	kubectl apply -f ngsa-local

	# deploy WebValidate after app starts
	@kubectl wait pod ngsa-memory --for condition=ready --timeout=30s
	kubectl apply -f webvalidate
	@kubectl wait pod webv --for condition=ready --timeout=30s

	@kubectl get po

	# display the app version
	-http localhost:30080/version

webv :
	# build and push the local image
	docker build ../webvalidate -t k3d-registry.localhost:5500/webv:local
	docker push k3d-registry.localhost:5500/webv:local
	
	# delete / create WebValidate
	-kubectl delete -f webvalidate --ignore-not-found=true
	kubectl apply -f webv-local
	kubectl wait pod webv --for condition=ready --timeout=30s
	@kubectl get po

	# display the current version
	-http localhost:30088/version

test :
	# use WebValidate to run a test
	cd webv && webv --verbose --summary tsv --server http://localhost:30080 http://localhost:30081 --files baseline.json baseline.json baseline.json
	@echo
	# the 400 and 404 results are expected
	# Errors and ValidationErrorCount should both be 0

load-test :
	# use WebValidate to run a 60 second test
	cd webv && webv --server http://localhost:30080 http://localhost:30081 http://localhost:30080 http://localhost:30081 --files benchmark.json --run-loop --sleep 4 --duration 60

jumpbox :
	# start a jumpbox pod
	@-kubectl delete pod jumpbox --ignore-not-found=true

	@kubectl run jumpbox --image=ghcr.io/cse-labs/jumpbox --restart=Always
	@kubectl wait pod jumpbox --for condition=ready --timeout=30s

get-pods :
	@kubectl get pods -A
