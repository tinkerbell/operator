package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/moadqassem/kubetink/pkg/resources"
)

type Options struct {
	Verbose         bool
	ChartsDirectory string
}

var options = &Options{}

func (o *Options) CopyInto(other *Options) {
	other.ChartsDirectory = o.ChartsDirectory
	other.Verbose = o.Verbose
}

func main() {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "18:01:01",
	}

	tinkCommand := &cobra.Command{
		Use:     "tinkerbell-installer",
		Short:   "Installs and updates the tinkerbell stack in a kubernetes cluster.",
		Version: resources.TinkerbellVersion,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if env := os.Getenv("TINKERBELL_CHARTS_DIRECTORY"); env != "" {
				options.ChartsDirectory = env
			}

			if options.ChartsDirectory == "" {
				options.ChartsDirectory = "charts"
			}

			if options.Verbose {
				logger.SetLevel(logrus.DebugLevel)
			}

			logger.SetFormatter(&logrus.JSONFormatter{})
		},
	}

	tinkCommand.SetFlagErrorFunc(func(c *cobra.Command, err error) error {
		if err := c.Usage(); err != nil {
			return err
		}

		// ensure we exit with code 1 later on
		return err
	})

	tinkCommand.PersistentFlags().BoolVarP(&options.Verbose, "verbose", "v", options.Verbose, "enable more verbose output")
	tinkCommand.PersistentFlags().StringVar(&options.ChartsDirectory, "charts-directory", "", "filesystem path to the Tinkerbell Helm charts (defaults to charts/)")

	if err := tinkCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
