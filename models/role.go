package models

type Role struct {
	// The name of this Role
	Name string
	// A collection of files that will be transferred to the agent machine
	*FileManifest
	// The packages that this role will have installed
	Packages []*Package
	// The commands this role will run before starting services and after packages are installed
	Commands []*Command
	// The services this role will run after packages are installed and commands are run
	Services []*Service
}
