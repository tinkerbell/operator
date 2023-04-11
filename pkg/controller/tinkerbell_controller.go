package controller

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/tinkerbell/operator/pkg/util"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	ctrlruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const (
	ControllerName = "TinkerbellController"
)

func Add(mgr manager.Manager, log *zap.SugaredLogger, clusterDNS, namespace string, workerCount int) error {
	reconciler := &Reconciler{
		Client:     mgr.GetClient(),
		log:        log,
		namespace:  namespace,
		clusterDNS: clusterDNS,
	}

	c, err := controller.New(ControllerName, mgr, controller.Options{Reconciler: reconciler, MaxConcurrentReconciles: workerCount})
	if err != nil {
		return err
	}

	typesToWatch := []client.Object{
		&corev1.Service{},
		&corev1.ServiceAccount{},
		&appsv1.Deployment{},
		&rbacv1.Role{},
		&rbacv1.RoleBinding{},
		&rbacv1.ClusterRole{},
		&rbacv1.ClusterRoleBinding{},
	}

	for _, t := range typesToWatch {
		if err := c.Watch(&source.Kind{Type: t}, &handler.EnqueueRequestForObject{}, util.ByNamespace(namespace)); err != nil {
			return fmt.Errorf("failed to create watch for %T: %w", t, err)
		}
	}

	return nil
}

type Reconciler struct {
	client.Client
	log *zap.SugaredLogger

	namespace  string
	clusterDNS string
}

func (r *Reconciler) Reconcile(ctx context.Context, req ctrlruntime.Request) (reconcile.Result, error) {
	r.log.Info("Reconciling tinkerbell resources..")

	if err := r.reconcile(ctx); err != nil {
		r.log.Errorf("failed to reconcile %q due to: %v", req.Name, err)
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

func (r *Reconciler) reconcile(ctx context.Context) error {
	if err := r.ensureTinkerbellServiceAccounts(ctx); err != nil {
		return fmt.Errorf("failed to ensure tinkerbell service accounts: %v", err)
	}

	if err := r.ensureTinkerbellClusterRole(ctx); err != nil {
		return fmt.Errorf("failed to ensure tinkerbell cluster role: %v", err)
	}

	if err := r.ensureTinkerbellClusterRoleBinding(ctx); err != nil {
		return fmt.Errorf("failed to ensure tinkerbell cluster role bindings: %v", err)
	}

	if err := r.ensureTinkerbellRole(ctx); err != nil {
		return fmt.Errorf("failed to ensure tinkerbell role: %v", err)
	}

	if err := r.ensureTinkerbellRoleBinding(ctx); err != nil {
		return fmt.Errorf("failed to ensure tinkerbell role binding: %v", err)
	}

	if err := r.ensureTinkerbellServices(ctx); err != nil {
		return fmt.Errorf("failed to ensure tinkerbell services: %v", err)
	}

	if err := r.ensureTinkerbellConfigMaps(ctx); err != nil {
		return fmt.Errorf("failed to ensure tinkerbell stack configmaps: %v", err)
	}

	if err := r.ensureTinkerbellDeployments(ctx); err != nil {
		return fmt.Errorf("failed to ensure tinkerbell deployments: %v", err)
	}
	return nil
}
