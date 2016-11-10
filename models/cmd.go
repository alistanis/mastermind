package models

type Command struct {
	Name string
	Args []string
}

type Commands []*Command
