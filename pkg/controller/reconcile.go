package controller

import (
	"context"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"

	"github.com/moadqassem/kubetink/pkg/resources"
	"github.com/moadqassem/kubetink/pkg/resources/boots"
	"github.com/moadqassem/kubetink/pkg/resources/hegel"
	"github.com/moadqassem/kubetink/pkg/resources/rufio"
	"github.com/moadqassem/kubetink/pkg/resources/tink"
)

func (r *Reconciler) ensureTinkerbellServiceAccounts(ctx context.Context) error {
	if err := boots.CreateServiceAccount(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create boots service account: %v", err)
	}

	if err := hegel.CreateServiceAccount(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create hegel service account: %v", err)
	}

	if err := rufio.CreateServiceAccount(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create rufio service account: %v", err)
	}

	if err := tink.CreateTinkControllerServiceAccount(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create tink controller service account: %v", err)
	}

	if err := tink.CreateTinkServerServiceAccount(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create tink server service account: %v", err)
	}

	return nil
}

func (r *Reconciler) ensureTinkerbellClusterRole(ctx context.Context) error {
	if err := boots.CreateClusterRole(ctx, r.Client); err != nil {
		return fmt.Errorf("failed to create boots cluster role: %v", err)
	}

	if err := rufio.CreateClusterRole(ctx, r.Client); err != nil {
		return fmt.Errorf("failed to create rufio cluster role: %v", err)
	}

	if err := tink.CreateTinkControllerClusterRole(ctx, r.Client); err != nil {
		return fmt.Errorf("failed to create tink controller cluster role: %v", err)
	}

	if err := tink.CreateTinkServerClusterRole(ctx, r.Client); err != nil {
		return fmt.Errorf("failed to create tink server cluster role: %v", err)
	}

	return nil
}

func (r *Reconciler) ensureTinkerbellClusterRoleBinding(ctx context.Context) error {
	if err := boots.CreateClusterRoleBinding(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create boots cluster role binding: %v", err)
	}

	if err := rufio.CreateClusterRoleBinding(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create rufio cluster role binding: %v", err)
	}

	if err := tink.CreateTinkControllerClusterRoleBinding(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create tink controller cluster role binding: %v", err)
	}

	if err := tink.CreateTinkServerClusterRoleBinding(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create tink server cluster role binding: %v", err)
	}

	return nil
}

func (r *Reconciler) ensureTinkerbellRole(ctx context.Context) error {
	if err := hegel.CreateRole(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create hegel role: %v", err)
	}

	if err := rufio.CreateRole(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create rufio role: %v", err)
	}

	if err := tink.CreateRole(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create tink leader election role: %v", err)
	}

	return nil
}

func (r *Reconciler) ensureTinkerbellRoleBinding(ctx context.Context) error {
	if err := hegel.CreateRoleBinding(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create hegel role binding: %v", err)
	}

	if err := rufio.CreateRoleBinding(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create rufio role binding: %v", err)
	}

	if err := tink.CreateRoleBinding(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create tink leader election role binding: %v", err)
	}

	return nil
}

func (r *Reconciler) ensureTinkerbellServices(ctx context.Context) error {
	if err := boots.CreateService(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create boots service: %v", err)
	}

	if err := hegel.CreateService(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create hegel service: %v", err)
	}

	if err := tink.CreateService(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create tink service: %v", err)
	}

	return nil
}

func (r *Reconciler) ensureTinkerbellDeployments(ctx context.Context) error {
	if err := boots.CreateDeployment(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create boots deployment: %v", err)
	}

	if err := hegel.CreateDeployment(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create hegel deployment: %v", err)
	}

	if err := rufio.CreateDeployment(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create rufio deployment: %v", err)
	}

	if err := tink.CreateTinkControllerDeployment(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create tink controller deployment: %v", err)
	}

	if err := tink.CreateTinkServerDeployment(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create tink server deployment: %v", err)
	}

	if err := tink.CreateTinkStackDeployment(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create tink stack deployment: %v", err)
	}

	return nil
}

func (r *Reconciler) ensureTinkerbellConfigMaps(ctx context.Context) error {
	if err := tink.CreateNginxConfigMap(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create stack nginx configmap: %v", err)
	}

	return nil
}

func (r *Reconciler) ensureTinkerbellNamespace(ctx context.Context) error {

	ns := &corev1.Namespace{}
	// TODO: We could read the namespace from the command line as well.
	err := r.Get(ctx, types.NamespacedName{Name: resources.TinkerbellNamespace}, ns)
	if err == nil {
		return nil // found it
	}
	if !apierrors.IsNotFound(err) {
		return err
	}

	ns = &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: resources.TinkerbellNamespace,
		},
	}

	//TODO: add IgnoreAlreadyExists utility in the error check
	if err := r.Create(ctx, ns); err != nil {
		return fmt.Errorf("failed to create Namespace %s: %w", resources.TinkerbellNamespace, err)
	}

	// make sure that the namespace is created and presented in the cache.
	err = wait.PollImmediate(1*time.Second, 30*time.Second, func() (bool, error) {
		ns := &corev1.Namespace{}
		err := r.Get(ctx, types.NamespacedName{Name: resources.TinkerbellNamespace}, ns)
		if err == nil {
			return true, nil
		}
		if apierrors.IsNotFound(err) {
			return false, nil
		}

		return false, err
	})
	if err != nil {
		return fmt.Errorf("failed to wait for cluster namespace to appear in cache: %w", err)
	}

	return nil
}
