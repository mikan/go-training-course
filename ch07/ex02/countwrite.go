// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

// CountWriter implements bufio.Writer
type CountWriter struct {
	writer io.Writer
	count  int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &CountWriter{
		count:  0,
		writer: w,
	}
	return cw, &cw.count
}

func (b *CountWriter) Write(p []byte) (nn int, err error) {
	nn, err = b.writer.Write(p)
	b.count += int64(nn)
	return nn, err
}

func main() {
	w, c := CountingWriter(ioutil.Discard)
	input := SingleLine("Text")
	n1, _ := w.Write([]byte(input))
	n2, _ := w.Write([]byte(input))
	n3, _ := w.Write([]byte(input))
	fmt.Printf("Wrote 3 times. %d %d %d -> %d\n", n1, n2, n3, *c)
}
