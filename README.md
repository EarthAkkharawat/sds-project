# sds-project

Kubernetes cluster at the edge deployed on Raspberry Pi, utilizing the lightweight [k3s](https://k3s.io/), and orchestrated with the assistance of [k3sup](https://github.com/alexellis/k3sup)

## Contents

- [List of hardwares](#list-of-hardwares-)
- [Prerequisites](#prerequisites-)
  - [Install k3sup](#install-k3sup)
  - [Config static IP via DHCP](#config-static-ip-via-dhcp)
  - [VM configuration](#vm-configuration)
  - [Raspberry Pi configuration](#raspberry-pi-configuration)
    - [hostname](#hostname)
    - [bootup](#bootup)
    - [Sudo with no password](#sudo-with-no-password)
  - [SSH-key management](#ssh-key-management)
- [Provision Cluster](#provision-cluster-)
- [Debugging or Troubleshooting](#debugging-or-troubleshooting-)
  - [View logs output](#view-logs-output)
  - [Uninstalling](#uninstalling)

## List of hardwares 🖥️

- Master Nodes

  - 2x Ubuntu 22.04 live-server installed VM

- Worker Nodes

  - 4x Raspberry Pi 4

- Networking

  - TL-WR841N Router
  - 4x LAN Cable

## Prerequisites 📝

### Install k3sup

via curl

```bash
curl -sLS https://get.k3sup.dev | sh
sudo install k3sup /usr/local/bin/
```

via brew

```bash
brew install k3sup
```

### Config static IP via DHCP

In these project, I setup router as [WISP](https://en.wikipedia.org/wiki/Wireless_distribution_system) mode, for receiving internet hotspot from my phone, and distribute it to all nodes, so that they can access internet.

Following this [Documents](https://www.tp-link.com/us/user-guides/tl-wr841n_v14/chapter-5-configure-the-router-in-wisp-mode) to do on your own.

Then, you can continue reserving static IP for each node, by matching thier MAC address to IP address.

### VM configuration

Nothing special, just make sure you have IP address on `Bridged mode` network, so they can acccess to internet and communicate with each other in the same network as Raspberry Pi.

### Raspberry Pi configuration

#### hostname

change hostname via GUI, as every node must have unique hostname.

```bash
sudo raspi-config
```

#### Sudo with no password

permit user pi to not use password when using sudo by

```bash
sudo visudo
```

then append these below lines at the end of file,

```bash
pi ALL=(ALL) NOPASSWD: ALL
```

#### bootup

Enable container features in the kernel, by editing `/boot/cmdline.txt`

Add the following to the end of the line:

```bash
cgroup_enable=cpuset cgroup_memory=1 cgroup_enable=memory
```

### SSH-key management

for coveninence, we will use ssh-copy-id to copy ssh-key to all nodes, since k3sup does not support password input or variable.

So, your need to copy all ssh-key in every instances to local, the machine that run k3sup.

```bash
ssh-copy-id user@ip
```

if it's error, then you need to generate ssh-key first.

```bash
ssh-keygen
```

## Provision Cluster 🚀

Config `node.json` likes,

```json
[
  {
    "hostname": "master1",
    "ip": "192.168.0.104"
  },
  {
    "hostname": "master2",
    "ip": "192.168.0.105"
  },
  {
    "hostname": "jindamanee",
    "ip": "192.168.0.100"
  },
  {
    "hostname": "cream",
    "ip": "192.168.0.101"
  },
  {
    "hostname": "earth",
    "ip": "192.168.0.102"
  },
  {
    "hostname": "singto",
    "ip": "192.168.0.103"
  }
]
```

Run k3sup plan via makefile.

```bash
make plan
```

Customize your `bootstrap.sh`, since k3sup plan api does not satisfy our setup, then run it.

More detail on Customizing `bootstrap.sh`, I will explain later.

```bash
./bootstrap.sh
```

If nothing failed, then copy `kubeconfig` to local, for monitoring cluster.

```bash
export KUBECONFIG=`pwd`/kubeconfig
kubectl get node
```

Taint Master Nodes for not scheduling pods on their own, since they are Ubuntu which not compatible with ARM image.

```bash
sudo kubectl taint nodes <master-hostname> master=true:NoSchedule
```

## Debugging or Troubleshooting 🔧

#### View logs output

for master nodes

```bash
sudo systemctl status k3s
```

for worker nodes

```bash
sudo systemctl status k3s-agent
```

#### Uninstalling

Uninstall server, master node

```bash
/usr/local/bin/k3s-uninstall.sh
```

Uninstall agent, worker node

```bash
/usr/local/bin/k3s-agent-uninstall.sh
```
