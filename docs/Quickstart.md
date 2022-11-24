# Quickstart
## Prerequisites
This quickstart requires:
- a Kubernetes cluster with permissions to create pods and secrets
- a host with kubectl installed and configured to access the Kubernetes cluster
- a TencentCloud account with permissions to create a VPC
- TencentCloud access keys

## Guided tour
### [Install Helm](https://helm.sh/docs/intro/install/)


### Install Crossplane
```
kubectl create namespace crossplane-system
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update

helm install crossplane --namespace crossplane-system crossplane-stable/crossplane
```

### Install Crossplane CLI
The Crossplane CLI extends kubectl with functionality to build, push, and install [Crossplane packages]:

```
curl -sL https://raw.githubusercontent.com/crossplane/crossplane/master/install.sh | sh

```

### Install Configuration Package
Fetch lasted VERSION in [DockerHub](https://hub.docker.com/r/crossplanecontrib/provider-tencentcloud/tags)
```
kubectl crossplane install provider crossplanecontrib/provider-tencentcloud:${VERSION}
```

### Create a Provider Secret
```
apiVersion: v1
kind: Secret
metadata:
  name: example-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "secret_id": "secret_id",
      "secret_key": "secret_key",
      "region": "region"
    }

```
Apply this configuration with kubectl apply -f

### Configure the Provider
We will create the following ProviderConfig object to configure credentials for TencentCloud Provider:

```
apiVersion: tencentcloud.crossplane.io/v1alpha1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: example-creds
      namespace: crossplane-system
      key: credentials
```
Apply this configuration with kubectl apply -f

### Create a managed resource
Create a managed resource to verify the provider is functioning.
This example creates an TencentCloud VPC.


```
apiVersion: vpc.tencentcloud.crossplane.io/v1alpha1
kind: VPC
metadata:
  name: example-vpc
spec:
  forProvider:
    cidrBlock: 10.1.0.0/16
    name: test-crossplane-vpc
```

Apply this configuration with kubectl apply -f, then use kubectl get VPC to verify VPC creation.

```
$ kubectl get VPC
NAME          READY   SYNCED   EXTERNAL-NAME   AGE
example-vpc   True    True     vpc-4c1o7d49    5d5h
```
