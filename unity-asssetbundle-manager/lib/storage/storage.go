package storage

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
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

// Get dir
func Get(path string) *Dir {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		err := os.MkdirAll(path, 0700)
		return &Dir{Path: path, Error: err}
	}

	return &Dir{Error: err, Path: path}
}

// Dir relative of root
func (d *Dir) Dir(path string) *Dir {
	if d.Error != nil {
		return &Dir{Error: d.Error}
	}

	npath := d.Path + "/" + path

	if err := os.MkdirAll(npath, 0700); err != nil {
		return &Dir{Error: err}
	}

	return &Dir{Path: npath}
}

// Save file to Dir
func (d *Dir) Save(f *multipart.FileHeader) *File {
	if d.Error != nil {
		return &File{Error: d.Error}
	}

	path := fmt.Sprintf("%s/%s", d.Path, f.Filename)
	data := make([]byte, f.Size)
	file, err := f.Open()

	if err != nil {
		return &File{Error: err}
	}

	if _, err := file.Read(data); err != nil {
		return &File{Error: err}
	}

	if err := ioutil.WriteFile(path, data, 0700); err != nil {
		return &File{Error: err}
	}

	return &File{
		Filepath: path,
		Filename: f.Filename,
	}
}
