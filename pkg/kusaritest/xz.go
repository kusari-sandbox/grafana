package kusaritest

import (
	"io"

	"github.com/ulikunitz/xz"
)

func Decompress(r io.Reader) (io.Reader, error) {
	return xz.NewReader(r)
}
