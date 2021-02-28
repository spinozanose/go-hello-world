package functions

import "fmt"

// PrintThings is . . .
func PrintThings(oneThing, anotherThing string){ // Types can be combined in the function signature
	fmt.Println(oneThing, anotherThing)
}

// PrintAndChangeThings is . . .
func PrintAndChangeThings(oneThing, anotherThing *string){
	fmt.Println(*oneThing, *anotherThing)
	*oneThing = "banana pudding"
}

// SumSomeInts is . . .
func SumSomeInts(ints ...int) *int { //this is a variatic function, and it must be the last parameter
	// the variatic parameter is received a a slice
	fmt.Println(ints)
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return &sum // Look at this! We can return a pointer to the variable, so no copying!
}

// SumSomeMoreInts is . . .
func SumSomeMoreInts(ints ...int) (sum int) { // another way to declare the return type
	fmt.Println(ints)
	for _, v := range ints {
		sum += v // Notice that the sum variable has been initialized
	}
	return // And the Function will return sum, because it was declared in the signature
}

// DivideFloats is . . .
func DivideFloats(a, b float64) (float64, error) { //This is a very common mechanism in Go
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a/b, nil
}

//AnonymousFunctions is . . .
func AnonymousFunctions() {
	func() {
		fmt.Println("Hello, I have no name!")
	}()

	message := "Hello, neither do I!"
	func(msg string) {
		fmt.Println(msg) 
	}(message) // It is best practice to pass in values needed from the outer scope

	f :=func() {
		fmt.Println("Hello, I am anonymous as well!")
	}
	f()
}

type greeter struct {
	greeting string
	name string
}

// It is perfect possible to put a method like this on any custom type
// Notice that greeter is passed by value here
// We could pass by reference, too
func (g greeter) greet() { 
	fmt.Printf("%v, %v!", g.greeting, g.name)
}

// Methods is . . .
func Methods() {
	g := greeter{
		greeting:"Hello", 
		name: "Human"}
	g.greet()
}