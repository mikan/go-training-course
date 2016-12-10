// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mikan/go-training-course/ch10/ex02/archive"
	_ "github.com/mikan/go-training-course/ch10/ex02/archive/tar"
	_ "github.com/mikan/go-training-course/ch10/ex02/archive/zip"
)

func main() {
	if err := printArchiveContents(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
}

func printArchiveContents(in io.Reader, out io.Writer) error {
	reader, err := archive.NewReader(in)
	if err != nil {
		return err
	}
	fmt.Println("Type: " + reader.FileType())
	for _, entry := range reader.ReadAll() {
		fmt.Printf("%s (%d kB)\n", entry.Name, entry.Size/1024)
	}
	return nil
}
