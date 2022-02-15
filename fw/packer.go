package fw

import (
	"bytes"
	"io"
)

type Packer struct {
	io.Reader
	MaxSize int
	Erase   bool

	size int
	buf  *bytes.Buffer
}

func (pa *Packer) Read(p []byte) (int, error) {
	n, err := pa.Reader.Read(p)
	pa.size += n
	if err == io.EOF {
		if pa.Erase {
			if pa.buf == nil {
				pa.buf = &bytes.Buffer{}
				s := pa.size + 0x800
				r := pa.MaxSize - s
				arr := make([]byte, r)
				for i := range arr {
					arr[i] = 0xff
				}
				pa.buf.Write(arr)
			}
			return pa.buf.Read(p)
		}
	}
	return n, err
}
