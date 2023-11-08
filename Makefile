.PHONY: plan ls

# server-args='--node-taint="master:NoSchedule" --kube-apiserver-arg="default-not-ready-toleration-seconds=10" --kube-apiserver-arg="default-unreachable-toleration-seconds=10"'
server-args='--node-taint="master:NoSchedule" --kube-apiserver-arg="default-not-ready-toleration-seconds=10" --kube-apiserver-arg="default-unreachable-toleration-seconds=10"'



plan:
	k3sup plan \
	node.json \
	--server-k3s-extra-args ${server-args} \
	--user pi \
	--servers 2 \
	--background > bootstrap.sh & \
	chmod +x bootstrap.sh &

ls:
	export KUBECONFIG=`pwd`/kubeconfig \
	& kubectl get nodes

