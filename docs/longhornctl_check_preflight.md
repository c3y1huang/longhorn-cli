## longhornctl check preflight

Check Longhorn preflight

### Synopsis

This command verifies your Kubernetes cluster environment. It performs a series of checks to ensure your cluster meets the requirements for Longhorn to function properly.
These checks can help to identify issues that might prevent Longhorn from functioning properly.

```
longhornctl check preflight [flags]
```

### Options

```
      --enable-spdk          Enable checking of SPDK required packages, modules, and setup.
  -h, --help                 help for preflight
      --huge-page-size int   Specify the huge page size in MiB for SPDK. (default 2048)
      --image string         Image containing longhornctl-local (default "longhornio/longhorn-cli:v1.8.0-dev")
      --kube-config string   Kubernetes config (kubeconfig) path
  -l, --log-level string     Log level (default "info")
      --uio-driver string    User space I/O driver for SPDK. (default "uio_pci_generic")
```

### SEE ALSO

* [longhornctl check](longhornctl_check.md)	 - Longhorn checking operations

###### Auto generated by spf13/cobra on 15-Jul-2024
