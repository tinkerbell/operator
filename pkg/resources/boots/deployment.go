package boots

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ptr "k8s.io/utils/pointer"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func CreateDeployment(ctx context.Context, client ctrlruntimeclient.Client, ns string) error {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "boots",
			Namespace: ns,
			Labels: map[string]string{
				"app": "boots",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptr.Int32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":   "boots",
					"stack": "tinkerbell",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":   "boots",
						"stack": "tinkerbell",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "boots",
							Image:           "quay.io/tinkerbell/boots:v0.8.0",
							ImagePullPolicy: corev1.PullIfNotPresent,
							Args:            []string{"--dhcp-addr", "0.0.0.0:67"},
							Env:             parsedEnvVars(),
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceMemory: resource.MustParse("64Mi"),
									corev1.ResourceCPU:    resource.MustParse("10m"),
								},
								Limits: corev1.ResourceList{
									corev1.ResourceMemory: resource.MustParse("128Mi"),
									corev1.ResourceCPU:    resource.MustParse("500m"),
								},
							},
						},
					},
					ServiceAccountName: serviceAccountName,
					HostNetwork:        true,
				},
			},
		},
	}

	if err := client.Create(ctx, deployment); err != nil {
		if kerrors.IsAlreadyExists(err) {
			return nil
		}

		return err
	}

	return nil
}

func parsedEnvVars() []corev1.EnvVar {
	return []corev1.EnvVar{
		{
			Name: "TRUSTED_PROXIES",
			// TODO: pass the TRUSTED_PROXIES as a command line
			Value: "",
		},
		{
			Name:  "DATA_MODEL_VERSION",
			Value: "kubernetes",
		},
		{
			Name:  "FACILITY_CODE",
			Value: "lab1",
		},
		{
			Name:  "HTTP_BIND",
			Value: ":80",
		},
		{
			Name: "MIRROR_BASE_URL",
			// TODO: pass the MIRROR_BASE_URL as a command line
			Value: "",
		},
		{
			Name: "BOOTS_OSIE_PATH_OVERRIDE",
			// TODO: pass the BOOTS_OSIE_PATH_OVERRIDE as a command line
			Value: "",
		},
		{
			Name: "PUBLIC_IP",
			// TODO: pass the PUBLIC_IP as a command line
			Value: "",
		},
		{
			Name: "PUBLIC_SYSLOG_FQDN",
			// TODO: pass the PUBLIC_SYSLOG_FQDN as a command line
			Value: "",
		},
		{
			Name:  "SYSLOG_BIND",
			Value: ":514",
		},
		{
			Name: "TINKERBELL_GRPC_AUTHORITY",
			// TODO: pass the TINKERBELL_GRPC_AUTHORITY as a command line
			Value: "",
		},
		{
			Name:  "TINKERBELL_TLS",
			Value: "false",
		},
		{
			Name:  "BOOTS_LOG_LEVEL",
			Value: "debug",
		},
		{
			Name:  "BOOTS_EXTRA_KERNEL_ARGS",
			Value: "tink_worker_image=quay.io/tinkerbell/tink-worker:v0.8.0",
		},
	}
}
