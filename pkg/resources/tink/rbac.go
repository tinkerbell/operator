package tink

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	tinkServerClusterRole        = "tink-server-cluster-role"
	tinkServerClusterRoleBinding = "tink-server-cluster-role-binding"

	tinkControllerClusterRole        = "tink-controller-cluster-role"
	tinkControllerClusterRoleBinding = "tink-controller-cluster-role-binding"

	role        = "tink-leader-election-role"
	roleBinding = "tink-leader-election-role-binding"
)

func CreateTinkControllerClusterRole(ctx context.Context, client ctrlruntimeclient.Client) error {
	clusterRole := &rbacv1.ClusterRole{
		ObjectMeta: v1.ObjectMeta{
			Name: tinkControllerClusterRole,
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{"tinkerbell.org"},
				Resources: []string{"hardware", "hardware/status"},
				Verbs:     []string{"get", "list", "patch", "update", "watch"},
			},
			{
				APIGroups: []string{"tinkerbell.org"},
				Resources: []string{"templates", "templates/status"},
				Verbs:     []string{"get", "list", "patch", "update", "watch"},
			},
			{
				APIGroups: []string{"tinkerbell.org"},
				Resources: []string{"workflows", "workflows/status"},
				Verbs:     []string{"delete", "get", "list", "patch", "update", "watch"},
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

func CreateTinkServerClusterRole(ctx context.Context, client ctrlruntimeclient.Client) error {
	clusterRole := &rbacv1.ClusterRole{
		ObjectMeta: v1.ObjectMeta{
			Name: tinkServerClusterRole,
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{"tinkerbell.org"},
				Resources: []string{"hardware", "hardware/status"},
				Verbs:     []string{"get", "list", "watch"},
			},
			{
				APIGroups: []string{"tinkerbell.org"},
				Resources: []string{"templates", "templates/status"},
				Verbs:     []string{"get", "list", "watch"},
			},
			{
				APIGroups: []string{"tinkerbell.org"},
				Resources: []string{"workflows", "workflows/status"},
				Verbs:     []string{"get", "list", "patch", "update", "watch"},
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

func CreateTinkControllerClusterRoleBinding(ctx context.Context, client ctrlruntimeclient.Client, ns string) error {
	clusterRoleBinding := &rbacv1.ClusterRoleBinding{
		ObjectMeta: v1.ObjectMeta{
			Name: tinkControllerClusterRoleBinding,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      rbacv1.ServiceAccountKind,
				Name:      tinkControllerServiceAccountName,
				Namespace: ns,
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: rbacv1.GroupName,
			Kind:     "ClusterRole",
			Name:     tinkControllerClusterRole,
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

func CreateTinkServerClusterRoleBinding(ctx context.Context, client ctrlruntimeclient.Client, ns string) error {
	clusterRoleBinding := &rbacv1.ClusterRoleBinding{
		ObjectMeta: v1.ObjectMeta{
			Name: tinkServerClusterRoleBinding,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      rbacv1.ServiceAccountKind,
				Name:      tinkServerServiceAccountName,
				Namespace: ns,
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: rbacv1.GroupName,
			Kind:     "ClusterRole",
			Name:     tinkServerClusterRole,
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
				Verbs:     []string{"get", "list", "watch", "create", "update", "patch"},
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
				Name:      tinkControllerServiceAccountName,
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
