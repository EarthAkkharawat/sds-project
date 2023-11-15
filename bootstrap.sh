#!/bin/sh

echo "Setting up primary server 1"
k3sup install --host 192.168.0.121 \
--user earth \
--cluster \
--local-path kubeconfig \
--context default \
--k3s-extra-args "--node-taint "master:NoSchedule" --kube-apiserver-arg "default-not-ready-toleration-seconds=10" --kube-apiserver-arg "default-unreachable-toleration-seconds=10""

echo "Fetching the server's node-token into memory"

export NODE_TOKEN=$(k3sup node-token --host 192.168.0.121 --user earth)

echo "Setting up additional server: 2"
k3sup join \
--host 192.168.0.122 \
--server-host 192.168.0.121 \
--server \
--node-token "$NODE_TOKEN" \
--server-user earth \
--user earth \
--k3s-extra-args "--node-taint "master:NoSchedule" --kube-apiserver-arg "default-not-ready-toleration-seconds=10" --kube-apiserver-arg "default-unreachable-toleration-seconds=10"" &

echo "Setting up additional server: 3"
k3sup join \
--host 192.168.0.123 \
--server-host 192.168.0.121 \
--server \
--node-token "$NODE_TOKEN" \
--server-user earth \
--user earthakkharawat \
--k3s-extra-args "--node-taint "master:NoSchedule" --kube-apiserver-arg "default-not-ready-toleration-seconds=10" --kube-apiserver-arg "default-unreachable-toleration-seconds=10"" &

echo "Setting up worker: 1"
k3sup join \
--host 192.168.0.100 \
--server-host 192.168.0.121 \
--node-token "$NODE_TOKEN" \
--server-user earth \
--user pi &

echo "Setting up worker: 2"
k3sup join \
--host 192.168.0.101 \
--server-host 192.168.0.121 \
--node-token "$NODE_TOKEN" \
--server-user earth \
--user pi &

echo "Setting up worker: 3"
k3sup join \
--host 192.168.0.102 \
--server-host 192.168.0.121 \
--node-token "$NODE_TOKEN" \
--server-user earth \
--user pi &

echo "Setting up worker: 4"
k3sup join \
--host 192.168.0.103 \
--server-host 192.168.0.121 \
--node-token "$NODE_TOKEN" \
--server-user earth \
--user pi &

