package consts

const (
	AppNamePreflightChecker              = "longhorn-preflight-checker"
	AppNamePreflightContainerOptimizedOS = "longhorn-gke-cos-node-agent"
	AppNamePreflightInstaller            = "longhorn-preflight-installer"

	AppNameCoreDNS = "coredns"
)

type DependencyModuleType int

const (
	DependencyModuleDefault DependencyModuleType = iota
	DependencyModuleSpdk
)
