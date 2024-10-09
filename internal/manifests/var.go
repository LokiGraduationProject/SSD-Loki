package manifests

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

func commonLabels(instanceName string, component string) map[string]string {
	return map[string]string{
		"app.kubernetes.io/name":      "loki",
		"app.kubernetes.io/instance":  instanceName,
		"app.kubernetes.io/component": component,
	}
}

func memberListLabels() map[string]string {
	return map[string]string{
		"app.kubernetes.io/part-of": "memberlist",
	}
}

func headlessServiceLabels() map[string]string {
	return map[string]string{
		"variant":                       "headless",
		"prometheus.io/service-monitor": "false",
	}
}

func fqdn(serviceName, namespace string) string {
	return fmt.Sprintf("%s.%s.svc.cluster.local", serviceName, namespace)
}

const (
	lokiHTTPPortName       = "http-metrics"
	lokiGRPCPortName       = "grpc"
	lokiMemberListPortName = "http-memberlist"
	httpPort               = 3100
	grpcPort               = 9095
	memberListPort         = 7946
	protocolTCP            = corev1.ProtocolTCP
	HeadLessClusterIP      = "None"
	defaultImage           = "docker.io/grafana/loki:3.1.1"
	defaultnamespace       = "default"
)
