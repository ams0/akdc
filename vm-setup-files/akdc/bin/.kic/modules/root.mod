# description to potentially read at runtime
kicRunCommand
	clone
	Clone the gitops repo to ~/gitops
	clone

kicRunCommand
	create
	Create a new local k3d cluster
	create

kicRunCommand
	delete
	Delete the k3d cluster
	delete

kicRunCommand
	jumpbox
	Deploy a 'jumpbox' to the local k3d cluster
	jumpbox

kicCommand
	sync
	Force Flux to sync (reconcile) to the local cluster
	flux reconcile source git gitops && kubectl get pods -A
