package boots

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	clusterRole        = "boots-cluster-role"
	clusterRoleBinding = "boots-cluster-role-binding"
)

func CreateClusterRole(ctx context.Context, client ctrlruntimeclient.Client) error {
	clusterRole := &rbacv1.ClusterRole{
		ObjectMeta: v1.ObjectMeta{
			Name: clusterRole,
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{"tinkerbell.org"},
				Resources: []string{"hardware", "hardware/status"},
				Verbs:     []string{"get", "list", "watch"},
			},
			{
				APIGroups: []string{"tinkerbell.org"},
				Resources: []string{"workflows", "workflows/status"},
				Verbs:     []string{"get", "list", "watch"},
			},
		},
	}

	if err := client.Create(ctx, clusterRole); err != nil {
		if kerrors.IsAlreadyExists(err) {
			return nil
		}

		return err
	}

	return nil
}

func CreateClusterRoleBinding(ctx context.Context, client ctrlruntimeclient.Client, ns string) error {
	clusterRoleBinding := &rbacv1.ClusterRoleBinding{
		ObjectMeta: v1.ObjectMeta{
			Name: clusterRoleBinding,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      rbacv1.ServiceAccountKind,
				Name:      serviceAccountName,
				Namespace: ns,
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: rbacv1.GroupName,
			Kind:     "ClusterRole",
			Name:     clusterRole,
		},
	}

	if err := client.Create(ctx, clusterRoleBinding); err != nil {
		if kerrors.IsAlreadyExists(err) {
			return nil
		}

		return err
	}

	return nil
}
