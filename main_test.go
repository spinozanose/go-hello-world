package main_test

import (
	"testing"

	"github.com/spinozanose/go-hello-world/hello"
)

func TestHello(t *testing.T) {
	if hello.SayHello() != "Hello!" {
		t.Fatal("I'm so sad! You didn't say hello . . .")
	}
}
