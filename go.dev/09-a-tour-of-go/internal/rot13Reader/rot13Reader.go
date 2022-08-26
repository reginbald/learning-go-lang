package rot13Reader

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(bytes []byte) (int, error) {
	n, err := reader.r.Read(bytes)
	for i := range bytes {
		x := bytes[i]
		switch {
		case 65 <= x && x <= 90:
			bytes[i] = 65 + (bytes[i]-65+13)%26
		case 97 <= x && x <= 122:
			bytes[i] = 97 + (bytes[i]-97+13)%26
		}
	}
	return n, err
}

func Run() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
