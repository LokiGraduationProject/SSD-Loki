package manifests

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func buildLokiMinioSecret() *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "loki-minio",
			Namespace: "default",
			Labels: map[string]string{
				"app": "minio",
			},
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{
			"rootUser":     []byte("minio"),
			"rootPassword": []byte("minio123"),
		},
	}
}
