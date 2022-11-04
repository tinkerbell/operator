package rufio

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	clusterRole        = "rufio-manager-role"
	clusterRoleBinding = "rufio-manager-cluster-role-binding"

	role        = "rufio-leader-election-role"
	roleBinding = "rufio-leader-election-role-binding"
)

func CreateClusterRole(ctx context.Context, client ctrlruntimeclient.Client) error {
	clusterRole := &rbacv1.ClusterRole{
		ObjectMeta: v1.ObjectMeta{
			Name: clusterRole,
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{""},
				Resources: []string{"secrets"},
				Verbs:     []string{"get", "list", "watch"},
			},
			{
				APIGroups: []string{"bmc.tinkerbell.org"},
				Resources: []string{"jobs"},
				Verbs:     []string{"create", "delete", "get", "list", "patch", "update", "watch"},
			},
			{
				APIGroups: []string{"bmc.tinkerbell.org"},
				Resources: []string{"jobs/finalizers"},
				Verbs:     []string{"update"},
			},
			{
				APIGroups: []string{"bmc.tinkerbell.org"},
				Resources: []string{"jobs/status"},
				Verbs:     []string{"get", "patch", "update"},
			},
			{
				APIGroups: []string{"bmc.tinkerbell.org"},
				Resources: []string{"machines"},
				Verbs:     []string{"create", "delete", "get", "list", "patch", "update", "watch"},
			},
			{
				APIGroups: []string{"bmc.tinkerbell.org"},
				Resources: []string{"machines/finalizers"},
				Verbs:     []string{"update"},
			},
			{
				APIGroups: []string{"bmc.tinkerbell.org"},
				Resources: []string{"machines/status"},
				Verbs:     []string{"get", "patch", "update"},
			},
			{
				APIGroups: []string{"bmc.tinkerbell.org"},
				Resources: []string{"tasks"},
				Verbs:     []string{"create", "delete", "get", "list", "patch", "update", "watch"},
			},
			{
				APIGroups: []string{"bmc.tinkerbell.org"},
				Resources: []string{"tasks/status"},
				Verbs:     []string{"get", "patch", "update"},
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

func CreateRole(ctx context.Context, client ctrlruntimeclient.Client, ns string) error {
	role := &rbacv1.Role{
		ObjectMeta: v1.ObjectMeta{
			Name:      role,
			Namespace: ns,
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{""},
				Resources: []string{"configmaps"},
				Verbs:     []string{"get", "list", "watch", "create", "update", "patch", "delete"},
			},
			{
				APIGroups: []string{"coordination.k8s.io"},
				Resources: []string{"leases"},
				Verbs:     []string{"get", "list", "watch", "create", "update", "patch", "delete"},
			},
			{
				APIGroups: []string{""},
				Resources: []string{"events"},
				Verbs:     []string{"create", "patch"},
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
