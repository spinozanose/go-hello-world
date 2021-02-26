package control

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Defer is . . .
func Defer() {
	fmt.Println("One deferred statement")
	fmt.Println("First")
	// defers until after the func is completed, but before the function returns
	defer fmt.Println("Second") 
	fmt.Println("Third")
}

// Defer2 is . . .
func Defer2() {
	fmt.Println("Multiple deferred statements")
	defer fmt.Println("First")
	defer fmt.Println("Second") 
	defer fmt.Println("Third")
	// deferred statements are executed in LIFO order
	// that order helps with closing resources
}

// PrintRobots is . . .
func PrintRobots() {
	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

// VariablesInDefer is . . .
func VariablesInDefer() {
	fmt.Println("VariablesInDefer:")
	a := "start"
	defer fmt.Println(a) // This will print out "start"
	a = "end"
}