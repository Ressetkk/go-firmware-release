package fw

import (
	"crypto/md5"
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestChecksum_Write(t *testing.T) {
	tc := []struct {
		name        string
		input       string
		expected    []byte
		expectedErr error
	}{
		{
			name:     "md5 of simple-checksum",
			input:    "simple-checksum",
			expected: []byte{0xfb, 0xca, 0xaa, 0x04, 0xc1, 0x82, 0x78, 0x60, 0xe4, 0x44, 0xd2, 0x5a, 0x33, 0x16, 0xa0, 0x21},
		},
	}
	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			test := strings.NewReader(c.input)
			buf := strings.Builder{}
			m := Checksum{
				Writer: &buf,
				Hash:   md5.New(),
			}
			_, err := io.Copy(&m, test)
			if !errors.Is(err, c.expectedErr) {
				t.Fatal(err)
			}
			got := m.Sum(nil)
			if !reflect.DeepEqual(got, c.expected) {
				t.Errorf("Checksum does not match. Got %x, Wanted %x", got, c.expected)
			}
		})
	}
}
