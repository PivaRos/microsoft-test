# Setup

## Macos

### Firstly install minikube and helm

```bash
brew install minikube
brew install helm
```

### Start Minikube

```bash
minikube start
```

### Domain setup

- Instead of setting up real domain I decided to mock the domain. (/hosts file)
- Since in macos docker run's on a vm and it may be tricky to setup the networking, i decided to pointed in the desired domain to localhost ip (172.0.0.1) in the host file located at /etc/host in mac

```go
127.0.0.1 production-ready.com
```

- And forward all traffic to minikube

```bash
minikube tunnel
```

## Linux

### 1. install minikube and helm

- Install the dependencies

```bash
sudo apt-get update
sudo apt-get install -y curl apt-transport-https
```

- Download minikube and add helm gpg key

```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
curl https://baltocdn.com/helm/signing.asc | sudo apt-key add -
```

- Move minikube to /usr/local/bin and add helm repo and install

```bash
sudo install minikube-linux-amd64 /usr/local/bin/minikube
sudo apt-get install apt-transport-https --yes
echo "deb https://baltocdn.com/helm/stable/debian/ all main" | sudo tee /etc/apt/sources.list.d/helm-stable-debian.list
sudo apt-get install helm

```

### Start Minikube

```bash
minikube start
```

### Domain setup

- When running on linux it can work with without running "minikube tunnel" since the docker runs on the same network as the host
- Simply point the domain to minikube's ip and you good to go !.

```go
<minikube-ip> production-ready.com
```
