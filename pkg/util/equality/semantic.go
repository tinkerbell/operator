/*
  Copied from https://github.com/kubermatic/kubermatic/tree/main/pkg/apis/equality
*/

package equality

import (
	"time"

	semverlib "github.com/Masterminds/semver/v3"

	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/conversion"
)

// Semantic can do semantic deep equality checks for objects.
// Example: equality.Semantic.DeepEqual(aPod, aPodWithNonNilButEmptyMaps) == true.
var Semantic = conversion.EqualitiesOrDie(
	func(a, b resource.Quantity) bool {
		return a.Cmp(b) == 0
	},
	func(a, b *semverlib.Version) bool {
		if a == nil && b == nil {
			return true
		}

		if a != nil && b != nil {
			return a.Equal(b)
		}

		return false
	},
	func(a, b time.Time) bool {
		return a.UTC() == b.UTC()
	},
)
