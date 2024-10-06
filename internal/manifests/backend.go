package manifests

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func BuildBackend(opts Options) ([]client.Object, error) {
	statefulset := NewBackendStatefulSet(opts)
	objs := []client.Object{
		statefulset,
		NewLokiBackendService(opts),
		NewLokiBackendHeadlessService(opts),
		NewBackendPodDisruptionBudget(opts),
	}

	return objs, nil
}

func NewBackendStatefulSet(opts Options) *appsv1.StatefulSet {
	backendLabels := commonWriteLabels("backend", opts.Name)

	// 컨테이너 정의
	container := corev1.Container{
		Name:            "loki",
		Image:           defaultImage,
		ImagePullPolicy: corev1.PullIfNotPresent,
		Args: []string{
			"-config.file=/etc/loki/config/config.yaml", //TODO
			"-target=backend",
			"-legacy-read-mode=false",
		},
		Ports: []corev1.ContainerPort{
			{
				Name:          lokiHTTPPortName,
				ContainerPort: httpPort,
				Protocol:      protocolTCP,
			},
			{
				Name:          lokiGRPCPortName,
				ContainerPort: grpcPort,
				Protocol:      protocolTCP,
			},
		},
		SecurityContext: &corev1.SecurityContext{
			AllowPrivilegeEscalation: ptr.To(false),
			Capabilities: &corev1.Capabilities{
				Drop: []corev1.Capability{"ALL"},
			},
			ReadOnlyRootFilesystem: ptr.To(true),
		},
		ReadinessProbe: &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: "/ready",
					Port: intstr.IntOrString{IntVal: httpPort},
				},
			},
			InitialDelaySeconds: 30,
			TimeoutSeconds:      1,
		},
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      "config",
				MountPath: "/etc/loki/config",
			},
			{
				Name:      "runtime-config",
				MountPath: "/etc/loki/runtime-config",
			},
			{
				Name:      "data",
				MountPath: "/var/loki",
			},
		},
	}

	// PodSpec 정의
	podSpec := corev1.PodSpec{
		ServiceAccountName:            "loki",
		AutomountServiceAccountToken:  ptr.To(true),
		EnableServiceLinks:            ptr.To(true),
		TerminationGracePeriodSeconds: ptr.To(int64(300)),
		SecurityContext: &corev1.PodSecurityContext{
			FSGroup:      ptr.To(int64(10001)),
			RunAsGroup:   ptr.To(int64(10001)),
			RunAsNonRoot: ptr.To(true),
			RunAsUser:    ptr.To(int64(10001)),
		},
		Containers: []corev1.Container{container},
		Affinity: &corev1.Affinity{
			PodAntiAffinity: &corev1.PodAntiAffinity{
				RequiredDuringSchedulingIgnoredDuringExecution: []corev1.PodAffinityTerm{
					{
						LabelSelector: &metav1.LabelSelector{
							MatchLabels: map[string]string{
								"app.kubernetes.io/component": "backend",
							},
						},
						TopologyKey: "kubernetes.io/hostname",
					},
				},
			},
		},
		Volumes: []corev1.Volume{
			{
				Name: "config",
				VolumeSource: corev1.VolumeSource{
					ConfigMap: &corev1.ConfigMapVolumeSource{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: "loki",
						},
						Items: []corev1.KeyToPath{
							{
								Key:  "config.yaml",
								Path: "config.yaml",
							},
						},
					},
				},
			},
			{
				Name: "runtime-config",
				VolumeSource: corev1.VolumeSource{
					ConfigMap: &corev1.ConfigMapVolumeSource{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: "loki-runtime",
						},
					},
				},
			},
			{
				Name: "data",
				VolumeSource: corev1.VolumeSource{
					EmptyDir: &corev1.EmptyDirVolumeSource{},
				},
			},
		},
	}

	// StatefulSet 리턴
	return &appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StatefulSet",
			APIVersion: appsv1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "loki-backend",
			Namespace: "default",
			Labels:    labels.Merge(memberListLabels(), backendLabels),
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: ptr.To(int32(3)),
			Selector: &metav1.LabelSelector{
				MatchLabels: backendLabels,
			},
			ServiceName: "loki-backend-headless",
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels.Merge(memberListLabels(), backendLabels),
				},
				Spec: podSpec,
			},
			VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
				{
					TypeMeta: metav1.TypeMeta{
						APIVersion: corev1.SchemeGroupVersion.String(),
					},
					ObjectMeta: metav1.ObjectMeta{
						Name: "data",
					},
					Spec: corev1.PersistentVolumeClaimSpec{
						AccessModes: []corev1.PersistentVolumeAccessMode{
							corev1.ReadWriteOnce,
						},
						Resources: corev1.ResourceRequirements{
							Requests: corev1.ResourceList{
								corev1.ResourceStorage: resource.MustParse("10Gi"),
							},
						},
					},
				},
			},
			UpdateStrategy: appsv1.StatefulSetUpdateStrategy{
				Type: appsv1.RollingUpdateStatefulSetStrategyType,
			},
			PodManagementPolicy:  appsv1.ParallelPodManagement,
			RevisionHistoryLimit: ptr.To(int32(10)),
		},
	}
}

func NewLokiBackendService(opts Options) *corev1.Service {
	serviceName := "loki-backend"
	backendLabels := commonWriteLabels(opts.Name, "backend")

	// Return the new service object
	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: corev1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceName,
			Namespace: "default",
			Labels:    backendLabels,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeClusterIP,
			Ports: []corev1.ServicePort{
				{
					Name:       lokiHTTPPortName,
					Port:       httpPort,
					Protocol:   protocolTCP,
					TargetPort: intstr.IntOrString{StrVal: lokiHTTPPortName},
				},
				{
					Name:       lokiGRPCPortName,
					Port:       grpcPort,
					Protocol:   protocolTCP,
					TargetPort: intstr.IntOrString{StrVal: lokiGRPCPortName},
				},
			},
			Selector: backendLabels,
		},
	}
}

// NewLokiBackendHeadlessService returns a new headless service for the Loki backend.
func NewLokiBackendHeadlessService(opts Options) *corev1.Service {
	serviceName := "loki-backend-headless"
	backendLabels := commonWriteLabels(opts.Name, "backend")
	headlessServiceLabels := map[string]string{
		"variant":                       "headless",
		"prometheus.io/service-monitor": "false",
	}
	// Return the new service object
	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: corev1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceName,
			Namespace: "default",
			Labels:    labels.Merge(backendLabels, headlessServiceLabels),
		},
		Spec: corev1.ServiceSpec{
			Type:      corev1.ServiceTypeClusterIP,
			ClusterIP: "None",
			Ports: []corev1.ServicePort{
				{
					Name:       lokiHTTPPortName,
					Port:       httpPort,
					Protocol:   protocolTCP,
					TargetPort: intstr.IntOrString{StrVal: lokiHTTPPortName},
				},
				{
					Name:       lokiGRPCPortName,
					Port:       grpcPort,
					Protocol:   protocolTCP,
					TargetPort: intstr.IntOrString{StrVal: lokiGRPCPortName},
				},
			},
			Selector: backendLabels,
		},
	}
}

// NewQuerierPodDisruptionBudget returns a PodDisruptionBudget for the LokiStack querier pods.
func NewBackendPodDisruptionBudget(opts Options) *policyv1.PodDisruptionBudget {
	name := "loki-backend"
	labels := commonWriteLabels(opts.Name, "backend")

	return &policyv1.PodDisruptionBudget{
		TypeMeta: metav1.TypeMeta{
			Kind:       "PodDisruptionBudget",
			APIVersion: policyv1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: "default",
			Labels:    labels,
		},
		Spec: policyv1.PodDisruptionBudgetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			MaxUnavailable: ptr.To((intstr.FromInt(1))),
		},
	}
}
