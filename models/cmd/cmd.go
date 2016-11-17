package cmd

type Command struct {
	Name string
	Args []string
}

type Commands []*Command
