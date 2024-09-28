// Package e2etests contains the infrastructure for end to
// end testing of Play-Go.
package e2etests

import (
	"os"
	"testing"

	"github.com/nwchandler/playgo"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m,
		map[string]func() int{
			"playgo": playgo.Main,
		}),
	)
}

func TestEndToEnd(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata/scripts",
	})
}
