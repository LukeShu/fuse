package mountinfo

import (
	"testing"
)

func FuzzParse(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		_, _ = parse(data)
	})
}
