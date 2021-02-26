package control_test

import (
	"testing"

	"github.com/spinozanose/go-hello-world/control"
)


func TestDefer(t *testing.T) {
	control.Defer()
	control.Defer2()
	control.VariablesInDefer()
}

func TestLoadingRobots(t *testing.T) {
	control.PrintRobots()
}
