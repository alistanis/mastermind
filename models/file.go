package models

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Files []*File

type File struct {
	*os.File
	Info     os.FileInfo
	MD5      []byte
	Path     string
	BaseName string
}

func NewFile(name string) (*File, error) {

	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	file := &File{}
	file.File = f
	file.Info = fi

	m := md5.New()
	// potentially super expensive and bad, because large files and reasons. can we trust fi.ModTime()?
	// if we do, a new "save" would change the modtime without changing the data, thus making service restarts
	// and redeployments more common even if the file itself doesn't change
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	m.Write(data)
	file.MD5 = m.Sum(nil)
	file.BaseName = filepath.Base(name)
	file.Path = name
	return file, nil
}

func (f *File) MD5String() string {
	return fmt.Sprintf("%x", f.MD5)
}

type FileManifest struct {
	Files
}

func NewFileManifest() *FileManifest {
	return &FileManifest{Files: make(Files, 0)}
}
