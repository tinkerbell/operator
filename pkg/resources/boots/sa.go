package boots

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	serviceAccountName = "boots"
	roleName           = "boots"
	roleBindingName    = "boots"
)

func CreateServiceAccount(ctx context.Context, client ctrlruntimeclient.Client, ns string) error {
	sa := &rbacv1.ClusterRole{
		ObjectMeta: v1.ObjectMeta{
			Name: serviceAccountName,
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
