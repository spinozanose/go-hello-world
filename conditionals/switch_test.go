package conditionals_test

import (
	"fmt"
	"testing"

	"github.com/spinozanose/go-hello-world/conditionals"
)

func TestSwitches(t *testing.T) {
	conditionals.Switches(1)
	conditionals.Switches(2)
	conditionals.Switches(3)
	//
	fmt.Println("GuessTheNumber2:")
	fmt.Printf("Number -12: %v\n", conditionals.GuessTheNumber2(-12))
	fmt.Printf("Number 101: %v\n", conditionals.GuessTheNumber2(101))
	fmt.Printf("Number 10: %v\n",  conditionals.GuessTheNumber2(10))
	fmt.Printf("Number 60: %v\n",  conditionals.GuessTheNumber2(60))
	fmt.Printf("Number 50: %v\n",  conditionals.GuessTheNumber2(50))
	//
	fmt.Println("Fallthrough with 9:")
	conditionals.Fallthrough(9) // this produces an incorrect result because of fallthrough!
	fmt.Println("Fallthrough with 10:")
	conditionals.Fallthrough(10)
	fmt.Println("Fallthrough with 11:")
	conditionals.Fallthrough(11)
	//
	fmt.Println("TypeCheck:")
	conditionals.TypeCheck(1)
	conditionals.TypeCheck(1.1)
	conditionals.TypeCheck("w")
	conditionals.TypeCheck(false)
	conditionals.TypeCheck([3]int{})
}