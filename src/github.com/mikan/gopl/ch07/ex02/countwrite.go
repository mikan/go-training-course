// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"github.com/mikan/libmikan/input"
	"io"
	"io/ioutil"
)

// See bufio.Writer
type CountWriter struct {
	err   error
	n     int
	wr    io.Writer
	count int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &CountWriter{
		count: 0,
		wr:    w,
	}
	return cw, &cw.count
}

func (b *CountWriter) Write(p []byte) (nn int, err error) {
	nn, err = b.wr.Write(p)
	b.count += int64(nn)
	return nn, err
}

func main() {
	w, c := CountingWriter(ioutil.Discard)
	input := input.SingleLine("Text")
	n1, _ := w.Write([]byte(input))
	n2, _ := w.Write([]byte(input))
	n3, _ := w.Write([]byte(input))
	fmt.Printf("Wrote 3 times. %d %d %d -> %d\n", n1, n2, n3, *c)
}
