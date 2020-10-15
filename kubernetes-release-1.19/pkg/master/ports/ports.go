package ports

import (
	"k8s.io/cloud-provider"
)

const (
	// ProxyStatusPort is the default port for the proxy metrics server.
	// May be overridden by a flag at startup.
	ProxyStatusPort = 10249
	// KubeletPort is the default port for the kubelet server on each host machine.
	// May be overridden by a flag at startup.
	KubeletPort = 10250
	// InsecureKubeControllerManagerPort is the default port for the controller manager status server.
	// May be overridden by a flag at startup.
	// Deprecated: use the secure KubeControllerManagerPort instead.
	InsecureKubeControllerManagerPort = 10252
	// KubeletReadOnlyPort exposes basic read-only services from the kubelet.
	// May be overridden by a flag at startup.
	// This is necessary for heapster to collect monitoring stats from the kubelet
	// until heapster can transition to using the SSL endpoint.
	// TODO(roberthbailey): Remove this once we have a better solution for heapster.
	KubeletReadOnlyPort = 10255
	// ProxyHealthzPort is the default port for the proxy healthz server.
	// May be overridden by a flag at startup.
	ProxyHealthzPort = 10256
	// KubeControllerManagerPort is the default port for the controller manager status server.
	// May be overridden by a flag at startup.
	KubeControllerManagerPort = 10257
	// CloudControllerManagerPort is the default port for the cloud controller manager server.
	// This value may be overridden by a flag at startup.
	CloudControllerManagerPort = cloudprovider.CloudControllerManagerPort
)
