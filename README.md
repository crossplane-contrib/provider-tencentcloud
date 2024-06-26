# Provider Tencentcloud

`provider-tencentcloud` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Tencentcloud API.

<div>
  <p>
    <a href="https://cloud.tencent.com">
        <img src=".github/01_Tcloud_logo_Eng.png" alt="logo" title="Terraform" height="69">
    </a>
    <br>
    <i>Tencent Infrastructure Automation for Crossplane.</i>
    <br>
  </p>
</div>

* Wechat Group:

    <img src=".github/02_Tcloud_wechat.jpg" width="200"/>

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/crossplane-contrib/provider-tencentcloud):
```
up ctp provider install crossplane-contrib/provider-tencentcloud:latest
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-tencentcloud
spec:
  package: crossplane-contrib/provider-tencentcloud:latest
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/crossplane-contrib/provider-tencentcloud).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane-contrib/provider-tencentcloud/issues).
