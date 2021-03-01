package interfaces

import (
	"bytes"
	"fmt"
	"io"
)

// Interfaces is . . .
func Interfaces() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello, Go!"))
	//
	fmt.Println("Incrementer:")
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 7; i++ {
		fmt.Println(inc.Increment())
	}

	// Type conversion. Conversion works if the functions in the struct are implemented
	writer, ok := w.(io.Reader)
	if ok {
		fmt.Println(writer, "This will not happen")
	} else {
		fmt.Println("Cannot convert incompatible types!")
	}

	var emptyInterface interface{} = w //Anything can be cast to an empty interface
	fmt.Println(emptyInterface)

}

// Writer is . . .
type Writer interface { // Convention for single-method interfaces: name the interface with -er
	Write([]byte) (int, error) // interfaces express behavior, i.e. functions
}

// Closer is . . .
type Closer interface {
	Close() error
}

// WriterCloser is an interface composed of two other interfaces
type WriterCloser interface {
	Writer
	Closer
}

//ConsoleWriter is . . .
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// Incrementer is . . .
type Incrementer interface {
	Increment() int
}

// IntCounter is . . .
type IntCounter int

// Increment is . . .
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// BufferedWriterCloser is . . .
type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

// Close is a method on the BufferedWriterCloser
func (bwc *BufferedWriterCloser) Close() error {
	return nil
}

// NewBufferedWriterCloser is a constructor of sorts
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
