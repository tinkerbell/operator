/*
  Copied from https://github.com/kubermatic/kubermatic/tree/main/pkg/install/helm
*/

package helm

import (
	semverlib "github.com/Masterminds/semver/v3"

	"github.com/moadqassem/kubetink/pkg/util/yamled"
)

// Client describes the operations that the Helm client is providing to
// the installer.
type Client interface {
	Version() (*semverlib.Version, error)
	BuildChartDependencies(chartDirectory string, flags []string) error
	InstallChart(namespace string, releaseName string, chartDirectory string, valuesFile string, values map[string]string, flags []string) error
	GetRelease(namespace string, name string) (*Release, error)
	ListReleases(namespace string) ([]Release, error)
	UninstallRelease(namespace string, name string) error
	RenderChart(namespace string, releaseName string, chartDirectory string, valuesFile string, values map[string]string) ([]byte, error)
	GetValues(namespace string, releaseName string) (*yamled.Document, error)
}
