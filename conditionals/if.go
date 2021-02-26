package conditionals

import (
	"fmt"
	"math"
)

// IfStatements is . . .
func IfStatements() {

	if true {
		fmt.Println("True")
	}

	statePopulations := map[string]int { // standard way
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}
	// with initializer
	// Note that both variables are scoped inside the if block
	if pop, ok:=statePopulations["Texas"]; ok {
		fmt.Println(ok)
		fmt.Println(pop)
	}
}

// GuessTheNumber is . . .
func GuessTheNumber(guess int) bool {
	number := 50

	// || an && can chain conditionals, as expected, and short-circuit, too
	if guess < 1 || guess >100 {
		fmt.Println("Guess must be between 1 and 100!")
	} else {
		fmt.Println("Checking your guess . . .")
	}
	if guess == number {
		fmt.Println("Nailed it!")
		return true
	}
	if guess < number {
		fmt.Println("Not there yet!")
	} else if guess > number {
		fmt.Println("Too high!")
	}
	return false
	// There are also the usual other operators: !, >=, <=, and !=
}

// CompareFloatingPointNumbers is . . .
func CompareFloatingPointNumbers() {
	fpNum := 1.0123
	// The square of the square root should be the same number, right?
	if math.Abs(fpNum / math.Pow(math.Sqrt(fpNum), 2) - 1) < .001 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
