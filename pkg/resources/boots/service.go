package boots

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func CreateService(ctx context.Context, client ctrlruntimeclient.Client, ns string) error {
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "boots",
			Namespace: ns,
			Labels: map[string]string{
				"app": "boots",
			},
		},
		Spec: corev1.ServiceSpec{
			ClusterIP: "None",
			Selector: map[string]string{
				"app": "boots",
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "boots-dhcp",
					Port:       67,
					TargetPort: intstr.FromInt(67),
					Protocol:   corev1.ProtocolUDP,
				},
				{
					Name:       "boots-http",
					Port:       80,
					TargetPort: intstr.FromInt(80),
					Protocol:   corev1.ProtocolTCP,
				},
				{
					Name:       "boots-syslog",
					Port:       514,
					TargetPort: intstr.FromInt(514),
					Protocol:   corev1.ProtocolUDP,
				},
				{
					Name:       "boots-tftp",
					Port:       69,
					TargetPort: intstr.FromInt(69),
					Protocol:   corev1.ProtocolUDP,
				},
			},
		},
	}

	if err := client.Create(ctx, service); err != nil {
		if kerrors.IsAlreadyExists(err) {
			return nil
		}

		return err
	}

	return nil
}
