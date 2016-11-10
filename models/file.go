package models

import (
	"crypto/md5"
	"io/ioutil"
	"os"
)

type Files []*File

type File struct {
	*os.File
	Info os.FileInfo
	MD5  []byte
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

	file.MD5 = m.Sum(data)
	return file, nil
}

type FileManifest struct {
	Files
}
