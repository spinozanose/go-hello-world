package control

import (
	"fmt"
	"log"
	"net/http"
)

// Panic is . . .
func Panic() {
	fmt.Println("Panic!")
	fmt.Println("First")
	panic("something bad happened!")
	//fmt.Println("Second") These never execute
	//fmt.Println("Third")  These never execute
}

// PanicAndDefer is . . .
func PanicAndDefer() {
	fmt.Println("Panic!")
	fmt.Println("First")
	defer fmt.Println("Second") // Now we get "Second", and the panic happens afterward!
	panic("something bad happened!")
	//fmt.Println("Third")  This never executes
}

// PanicAndRecover is . . .
func PanicAndRecover() {
	fmt.Println("Panic!")
	fmt.Println("First")
	defer func() {
		// recover means the *calling* function continues. recover is only used in deffered functions
		if err := recover(); err != nil { 
			log.Println("Error: ", err)
			// if you needed to, here you can call panic again
		}
	}()
	panic("something bad happened!")
	// fmt.Println("Second") // These never execute
	// fmt.Println("Third")  // These never execute
}

// HTTPPanic is . . .
func HTTPPanic(path string, reply string) {
	
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(reply))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}

	// This will result in a panic because the port is blocked:
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(reply))
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}