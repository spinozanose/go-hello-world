package control_test

import (
	"testing"

	"github.com/spinozanose/go-hello-world/control"
)

func TestPanicAndRecover(t *testing.T) {
	control.PanicAndRecover()
}

func TestPanic(t *testing.T) {
	control.Panic()
}

func TestPanicAndDefer(t *testing.T) {
	control.PanicAndDefer()
}

func TestHTTPPanic(t *testing.T) {
	control.HTTPPanic("/", "Hello, World!")
}