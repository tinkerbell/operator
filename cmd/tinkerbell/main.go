package main

import (
	"fmt"

	"go.uber.org/zap"

	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func main() {
	// Configure logger
	logger, err := zap.NewProduction()
	if err != nil {
		klog.Fatal(err)
	}
	log := logger.Sugar()

	opts := newControllerOptions()
	mgr, err := createManager(opts)
	if err != nil {
		klog.Fatalf("failed to create runtime manager: %v", err)
	}
}

func createManager(opts *controllerRunOptions) (manager.Manager, error) {
	// Manager options
	options := manager.Options{
		LeaderElection:          opts.enableLeaderElection,
		LeaderElectionID:        "tinkerbell-controller",
		LeaderElectionNamespace: opts.namespace,
		HealthProbeBindAddress:  opts.healthProbeAddress,
		MetricsBindAddress:      opts.metricsAddress,
		Port:                    9443,
	}

	mgr, err := manager.New(config.GetConfigOrDie(), options)
	if err != nil {
		return nil, fmt.Errorf("error building ctrlruntime manager: %w", err)
	}

	// Add health endpoints
	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		return nil, fmt.Errorf("failed to add health check: %w", err)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		return nil, fmt.Errorf("failed to add readiness check: %w", err)
	}
	return mgr, nil
}
