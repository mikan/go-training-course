// Copyright 2016 mikan. All rights reserved.

package tar

import (
	"archive/zip"

	"encoding/binary"
	"io"
	"io/ioutil"

	"github.com/mikan/gopl/ch10/ex02/archive"
)

func init() {
	archive.RegisterFormat("zip", newZipReader)
}

type zipReader zip.Reader

func (r zipReader) ReadAll() []archive.Entry {
	var entries []archive.Entry
	for _, file := range r.File {
		entries = append(entries, archive.Entry{file.Name, int(file.UncompressedSize)})
	}
	return entries
}

func (r zipReader) FileType() string {
	return "zip"
}

type readerBuf []byte

func (r readerBuf) ReadAt(b []byte, off int64) (int, error) {
	copy(b, r[int(off):int(off)+len(b)])
	return len(b), nil
}

func newZipReader(in io.Reader) (archive.Reader, error) {
	var data readerBuf
	var err error
	data, err = ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}
	reader, err := zip.NewReader(data, int64(binary.Size(data)))
	if err != nil {
		return nil, err
	}
	return zipReader(*reader), nil
}
