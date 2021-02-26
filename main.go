package main

import (
	"fmt"

	"rsc.io/quote"

	"github.com/spinozanose/go-hello-world/hello"
	"github.com/spinozanose/go-hello-world/variables"
)

func main() {
	fmt.Println(hello.SayHello())
	fmt.Println(quote.Go())

	variables.DeclareVariables()
	variables.DeclareTypes()
	variables.Types()
	variables.Arrays()
	variables.Slices()
}
