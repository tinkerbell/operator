package tink

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
			Name:      "tink-server",
			Namespace: ns,
			Labels: map[string]string{
				"app": "tink-server",
			},
		},
		Spec: corev1.ServiceSpec{
			ClusterIP: "None",
			Selector: map[string]string{
				"app": "tink-server",
			},
			Ports: []corev1.ServicePort{
				{
					Port:       42113,
					TargetPort: intstr.IntOrString{StrVal: "tink-grpc"},
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
