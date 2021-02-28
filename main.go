package main

import (
	"fmt"

	"rsc.io/quote"

	"github.com/spinozanose/go-hello-world/hello"
)

func main() {
	fmt.Println(hello.SayHello())
	fmt.Println(quote.Go())
}
