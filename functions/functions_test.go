package functions_test

import (
	"fmt"
	"testing"

	"github.com/spinozanose/go-hello-world/functions"
)


func TestFunction(t *testing.T) {
	fmt.Println("Arguments passed by value:")
	oneThing := "thing1"
	anotherThing := "thing2"
	functions.PrintThings(oneThing, anotherThing)
	fmt.Println("Arguments passed by reference:")
	functions.PrintAndChangeThings(&oneThing, &anotherThing)
	fmt.Println(oneThing, anotherThing)
	fmt.Println("Variatic Functions:")
	fmt.Println(*functions.SumSomeInts(0, 1, 2, 3)) // because we returned a pointer we have to dereference
	fmt.Println("Another syntax:")
	fmt.Println(functions.SumSomeMoreInts(0, 1, 2, 3))
	fmt.Println("Multiple return (with error):")
	fmt.Println(functions.DivideFloats(2, 2)) //implicit conversion to floats
	returnValue, err := functions.DivideFloats(2, 0) // returns 0 and an error object
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(returnValue)
	fmt.Println("Anonymous Functions:")
	functions.AnonymousFunctions()
	fmt.Println("Methods:")
	functions.Methods()
}