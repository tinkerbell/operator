/*
  Copied from https://github.com/kubermatic/kubermatic/tree/main/pkg/util/yamled
*/

package yamled

import (
	"fmt"
	"strings"
)

type Step interface{}

type Path []Step

func (p Path) Append(s Step) Path {
	return append(p, s)
}

// Parent returns the path except for the last element.
func (p Path) Parent() Path {
	if len(p) < 1 {
		return nil
	}

	return p[0 : len(p)-1]
}

func (p Path) End() Step {
	if len(p) == 0 {
		return nil
	}

	return p[len(p)-1]
}

func (p Path) String() string {
	parts := []string{}

	for _, p := range p {
		if s, ok := p.(string); ok {
			parts = append(parts, s)
			continue
		}

		if i, ok := p.(int); ok {
			parts = append(parts, fmt.Sprintf("[%d]", i))
			continue
		}

		parts = append(parts, fmt.Sprintf("%v", p))
	}

	return strings.Join(parts, ".")
}
