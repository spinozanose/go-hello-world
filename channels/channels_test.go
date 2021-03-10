package channels_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Channels are a mechanism to pass data between goroutines
// Channels can be two-way
// Channels pass messages

var wg = sync.WaitGroup{}

func TestSimpleChannel(t *testing.T) {

	// channels HAVE to be set up using the make function
	ch := make(chan int) // Channels are strongly typed. This channel can only pass ints
	wg.Add(2)            // adding two items to the wait group
	go func() {          // anonymous function in a goroutine. This is the receiving goroutine
		i := <-ch // arrow pointing from the channel into the variable, so this receives the message from the channel
		fmt.Println(i)
		wg.Done() // subtracts 1 from the count in wg
	}() // this starts the routine
	go func() { // This is the sending goroutine
		ch <- 42 //arrow pointing to the channel, so this is to push it in (this is pass by value)
		wg.Done()
	}()
	wg.Wait() //waits until count is zero
}

func TestMultipleSourcesSinks(t *testing.T) {
	// Channels can have multiple sources and sinks, and these must be balanced.
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch // This will block until there is an int to receive
			fmt.Println(i)
			// because this finishes after receiving one message, there must be the same number
			// of senders and receivers on this channel.
			wg.Done()
		}()
		go func() {
			// Channels can only hold one thing at a time.
			// This is an unbuffered channel, so this will block until there is a space in the channel
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestBiDirectionalChannels(t *testing.T) {
	ch := make(chan int)

	wg.Add(2)
	go func() {
		i := <-ch // this pulls off
		fmt.Println(i)
		ch <- 27 // this puts back
		wg.Done()
	}()
	go func() {
		ch <- 42          // this puts
		fmt.Println(<-ch) // this pulls out
		wg.Done()
	}()
	wg.Wait()
}

func TestMonoDirectionalChannels(t *testing.T) {
	ch := make(chan int)

	wg.Add(2)
	go func(ch <-chan int) { // set for receive-only
		i := <-ch
		fmt.Println(i)
		// ch <- 27 // this is now disallowed
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // set for send-only
		ch <- 42
		// fmt.Println(<-ch) // this is now disallowed
		wg.Done()
	}(ch)
	wg.Wait()
}

func TestBufferedChannels(t *testing.T) {

	ch := make(chan int, 50) //set channel buffer to 50 messages

	for j := 0; j < 50; j++ {
		wg.Add(2)
		go func(ch <-chan int) {
			time.Sleep(1 * time.Second) //This puts a delay into the receivers
			i := <-ch
			fmt.Println(i, len(ch)) // len shows the number of messages
			wg.Done()
		}(ch)
		// these start up and immediately write into the channel
		go func(ch chan<- int, index int) {
			ch <- index
			wg.Done()
		}(ch, j)
	}
	wg.Wait()
}

func TestRangingOverTheChannel(t *testing.T) {
	ch := make(chan int, 50)

	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch { // This blocks until the channel is closed
			fmt.Println(i, len(ch))
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch) // this closes the channel, which causes a break in the range in the for loop in the receiver
		wg.Done()
	}(ch)

	wg.Wait()
}

func TestTestingForClosedChannel(t *testing.T) {
	ch := make(chan int, 50)

	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i, len(ch))
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch) // this closes the channel, which causes a break in the range in the for loop in the receiver
		wg.Done()
	}(ch)

	wg.Wait()
}

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)

func TestLogger(t *testing.T) {
	go logger()    //This starts a function that outputs log messages.
	defer func() { // This closes the channel after the function completes. If this were a main method, this would be when the application ends.
		close(logCh)
	}()

	// let's add a log entry
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	time.Sleep(5 * time.Second) // Application runs here

	// then end
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond) // gives the logger time to finish

}

func logger() {
	for entry := range logCh {
		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	}
}

// this construct, a type of a struct with no fields, cannot carry a message. It is a
// "signal-only" channel. It can only have the properties of whether a message was sent or received.
var doneCh = make(chan struct{})

func TestSelectLogger(t *testing.T) {

	go selectLogger()

	// let's add a log entry
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	time.Sleep(5 * time.Second) // Application runs here

	// then end
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond) // gives the logger time to finish
	doneCh <- struct{}{}               // this fails to send anything through the channel, but does carry a signal
}

func selectLogger() {
	for {
		select { // this like a switch statement, but only used for channels
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			fmt.Println("Logger shutting done, too")
			break
			// with no default case, this is a blocking select. Add a default case and it no longer blocks.
		}
	}
}
