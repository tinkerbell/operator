/*
  Copied from https://github.com/kubermatic/kubermatic/tree/main/pkg/install/helm
*/

package helm

type ReleaseStatus string

const (
	ReleaseCheckFailed ReleaseStatus = ""

	// these constants mirror the Helm status from
	// `helm status --help`.

	ReleaseStatusUnknown         ReleaseStatus = "unknown"
	ReleaseStatusDeployed        ReleaseStatus = "deployed"
	ReleaseStatusDeleted         ReleaseStatus = "uninstalled"
	ReleaseStatusSuperseded      ReleaseStatus = "superseded"
	ReleaseStatusFailed          ReleaseStatus = "failed"
	ReleaseStatusDeleting        ReleaseStatus = "uninstalling"
	ReleaseStatusPendingInstall  ReleaseStatus = "pending-install"
	ReleaseStatusPendingUpgrade  ReleaseStatus = "pending-upgrade"
	ReleaseStatusPendingRollback ReleaseStatus = "pending-rollback"
)
