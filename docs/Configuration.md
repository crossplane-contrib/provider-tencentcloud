### Debug and open terraform trace log
If you want to debug and open terraform trace log, you can create ProviderConfig with the following configuration file.

```
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: crossplanecontrib-provider-tencentcloud
spec:
  package: crossplanecontrib/provider-tencentcloud:<version>
  controllerConfigRef:
    name: debug-config

---

apiVersion: pkg.crossplane.io/v1alpha1
kind: ControllerConfig
metadata:
  name: debug-config
spec:
  args:
    - --debug
  env:
    -
      name: TF_LOG
      value: TRACE
    -
      name: TF_LOG_PATH
      value: /tmp/terraform-trace.log

```


### Import resource
when you need import resource, you can use `crossplane.io/external-name`, like following configuration.
```
apiVersion: vpc.tencentcloud.crossplane.io/v1alpha1
kind: VPC
metadata:
  name: example-vpc
  annotations:
    crossplane.io/external-name: "vpc-1qkhjxt3"
spec:
  forProvider:
    cidrBlock: "10.1.0.0/16"
    name: "crossplane-import"
```

