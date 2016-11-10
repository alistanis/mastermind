package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoadRoles(t *testing.T) {
	Convey("We can load some roles", t, func() {
		*roleFolder = Gopath + "/src/github.com/alistanis/mastermind/example_data/roles"
		err := LoadRoles()
		So(err, ShouldBeNil)
	})
}
