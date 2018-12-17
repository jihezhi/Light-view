package main

import (
	//"golang.org/x/tour/reader"
	"strings"
	//"io"
	"fmt"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r *MyReader) Read(b []byte) (int, error) {

	n := len(b)
	return n, nil
}

func main() {
	r := strings.NewReader("A")
	b := make([]byte, 1)
	//reader.Validate(MyReader{})
	n, err := r.Read(b)
	for err == nil {
		fmt.Print(string(b[:n]))
		//break
	}
}
