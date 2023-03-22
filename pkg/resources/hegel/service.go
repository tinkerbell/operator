package hegel

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
			Name:      "hegel",
			Namespace: ns,
			Labels: map[string]string{
				"app": "hegel",
			},
		},
		Spec: corev1.ServiceSpec{
			ClusterIP: "None",
			Selector: map[string]string{
				"app": "hegel",
			},
			Ports: []corev1.ServicePort{
				{
					Port:       50061,
					TargetPort: intstr.IntOrString{IntVal: 50061},
					Protocol:   corev1.ProtocolTCP,
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
