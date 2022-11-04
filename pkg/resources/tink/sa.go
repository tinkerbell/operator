package tink

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	tinkServerServiceAccountName = "tink-server"

	tinkControllerServiceAccountName = "tink-controller"
)

func CreateTinkServerServiceAccount(ctx context.Context, client ctrlruntimeclient.Client, ns string) error {
	sa := &corev1.ServiceAccount{
		ObjectMeta: v1.ObjectMeta{
			Name:      tinkServerServiceAccountName,
			Namespace: ns,
		},
	}

	if err := client.Create(ctx, sa); err != nil {
		if kerrors.IsAlreadyExists(err) {
			return nil
		}

		return err
	}

	return nil
}

func CreateTinkControllerServiceAccount(ctx context.Context, client ctrlruntimeclient.Client, ns string) error {
	sa := &corev1.ServiceAccount{
		ObjectMeta: v1.ObjectMeta{
			Name:      tinkControllerServiceAccountName,
			Namespace: ns,
		},
	}

	if err := client.Create(ctx, sa); err != nil {
		if kerrors.IsAlreadyExists(err) {
			return nil
		}

		return err
	}

	return nil
}
