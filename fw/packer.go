package fw

import (
	"io"
)

type Packer struct {
	io.Reader
	MaxSize int
	Erase   bool

	size, missing int
}

func (pa *Packer) Read(p []byte) (int, error) {
	n, err := pa.Reader.Read(p)
	if err == io.EOF && pa.Erase {
		if pa.missing == 0 {
			pa.missing = pa.MaxSize - (pa.size + 0x800)
		}
		for n = range p {
			p[n] = 0xff
			if pa.missing-n == 0 {
				return n, io.EOF
			}
		}
		pa.missing -= n
		return n, nil
	}
	pa.size += n
	return n, err
}
