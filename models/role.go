package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/alistanis/mastermind/models/cmd"

	"os"
)

// Role
// Having a "FilesRoot" is incredibly naive, but it works for this system's initial purpose. If files transferred need to live
// somewhere else, this can be accomplished by providing command(s) in order to move those files from the root to another specified
// path on the target machine.
// yes, this sucks, but it's a stop gap measure that will work. plus, hopefully we're handling config files with dynamic tools like consul
type Role struct {
	// The name of this Role
	Name string
	// The folder where these files will be stored
	FilesRoot string
	// A collection of files that will be transferred to the agent machine
	*FileManifest
	// The packages that this role will have installed
	Packages
	// The commands this role will run before starting services and after packages are installed
	// These may execute files that have already been transferred.
	cmd.Commands
	// The services this role will run after packages are installed and commands are run
	Services
}

var (
	Roles = make(map[string]*Role)
)

func NewRole() *Role {

	return &Role{Packages: make(Packages, 0),
		Commands:     make(cmd.Commands, 0),
		Services:     make(Services, 0),
		FileManifest: NewFileManifest()}
}

type roleDirStructure struct {
	Name       string
	Path       string
	FilesPath  string
	ConfigPath string
}

// LoadRoles loads roles from a roles folder. The structure is assumed, so this is not a recursive function (with exception to the files directory)
func LoadRoles() error {

	roleDirs, err := ioutil.ReadDir(*roleFolder)
	if err != nil {
		return err
	}
	roleDirStructs := []*roleDirStructure{}
	for _, rd := range roleDirs {
		r := &roleDirStructure{}
		r.Name = filepath.Base(rd.Name())

		r.Path = *roleFolder + string(filepath.Separator) + rd.Name()

		contents, err := ioutil.ReadDir(r.Path)
		if err != nil {
			return err
		}

		// i kinda hate this but hey, it works
		for _, d := range contents {
			if d.IsDir() && d.Name() == "files" {
				r.FilesPath = r.Path + string(filepath.Separator) + d.Name()
			}
			// unforunately binds every role's config name to config.json - we can deal with this at a later date, but like everything else, fuck it for now
			if d.Name() == "config.json" {
				r.ConfigPath = r.Path + string(filepath.Separator) + d.Name()
			}
		}

		if r.ConfigPath == "" {
			return fmt.Errorf("Could not find config file for role %s", r.Name)
		}

		roleDirStructs = append(roleDirStructs, r)
	}

	for _, rd := range roleDirStructs {
		r := NewRole()
		data, err := ioutil.ReadFile(rd.ConfigPath)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, &r)
		if err != nil {
			return err
		}

		if rd.FilesPath != "" {
			err = filepath.Walk(rd.FilesPath,
				func(path string, f os.FileInfo, err error) error {
					if err != nil {
						return err
					}

					if path == *roleFolder {
						return nil
					}

					if f.IsDir() {
						return nil
					}
					newFile, err := NewFile(path)
					if err != nil {
						return err
					}
					r.FileManifest.Files = append(r.FileManifest.Files, newFile)
					return nil
				})

			if err != nil {
				return err
			}

		}

		// TODO - other things with files manifest
		for _, f := range r.FileManifest.Files {
			fmt.Println(f)
		}

		// not thread safe
		Roles[r.Name] = r

	}
	//spew.Dump(Roles)
	return nil
}

func visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if path == *roleFolder {
		return nil
	}
	fmt.Println(path)
	fmt.Println(f.Name())
	return nil
}
