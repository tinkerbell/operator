package main

//
//import (
//	"errors"
//	"fmt"
//	"os"
//	"time"
//
//	"github.com/sirupsen/logrus"
//	"github.com/spf13/cobra"
//
//	"github.com/moadqassem/kubetink/pkg/helm"
//	"github.com/moadqassem/kubetink/pkg/resources"
//)
//
//func DeployCommand(logger *logrus.Logger) *cobra.Command {
//	opt := DeployOptions{
//		HelmTimeout: 5 * time.Minute,
//		HelmBinary:  "helm",
//	}
//
//	cmd := &cobra.Command{
//		Use:          "deploy",
//		Short:        "Install tinkerbell stack to the installer's built-in version",
//		Long:         "Installs tinkerbell stack to the installer's built-in version",
//		RunE:         DeployFunc(logger, &opt),
//		SilenceUsage: true,
//		PreRun: func(cmd *cobra.Command, args []string) {
//			options.CopyInto(&opt.Options)
//
//			if opt.Kubeconfig == "" {
//				opt.Kubeconfig = os.Getenv("KUBECONFIG")
//			}
//			if opt.KubeContext == "" {
//				opt.KubeContext = os.Getenv("KUBE_CONTEXT")
//			}
//			if opt.HelmValues == "" {
//				opt.HelmValues = os.Getenv("HELM_VALUES")
//			}
//			if opt.HelmBinary == "" {
//				opt.HelmBinary = os.Getenv("HELM_BINARY")
//			}
//		},
//	}
//
//	cmd.PersistentFlags().StringVar(&opt.Kubeconfig, "kubeconfig", "", "full path to where a kubeconfig with cluster-admin permissions for the target cluster")
//	cmd.PersistentFlags().StringVar(&opt.KubeContext, "kube-context", "", "context to use from the given kubeconfig")
//
//	cmd.PersistentFlags().StringVar(&opt.HelmValues, "helm-values", "", "full path to the Helm values.yaml used for customizing all charts")
//	cmd.PersistentFlags().DurationVar(&opt.HelmTimeout, "helm-timeout", opt.HelmTimeout, "time to wait for Helm operations to finish")
//	cmd.PersistentFlags().StringVar(&opt.HelmBinary, "helm-binary", opt.HelmBinary, "full path to the Helm 3 binary to use")
//	cmd.PersistentFlags().BoolVar(&opt.SkipDependencies, "skip-dependencies", false, "skip pulling Helm chart dependencies (requires chart dependencies to be already downloaded)")
//	cmd.PersistentFlags().BoolVar(&opt.Force, "force", false, "perform Helm upgrades even when the release is up-to-date")
//
//	return cmd
//}
//
//func DeployFunc(logger *logrus.Logger, opt *DeployOptions) cobraFuncE {
//	return handleErrors(logger, func(cmd *cobra.Command, args []string) error {
//		helmClient, err := helm.NewCLI(opt.HelmBinary, opt.Kubeconfig, opt.KubeContext, opt.HelmTimeout, logger)
//		if err != nil {
//			return fmt.Errorf("failed to create Helm client: %w", err)
//		}
//
//		helmVersion, err := helmClient.Version()
//		if err != nil {
//			return fmt.Errorf("failed to check Helm version: %w", err)
//		}
//
//		if helmVersion.LessThan(resources.MinHelmVersion) {
//			return fmt.Errorf(
//				"the installer requires Helm >= %s, but detected %q as %s (use --helm-binary or $HELM_BINARY to override)",
//				resources.MinHelmVersion,
//				opt.HelmBinary,
//				helmVersion,
//			)
//		}
//
//		logger.Info("Initializing tinkerbell installer…")
//
//		if len(opt.Kubeconfig) == 0 {
//			return errors.New("no kubeconfig (--kubeconfig or $KUBECONFIG) given")
//		}
//
//		helmValues, err := loadHelmValues(opt.HelmValues)
//		if err != nil {
//			return fmt.Errorf("failed to load Helm values: %w", err)
//		}
//
//		helmChart, err := helm.LoadChart(opt.ChartsDirectory)
//		if err != nil {
//			return fmt.Errorf("failed to load Helm chart: %w", err)
//		}
//
//
//		return nil
//	})
//}
