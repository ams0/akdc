.PHONY: docker gitops all

docker :
	@# force docker to rebuild
	@date > /workspaces/ngsa-app/version
	@docker build /workspaces/ngsa-app -t ghcr.io/cse-labs/aioa\:beta

	@# force GitOps to run
	@docker run --rm ghcr.io/cse-labs/aioa\:beta --version > version

	@#docker login ghcr.io -u bartr -p ${AKDC_PAT}
	@docker push ghcr.io/cse-labs/aioa\:beta

	# build complete

all : docker gitops

gitops :

	@# run GitOps
	@kic gitops

	@echo ""
	@echo ""
	# force flux sync
	@echo ""
	@kic fleet sync
	@echo ""
	@echo ""
	@sleep 2

	# check pod status for restart
	@kic fleet exec "kubectl get pods -n aioa-beta"
	@echo ""
	@echo ""
	@sleep 2

	@kic fleet exec "kubectl get pods -n aioa-beta"
	@echo ""
	@echo ""
	@sleep 2

	@kic fleet exec "kubectl get pods -n aioa-beta"
	@echo ""
	@echo ""

	# GitOps complete
	@echo ""
	# check status command
	#    kic fleet exec "kubectl get pods -n aioa-beta"
	@echo ""
