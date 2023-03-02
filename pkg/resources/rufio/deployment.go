package rufio

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/intstr"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ptr "k8s.io/utils/pointer"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func CreateDeployment(ctx context.Context, client ctrlruntimeclient.Client, ns string) error {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "rufio",
			Namespace: ns,
			Labels: map[string]string{
				"app":           "rufio",
				"control-plane": "controller-manager",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptr.Int32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":           "rufio",
					"control-plane": "controller-manager",
					"stack":         "tinkerbell",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"kubectl.kubernetes.io/default-container": "manager",
					},
					Labels: map[string]string{
						"app":           "rufio",
						"control-plane": "controller-manager",
						"stack":         "tinkerbell",
					},
				},
				Spec: corev1.PodSpec{
					SecurityContext: &corev1.PodSecurityContext{
						RunAsNonRoot: ptr.Bool(true),
					},
					Containers: []corev1.Container{
						{
							Name:    "manager",
							Command: []string{"/manager"},
							Image:   "quay.io/tinkerbell/rufio:v0.1.0",
							SecurityContext: &corev1.SecurityContext{
								AllowPrivilegeEscalation: ptr.Bool(false),
							},
							ImagePullPolicy: corev1.PullIfNotPresent,
							Args:            []string{"--leader-elect"},
							Env: []corev1.EnvVar{
								{
									Name: "HEGEL_TRUSTED_PROXIES",
									// TODO: pass the TRUSTED_PROXIES as a command line
									Value: "10.244.0.0/24,10.244.1.0/24,10.244.2.0/24",
								},
							},
							LivenessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/healthz",
										Port: intstr.IntOrString{IntVal: 8081},
									},
								},
								InitialDelaySeconds: 15,
								PeriodSeconds:       20,
							},
							ReadinessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/readyz",
										Port: intstr.IntOrString{IntVal: 8081},
									},
								},
								InitialDelaySeconds: 5,
								PeriodSeconds:       10,
							},
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
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: int32(50061),
									Name:          "rufio-http",
								},
							},
						},
					},
					ServiceAccountName:            serviceAccountName,
					TerminationGracePeriodSeconds: ptr.Int64(10),
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
