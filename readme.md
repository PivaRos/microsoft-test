# Microsoft Test

<span style="font-size:12px;font-weight:600;">Made by Daniel Gurbin</span>

- Assuming you have minikube and helm installed

### 1. Before running helm install run the following commands to :

- enable ingress
- enable metrics-server
- add certificate manager

```bash
    minikube addons enable ingress
    minikube addons enable metrics-server
    helm repo add jetstack https://charts.jetstack.io
    helm repo update
    helm install cert-manager jetstack/cert-manager --namespace cert-manager --create-namespace --set installCRDs=true
```

### 2. Edit the values file

change the variables as you please

```yaml
image:
  repository: <image-name> # pivaros/service-b or pivaros/service-a
  tag: "latest"

ingress:
  host: <domain-name>
  route: <route-name>
```

### 4. Apply read-only cluster-role first

```bash
kubectl apply -f ./service-chart/cluster-role.yaml
```

### 3. Run helm install

```bash
cd service-chart
helm install <release-name> . -f values.yaml --create-namespace
```

<span style="color: green; font-size:22px;font-weight:600;">Done !</span>

## Finally when both services are up and running you can check the roles by running :

```bash
kubectl auth can-i get pods -n a --as=system:serviceaccount:a:admin-sa
```

to check that the a's admin service-account can access the pods in namespace a

```bash
kubectl auth can-i get pods -n b --as=system:serviceaccount:a:admin-sa
```

to check that the a's admin service-account can access the pods in namespace b

```bash
kubectl auth can-i delete pods -n b --as=system:serviceaccount:a:admin-sa
```

to check that the a's admin service-account cannot delete pods in namespace b

and so on ..
