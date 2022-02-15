package fw

import (
	"hash"
	"io"
)

type Checksum struct {
	io.Writer
	hash.Hash
}

func (c *Checksum) Write(p []byte) (int, error) {
	c.Hash.Write(p)
	n, err := c.Writer.Write(p)
	return n, err
}
