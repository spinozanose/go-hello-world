package goroutines_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/spinozanose/go-hello-world/goroutines"
)

var waitGroup = sync.WaitGroup{}

func TestGoRoutines(t *testing.T) {
	go goroutines.SayHello() // <-- note the keyword go, which creates a green thread
	// we can also put go on an anonymous function
	var msg = "Hello"
	go func() {
		fmt.Println(msg) // go has closures, so this works, but not a good idea
	}() // in this case, if we want to share data, maybe sending a pointer is better
	msg = "Goodbye"       // and this is why relying on closures is a bad idea
	go func(msg string) { //this is better
		fmt.Println(msg)
	}(msg)
	// we have to put this in or the spawned threads can be closed by the main program before it can run
	// But this is bad practice
	time.Sleep(100 * time.Microsecond)
}

func TestWaitGroup(t *testing.T) {
	var msg = "Hello"
	waitGroup.Add(1) // this tells the wait group that there is a new thread about to be spawned
	go func(msg string) {
		fmt.Println(msg)
		waitGroup.Done() //decrement the number of threads in the group to wait on
	}(msg)
	waitGroup.Wait() // this is inefficient, but often necessary
}

var counter = 0
var m = sync.RWMutex{} //multiple readers, write is locked

//runtime.GOMAXPROCS(1) <--single threaded
//runtime.GOMAXPROCS(4) <--one thread per core is typical (and default)
//runtime.GOMAXPROCS(-1) <--does not change, just reads

// Best Practices:
// - Don't create goroutines in libraries (unless returning channels)
// - To avoid memory leaks, make sure you know how a routine will end
// - Check for race conditions at compile time
//    - go run -race

func TestSynchronization(t *testing.T) {
	// This actually eliminates the benefits of parallelism because it now
	// functions sequentially.
	for i := 0; i < 10; i++ {
		waitGroup.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	waitGroup.Wait()
}

func sayHello() {
	fmt.Printf("Hello, %v!\n", counter)
	m.RUnlock()
	waitGroup.Done()
}

func increment() {
	counter++
	m.Unlock()
	waitGroup.Done()
}
