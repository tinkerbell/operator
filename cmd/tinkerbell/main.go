package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	kubetinkctrl "github.com/moadqassem/kubetink/pkg/controller"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func main() {
	// Configure logger
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("failed to create logger: %v", err)
		os.Exit(1)
	}

	log := logger.Sugar()

	opts := newControllerOptions()
	mgr, err := createManager(opts)
	if err != nil {
		log.Fatalf("failed to create runtime manager: %v", err)
	}

	if err := kubetinkctrl.Add(mgr, log, opts.namespace, opts.workerCount); err != nil {
		log.Fatalf("failed to add kubetink controller to manager: %v", err)
	}

	log.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		log.Fatalf("Failed to start kubetink controller: %v", zap.Error(err))
	}
}

func createManager(opts *controllerRunOptions) (manager.Manager, error) {
	// Manager options
	options := manager.Options{
		LeaderElection:          opts.enableLeaderElection,
		LeaderElectionID:        "tinkerbell-controller",
		LeaderElectionNamespace: opts.leaderElectionNamespace,
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
