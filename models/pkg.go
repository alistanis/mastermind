package models

type Package struct {
	Name       string
	Repository string
	OS         string
	Version    string
	Arch       string
	Manager    string
}

type Packages []*Package
