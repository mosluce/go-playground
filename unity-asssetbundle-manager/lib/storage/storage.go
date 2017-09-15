package storage

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Dir hold path of dir
type Dir struct {
	Error error
	Path  string
}

// File hold file infomation
type File struct {
	Error    error
	Filepath string
	Filename string
}

// Root dir
func Root(path string) *Dir {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		err := os.MkdirAll(path, 0700)
		return &Dir{Path: path, Error: err}
	}

	return &Dir{Error: err}
}

// Dir relative of root
func (d *Dir) Dir(path string) *Dir {
	npath := d.Path + path

	if err := os.MkdirAll(npath, 0700); err != nil {
		return &Dir{Error: err}
	}

	return &Dir{Path: path}
}

// Write file to Dir
func (d *Dir) Write(filename string, data []byte) *File {
	npath := fmt.Sprintf("%s/%s", d.Path, filename)
	err := ioutil.WriteFile(npath, data, 0700)
	return &File{
		Error:    err,
		Filepath: npath,
		Filename: filename,
	}
}
