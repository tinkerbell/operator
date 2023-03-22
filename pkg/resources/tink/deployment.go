package tink

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ptr "k8s.io/utils/pointer"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var hostPathType = corev1.HostPathDirectoryOrCreate

func CreateTinkControllerDeployment(ctx context.Context, client ctrlruntimeclient.Client, ns string) error {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tink-controller",
			Namespace: ns,
			Labels: map[string]string{
				"app": "tink-controller",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptr.Int32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "tink-controller",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "tink-controller",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "tink-controller",
							Image:           "quay.io/tinkerbell/tink-controller:v0.8.0",
							ImagePullPolicy: corev1.PullIfNotPresent,
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
					ServiceAccountName: tinkControllerServiceAccountName,
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

func CreateTinkServerDeployment(ctx context.Context, client ctrlruntimeclient.Client, ns string) error {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tink-server",
			Namespace: ns,
			Labels: map[string]string{
				"app": "tink-server",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptr.Int32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":   "tink-server",
					"stack": "tinkerbell",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":   "tink-server",
						"stack": "tinkerbell",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "server",
							Image: "quay.io/tinkerbell/tink:v0.8.0",
							Args:  []string{"--backend", "kubernetes"},
							Env: []corev1.EnvVar{
								{
									Name:  "TINKERBELL_TLS",
									Value: "false",
								},
							},
							ImagePullPolicy: corev1.PullIfNotPresent,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: int32(42113),
									Name:          "tink-grpc",
								},
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
						},
					},
					ServiceAccountName: tinkServerServiceAccountName,
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

func CreateTinkStackDeployment(ctx context.Context, client ctrlruntimeclient.Client, ns string) error {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tink-stack",
			Namespace: ns,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptr.Int32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "tink-stack",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						// TODO: pass the checksum/config as a command line
						"checksum/config": "75fffb14e7848a2319212c0422af0eb693157e9359a7cd10b59518125ad9822a",
					},
					Labels: map[string]string{
						"app": "tink-stack",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "tink-stack",
							Image:           "nginx:1.23.1",
							ImagePullPolicy: corev1.PullIfNotPresent,
							Command:         []string{"/bin/bash", "-c"},
							Args:            []string{"nginx -g 'daemon off;'"},
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: int32(67),
									Name:          "boots-dhcp",
									Protocol:      corev1.ProtocolUDP,
								},
								{
									ContainerPort: int32(80),
									Name:          "boots-http",
									Protocol:      corev1.ProtocolTCP,
								},
								{
									ContainerPort: int32(69),
									Name:          "boots-tftp",
									Protocol:      corev1.ProtocolUDP,
								},
								{
									ContainerPort: int32(514),
									Name:          "boots-syslog",
									Protocol:      corev1.ProtocolUDP,
								},
								{
									ContainerPort: int32(50061),
									Name:          "hegel-http",
									Protocol:      corev1.ProtocolTCP,
								},
								{
									ContainerPort: int32(42113),
									Name:          "tink-grpc",
									Protocol:      corev1.ProtocolTCP,
								},
								{
									ContainerPort: int32(8080),
									Name:          "hook-http",
									Protocol:      corev1.ProtocolTCP,
								},
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
							VolumeMounts: []corev1.VolumeMount{
								{
									MountPath: "/usr/share/nginx/html",
									Name:      "hook-artifacts",
								},
								{
									MountPath: "/tmp",
									ReadOnly:  true,
									Name:      "nginx-conf",
								},
							},
						},
					},
					InitContainers: []corev1.Container{
						{
							Name:    "init-hook-download",
							Image:   "alpine",
							Command: []string{"/bin/sh", "-xc"},
							Args: []string{
								"rm -rf /usr/share/nginx/html/checksums.txt;",
								"touch /usr/share/nginx/html/checksums.txt;",
								"echo \"7c35042d35c003ae1f424e503ad6edf21854bc70b24b37006e810c3c8a92543420eed129c14e364769b0f32c27bdf4c61299fce8f8156af7477cac6a43931a20  vmlinuz-x86_64\" >> /usr/share/nginx/html/checksums.txt;",
								"echo \"be7c3d57e2d73bfa4e41a2b5740c722b1c83722e4388b3cff9017192fce43ede360221e3095c800e511d7b4bce6065f2906883421409dd6d983412418a8d903e  initramfs-x86_64\" >> /usr/share/nginx/html/checksums.txt;",
								"echo \"2f1bdbf64380e281288f54c6ddd29221d8a007d29b40f405da0592ed32ef6e52695fc5071e05b2db3f075122943d62a2c266704d154a16ffb7b278c70538e7da  vmlinuz-aarch64\" >> /usr/share/nginx/html/checksums.txt;",
								"echo \"5adc51798c8699f5f257599aabb999e2c2f65a07c9f8607c65510e57122b3e5c53196819e7ececdcda7b8fef47ba597ea7c4b53f2f4a92e236b20e355443eefe  initramfs-aarch64\" >> /usr/share/nginx/html/checksums.txt;",
								"cd /usr/share/nginx/html/",
								"sha512sum -c, checksums.txt && exit 0",
								"apk add wget",
								"echo downloading HOOK...",
								"wget -O /tmp/hook0.tar.gz https: //github.com/tinkerbell/hook/releases/download/v0.7.0/hook_x86_64.tar.gz;",
								"tar -zxvf /tmp/hook0.tar.gz -C \"/usr/share/nginx/html/\"",
								"rm -rf /tmp/hook0.tar.gz",
								"apk add wget",
								"echo downloading HOOK...",
								"wget -O /tmp/hook1.tar.gz https://github.com/tinkerbell/hook/releases/download/v0.7.0/hook_aarch64.tar.gz;",
								"tar -zxvf /tmp/hook1.tar.gz -C \"/usr/share/nginx/html/\"",
								"rm -rf /tmp/hook1.tar.gz",
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									MountPath: "/usr/share/nginx/html",
									Name:      "hook-artifacts",
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "hook-artifacts",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/opt/hook",
									Type: &hostPathType,
								},
							},
						},
						{
							Name: "nginx-conf",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "nginx-conf",
									},
									Items: []corev1.KeyToPath{
										{
											Key:  "nginx.conf",
											Path: "nginx.conf.template",
										},
									},
								},
							},
						},
					},
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
