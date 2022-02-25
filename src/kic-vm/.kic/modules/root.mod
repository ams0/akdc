# root commands with no sub-commands
kicRunCommand
	clone
	Clone the gitops repo to ~/gitops
	clone

kicCommand
	sync
	Force Flux to sync (reconcile) to the local cluster
	flux reconcile source git gitops && kubectl get pods -A
