package resources

import semverlib "github.com/Masterminds/semver/v3"

const (
	TinkerbellVersion = "v0.0.1"
)

var (
	MinHelmVersion = semverlib.MustParse("v3.0.0")
)
