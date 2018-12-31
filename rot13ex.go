package main

import (
	"io"
	"os"
	"strings"
	//"fmt"
)

type rot13Reader struct {
	r io.Reader
}

const ad = ('z'-'a')/2 + 1

func (rot *rot13Reader) Read(b []byte) (int, error) {
	//fmt.Print("*")
	for i := 0; i < len(b); i++ {
		switch {
		case 'A' <= b[i] && b[i] <= 'Z':
			b[i] = (b[i]-'A'+ad)%26 + 'A'
		case 'a' <= b[i] && b[i] <= 'z':
			b[i] = (b[i]-'a'+ad)%26 + 'a'
		default:
			b[i] = b[i]
		}
	}
	return len(b), nil
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	v := rot13Reader{s}
	io.Copy(os.Stdout, v.r)
}
