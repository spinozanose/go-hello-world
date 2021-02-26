package variables

import (
	"fmt"
)

// declaring at the file level can only use full declaration format.
var l int = 99

// DeclareVariables is . . .
func DeclareVariables() {

	fmt.Println("DeclareVariables:")

	var i int //declare
	i = 42    //assign
	i = 27    //reassign

	// one line asssignment
	var j int = 21

	// or just let the compiler figure it out
	k := 11

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}

// DeclareTypes is . . .
func DeclareTypes() {

	fmt.Println("DeclareTypes:")

	var i int = 19
	fmt.Printf("%v, %T\n", i, i)

	var j float32 = 24
	fmt.Printf("%v, %T\n", j, j)
}

// Types is . . .
func Types() {

	fmt.Println("Types:")

	var m bool = true
	fmt.Printf("m %v, %T\n", m, m)

	n := 1 == 1
	p := 1 == 2
	fmt.Printf("n %v, %T\n", n, n)
	fmt.Printf("p %v, %T\n", p, p)
}

// Arrays is . . .
func Arrays() {

	fmt.Println("Arrays:")

	a := [3]int{1, 2, 3} // type specified
	fmt.Printf("a %v\n", a)
	fmt.Printf("length of a is %v\n", len(a)) // there is a len function
	b := [...]int{1, 2, 3} // the dots are used instead of declaring a size directly
	// Arrays are fixed size!
	c := b // this is a clone operation for Arrays
	c[0] = 4
	fmt.Printf("b %v\n", b)
	fmt.Printf("c %v\n", c)
	d := &b // this declares d to refer to the b array
	d[0] = 4
	fmt.Printf("b %v\n", b)
	fmt.Printf("d %v <-- Note the reference '&' sign \n", d)
}

// Slices is . . .
func Slices() {

	fmt.Println("Slices:")
	// Slices are not fixed sizes like arrays
	// A slice is backed by an array
	a := []int{1, 2, 3} // empty brackets indicate a slice
	fmt.Printf("a %v\n", a) // initialized with zeroes
	a = make([]int, 3) // we can also use the make function
	fmt.Printf("a %v\n", a)
	// The capacity function returns the size of the underlying array
	fmt.Printf("Capacity of a is %v\n", cap(a)) // note the capacity is the same
	a = make([]int, 3, 10) // the third number sets the capacity
	fmt.Printf("a %v\n", a)
	fmt.Printf("Capacity of a is %v\n", cap(a)) // note the capacity again
	a = make([]int, 0) // set to an empty slice
	fmt.Printf("a %v\n", a)
	fmt.Printf("Capacity of a is %v\n", cap(a))
	a = append(a, 1) // we can append to a slice
	a = append(a, 2, 3) // append is a variatic function, so this is good, too.
	fmt.Printf("a %v\n", a)
	fmt.Printf("Length of a is %v\n", len(a)) // can use functionality from arrays
	fmt.Printf("Capacity of a is %v\n", cap(a))
	b := a // slice declarations are naturally a reference type!
	fmt.Printf("b %v\n", b)
	b[1] = 4
	fmt.Printf("a %v\n", a)
	// slicing operations work on both arrays and other slices
	// first index is inclusive, second reference is exclusive
	c := a[:] // this creates a copy of the slice
	d := a[1:]
	e := a[:3]
	f := a[1:2]
	fmt.Printf("c %v\n", c)
	fmt.Printf("d %v\n", d)
	fmt.Printf("e %v\n", e)
	fmt.Printf("f %v\n", f)
	g := a[1:1] // notice that this retrieves an empty array!
	fmt.Printf("g %v\n", g)
	a = []int{1, 2, 3}
	// There is a spread operator in Go, for use with variatic functions
	// a = append(a, []int{4,5,6}) // this does not work
	a = append(a, []int{4,5,6}...) // but this does
	fmt.Printf("a %v\n", a)
	// there are stack operations in slices
	a = append(a, 7) // this is a push
	fmt.Printf("a %v\n", a)
	a = a[1:] // this is a pop (but does not return the element)
	fmt.Printf("a %v\n", a)
	a = a[:len(a)-1] // this is a pop from the tail
	fmt.Printf("a %v\n", a)
	// To return the element would involve writing a function
	//Removing an element from the slice:
	a = []int{1, 2, 3, 4, 5}
	a = append(a[:2], a[3:]...) // this removes the 3rd element, element 2
	fmt.Printf("a %v\n", a) 
	// Note that if you were to assign the created array to b, a would be 
	// messed up, as would anything that refers to a!
	a = []int{1, 2, 3, 4, 5}
	b = append(a[:2], a[3:]...)
	fmt.Printf("b %v\n", b)
	fmt.Printf("a %v\n", a)
	// If this behavior is a problem, then we would need to write a loop.
}

// Maps is . . .
func Maps()  {
	fmt.Println("Maps:")

	statePopulations := make(map[string]int) // using make function
	statePopulations = map[string]int { // standard way
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}
	// Keys have to be testable for equality
	// No slices, maps, or other functions
	// But arrays will work, believe it or not

	// Maps are not ordered, so this could print out in any order
	fmt.Println(statePopulations)
	// Just one element
	fmt.Printf("California population is %v\n", statePopulations["California"])
	// Adding
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations) // again, indeterminate order
	// Remove Georgia
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
	// Watch this, though
	fmt.Printf("Georgia population is %v\n", statePopulations["Georgia"])
	// This can cause lots of problems!
	pop, ok := statePopulations["Georgia"] // <-- this can be used to verify misspellings
	fmt.Println(pop, ok) // <-- if ok is false, it failed
	// if we don't care about the value, we can do this:
	_, ok = statePopulations["Georgia"]
	fmt.Println(ok)
	// len works with maps
	fmt.Printf("Length of statePopulations is %v\n", len(statePopulations))
	// Like slices, assignment is by reference
	sp := statePopulations
	sp["California"] = 42
	fmt.Printf("California population is %v in statePopulations now\n", statePopulations["California"])
	sp["California"] = 39250017
	fmt.Printf("California population is %v in statePopulations now\n", statePopulations["California"])
}
