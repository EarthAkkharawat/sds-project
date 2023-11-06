plan:
	k3sup plan \
	node.json \
	--user pi \
	--servers 2 \
	--background > bootstrap.sh & \
	chmod +x bootstrap.sh &

ls:
	export KUBECONFIG=`pwd`/kubeconfig \
	& kubectl get nodes

