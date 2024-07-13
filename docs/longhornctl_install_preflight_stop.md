## longhornctl install preflight stop

Stop Longhorn preflight installer

### Synopsis

This command terminates the preflight installer.

```
longhornctl install preflight stop [flags]
```

### Options

```
  -h, --help                      help for stop
      --image string              Image containing longhornctl-local (default "longhornio/longhorn-cli:v1.8.0-dev")
      --kube-config string        Kubernetes config (kubeconfig) path
  -l, --log-level string          Log level (default "info")
      --operating-system string   Specify the operating system ("", cos). Leave this empty to use the package manager for installation.
```

### SEE ALSO

* [longhornctl install preflight](longhornctl_install_preflight.md)	 - Install Longhorn preflight

###### Auto generated by spf13/cobra on 15-Jul-2024