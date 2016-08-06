// Copyright 2016 mikan. All rights reserved.

package tar

import (
	"archive/tar"

	"bytes"
	"io"

	"github.com/mikan/gopl/ch10/ex02/archive"
)

func init() {
	archive.RegisterFormat("tar", newTarReader)
}

type tarReader tar.Reader

var entries []archive.Entry

func (r tarReader) ReadAll() []archive.Entry {
	return entries
}

func (r tarReader) FileType() string {
	return "tar"
}

func newTarReader(in io.Reader) (archive.Reader, error) {
	reader := tar.NewReader(in)
	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		buf := new(bytes.Buffer)
		if _, err = io.Copy(buf, reader); err != nil {
			return nil, err
		}
		entries = append(entries, archive.Entry{header.Name, int(header.Size)})
	}
	return tarReader(*reader), nil
}
