// Copyright 2016 mikan. All rights reserved.

package archive

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type Reader interface {
	ReadAll() []Entry
	FileType() string
}

type Entry struct {
	Name string
	Size int
}

type format struct {
	name      string
	newReader func(io.Reader) (Reader, error)
}

var formats []format

func RegisterFormat(name string, newReader func(io.Reader) (Reader, error)) {
	formats = append(formats, format{name, newReader})
}

func NewReader(in io.Reader) (Reader, error) {
	input, err := ioutil.ReadAll(in) // save stdin
	if err != nil {
		return nil, errors.New("stdin unavailable")
	}

	var errMsgs []string
	for _, f := range formats {
		reader, err := f.newReader(bytes.NewReader(input))
		if err == nil {
			return reader, nil
		}
		errMsgs = append(errMsgs, fmt.Sprintf("%v", err))
	}
	return nil, errors.New(strings.Join(errMsgs, "\n"))
}
