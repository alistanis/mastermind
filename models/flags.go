package models

import "flag"

var (
	roleFolder = flag.String("role", defaultRoleDir, "The directory to search for roles. Defaults to the running location of the binary.")
)
