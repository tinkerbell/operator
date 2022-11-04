package main

import (
	"flag"
)

type controllerRunOptions struct {
	kubeconfig              string
	enableLeaderElection    bool
	leaderElectionNamespace string

	workerCount              int
	overwriteRegistry        string
	dockerPullConfigJSONFile string
	namespace                string

	healthProbeAddress string
	metricsAddress     string
}

func newControllerOptions() *controllerRunOptions {
	opts := &controllerRunOptions{}

	if flag.Lookup("kubeconfig") == nil {
		flag.StringVar(&opts.kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	}
	flag.BoolVar(&opts.enableLeaderElection, "enable-leader-election", true, "Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	flag.StringVar(&opts.leaderElectionNamespace, "leader-election-namespace", "", "Leader election namespace. In-cluster discovery will be attempted in such case.")
	flag.IntVar(&opts.workerCount, "worker-count", 1, "Number of workers which process the clusters in parallel.")
	flag.StringVar(&opts.overwriteRegistry, "overwrite-registry", "", "Registry to use for all images")
	flag.StringVar(&opts.dockerPullConfigJSONFile, "docker-pull-config-json-file", "", "The file containing the docker auth config.")
	flag.StringVar(&opts.namespace, "namespace", "kube-system", "The namespace where the tinkerbell controller runs in.")
	flag.StringVar(&opts.healthProbeAddress, "health-probe-address", "127.0.0.1:8085", "The address on which the liveness check on /healthz and readiness check on /readyz will be available")
	flag.StringVar(&opts.metricsAddress, "metrics-address", "127.0.0.1:8080", "The address on which Prometheus metrics will be available under /metrics")

	flag.Parse()

	opts.kubeconfig = flag.Lookup("kubeconfig").Value.(flag.Getter).Get().(string)

	return opts
}
