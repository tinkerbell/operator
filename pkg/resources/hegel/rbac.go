package hegel

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	role        = "hegel"
	roleBinding = "hegel-role-binding"
)

func CreateRole(ctx context.Context, client ctrlruntimeclient.Client, ns string) error {
	role := &rbacv1.Role{
		ObjectMeta: v1.ObjectMeta{
			Name:      role,
			Namespace: ns,
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

	if err := client.Create(ctx, role); err != nil {
		if kerrors.IsAlreadyExists(err) {
			return nil
		}

		return err
	}

	return nil
}

func CreateRoleBinding(ctx context.Context, client ctrlruntimeclient.Client, ns string) error {
	roleBinding := &rbacv1.RoleBinding{
		ObjectMeta: v1.ObjectMeta{
			Name:      roleBinding,
			Namespace: ns,
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
			Kind:     "Role",
			Name:     role,
		},
	}

	if err := client.Create(ctx, roleBinding); err != nil {
		if kerrors.IsAlreadyExists(err) {
			return nil
		}

		return err
	}

	return nil
}
