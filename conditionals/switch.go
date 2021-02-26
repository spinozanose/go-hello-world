package conditionals

import "fmt"

// Switches is . . .
func Switches(in int) {

	switch i := in; i { // Can use an initializer
	case 1:
		fmt.Println("You selected 1!")
	case 2:
		fmt.Println("You selected 2!")
	case 3, 4, 5:
		fmt.Println("You selected 3, 4, or 5!")
	default:
		fmt.Printf("You selected another number, %v!\n", i)
	}
}

// GuessTheNumber2 is . . .
func GuessTheNumber2(guess int) bool {

	number := 50

	switch {
	case guess == number:
		fmt.Println("Nailed it!")
		return true
	case guess < 1 || guess > 100:
		fmt.Println("Guess must be between 1 and 100!")
	case guess < number:
		fmt.Println("Not there yet!")
	case guess > number:
		fmt.Println("Too high!")
	}
	return false
}

// Fallthrough is . . .
func Fallthrough(i int) {
	// There is no "falling through" in Go by default, but there is a 
	// fallthrough keyword
	switch {
	case i < 10:
		fmt.Println("less than 10")
		fallthrough //warning: this does not execute the logic in the next case!
	case i > 10:
		fmt.Println("greater than 10")
		break //but there is a break statement, used to bail early
		//fmt.Println("This never prints!")
	default:
		fmt.Println("equals 10")
	}
}

// TypeCheck is . . .
func TypeCheck(i interface{}) {

	switch i.(type) {
	case int:
		fmt.Println("Integer")
	case float64:
		fmt.Println("Float64")
	case string:
		fmt.Println("String")
	case [3]int:
		fmt.Println("Array of 3 Integers")
	default:
		fmt.Println("Another Type")
	}

}