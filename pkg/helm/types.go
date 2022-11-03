/*
  Copied from https://github.com/kubermatic/kubermatic/tree/main/pkg/install/helm
*/

package helm

import (
	"fmt"
	"os"
	"path/filepath"

	semverlib "github.com/Masterminds/semver/v3"
	"gopkg.in/yaml.v3"
)

type Release struct {
	Name      string             `json:"name"`
	Namespace string             `json:"namespace"`
	Chart     string             `json:"chart"`
	Revision  string             `json:"revision"`
	Version   *semverlib.Version `json:"-"`
	// AppVersion is not a semver, for example Minio has date-based versions.
	AppVersion string        `json:"app_version"`
	Status     ReleaseStatus `json:"status"`
}

func (r *Release) Clone() Release {
	releaseCopy := *r
	releaseCopy.Version = semverlib.MustParse(r.Version.Original())

	return releaseCopy
}

type Chart struct {
	Name       string             `yaml:"name"`
	Version    *semverlib.Version `yaml:"-"`
	VersionRaw string             `yaml:"version"`
	// AppVersion is not a semver, for example Minio has date-based versions.
	AppVersion   string `yaml:"appVersion"`
	Directory    string
	Dependencies []Dependency `yaml:"dependencies,omitempty"`
}

func (c *Chart) Clone() Chart {
	chartCopy := *c
	chartCopy.Version = semverlib.MustParse(c.Version.Original())

	return chartCopy
}

func LoadChart(directory string) (*Chart, error) {
	f, err := os.Open(filepath.Join(directory, "Chart.yaml"))
	if err != nil {
		return nil, fmt.Errorf("failed to open Chart.yaml: %w", err)
	}
	defer f.Close()

	chart := &Chart{}
	if err := yaml.NewDecoder(f).Decode(chart); err != nil {
		return nil, fmt.Errorf("failed to read Chart.yaml: %w", err)
	}

	version, err := semverlib.NewVersion(chart.VersionRaw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse version %q: %w", chart.VersionRaw, err)
	}

	chart.Version = version
	chart.Directory = directory

	return chart, nil
}

type Dependency struct {
	Name         string        `yaml:"name"`
	Version      string        `yaml:"version,omitempty"`
	Repository   string        `yaml:"repository"`
	Condition    string        `yaml:"condition,omitempty"`
	Tags         []string      `yaml:"tags,omitempty"`
	Enabled      bool          `yaml:"enabled,omitempty"`
	ImportValues []interface{} `json:"import-values,omitempty"`
	Alias        string        `yaml:"alias,omitempty"`
}
