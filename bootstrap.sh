#!/bin/sh

echo "Setting up primary server 1"
k3sup install --host 192.168.0.104 \
--user pi \
--cluster \
--local-path kubeconfig \
--context default \
--k3s-extra-args '--node-taint="key1=value1:NoSchedule" --kube-apiserver-arg="default-not-ready-toleration-seconds=10" --kube-apiserver-arg="default-unreachable-toleration-seconds=10"'

echo "Fetching the server's node-token into memory"

export NODE_TOKEN=$(k3sup node-token --host 192.168.0.104 --user pi)

echo "Setting up additional server: 2"
k3sup join \
--host 192.168.0.105 \
--server-host 192.168.0.104 \
--node-token "$NODE_TOKEN" \
--server-user earth \
--user earth \
--k3s-extra-args '--node-taint="key1=value1:NoSchedule" --kube-apiserver-arg="default-not-ready-toleration-seconds=10" --kube-apiserver-arg="default-unreachable-toleration-seconds=10"' &

echo "Setting up worker: 1"
k3sup join \
--host 192.168.0.100 \
--server-host 192.168.0.104 \
--node-token "$NODE_TOKEN" \
--server-user earth \
--user pi &

echo "Setting up worker: 2"
k3sup join \
--host 192.168.0.101 \
--server-host 192.168.0.104 \
--node-token "$NODE_TOKEN" \
--server-user earth \
--user pi &

echo "Setting up worker: 3"
k3sup join \
--host 192.168.0.102 \
--server-host 192.168.0.104 \
--node-token "$NODE_TOKEN" \
--server-user earth \
--user pi &

echo "Setting up worker: 4"
k3sup join \
--host 192.168.0.103 \
--server-host 192.168.0.104 \
--node-token "$NODE_TOKEN" \
--server-user earth \
--user pi &

