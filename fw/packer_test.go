package fw

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestPacker_Read(t *testing.T) {
	tc := []struct {
		name        string
		erase       bool
		expectedLen int
		expectedErr error
		expectedMD5 string
	}{
		{
			name:        "file is copied without additional bytes",
			erase:       false,
			expectedLen: 12,
			expectedMD5: "84825e912a1c0b18813f31a7aa366c57",
		},
		{
			name:        "file copied, erase eeprom file always has 489472 bytes.",
			erase:       true,
			expectedLen: 0x77800,
			expectedMD5: "8bc1d97d5d750748181c902979c8b795",
		},
	}
	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			input := strings.NewReader("input-string")
			p := Packer{
				Reader:  input,
				MaxSize: 0x78000,
				Erase:   c.erase,
			}
			buf := bytes.Buffer{}
			_, err := io.Copy(&buf, &p)
			if !errors.Is(err, c.expectedErr) {
				t.Fatal(err)
			}
			got := buf.Len()
			if got != c.expectedLen {
				t.Errorf("Result buffer has wrong size. Got %x, Wanted %x", got, c.expectedLen)

			}
			gotMD5 := fmt.Sprintf("%x", md5.Sum(buf.Bytes()))
			if c.expectedMD5 != gotMD5 {
				t.Errorf("File MD5 checksum is wrong. Got %s, Wanted, %s", gotMD5, c.expectedMD5)
			}
		})
	}
}
