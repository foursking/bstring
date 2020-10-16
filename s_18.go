// +build !go1.10

package ffstring

import "bytes"

// Buffer Buffer
func Buffer(size ...int) *bytes.Buffer {
	b := bytes.NewBufferString("")
	return &b
}
