// +build go1.10

package bstring

import (
	"strings"
)

// Buffer Buffer
func Buffer(size ...int) *strings.Builder {
	var b strings.Builder
	if len(size) > 0 {
		b.Grow(size[0])
	}
	return &b
}
