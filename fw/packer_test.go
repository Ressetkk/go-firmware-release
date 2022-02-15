package fw

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func TestPacker_Read(t *testing.T) {
	tc := []struct {
		name        string
		erase       bool
		expected    int
		expectedErr error
	}{
		{
			name:     "file is copied without additional bytes",
			erase:    false,
			expected: 12,
		},
		{
			name:     "file copied, erase eeprom file always has 489472 bytes.",
			erase:    true,
			expected: 0x77800,
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
			buf := strings.Builder{}
			_, err := io.Copy(&buf, &p)
			if !errors.Is(err, c.expectedErr) {
				t.Fatal(err)
			}
			got := buf.Len()
			if got != c.expected {
				t.Errorf("Result buffer has wrong size. Got %x, Wanted %x", got, c.expected)

			}
		})
	}
}
