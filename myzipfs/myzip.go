package myzipfs

import (
	"archive/zip"
	"io"
	"io/fs"
)

type ZipFs struct {
	*zip.Reader
}

func NewReader(r io.ReaderAt, size int64) (fs.FS, error) {
	var err error
	zfs := ZipFs{}
	zfs.Reader, err = zip.NewReader(r, size)
	if err != nil {
		return nil, err
	}
	return &zfs, nil
}

func (zfs *ZipFs) ReadDir(name string) ([]fs.DirEntry, error) {
	ds := []fs.DirEntry{}
	for _, d := range zfs.File {
		ds = append(ds, &DirEntry{File: d})
	}
	return ds, nil
}

type DirEntry struct {
	*zip.File
}

func (de DirEntry) Info() (fs.FileInfo, error) {
	return de.FileInfo(), nil
}
func (de DirEntry) IsDir() bool {
	return de.Mode().IsDir()
}

func (de DirEntry) Name() string {
	return de.File.Name
}

func (de DirEntry) Type() fs.FileMode {
	return de.Mode()
}
