package models

type Service struct {
	// runner is the system runner (init.d, systemd, etc)
	Runner     string
	Executable string
	Args       []string
}
