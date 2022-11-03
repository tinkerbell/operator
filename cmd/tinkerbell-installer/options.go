package main

import "time"

type Options struct {
	Verbose         bool
	ChartsDirectory string
}

var options = &Options{}

func (o *Options) CopyInto(other *Options) {
	other.ChartsDirectory = o.ChartsDirectory
	other.Verbose = o.Verbose
}

type DeployOptions struct {
	Options

	Kubeconfig  string
	KubeContext string

	HelmBinary       string
	HelmValues       string
	HelmTimeout      time.Duration
	SkipDependencies bool
	Force            bool
}
