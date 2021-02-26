package pointers

import "fmt"

// Pointers is . . .
func Pointers() {
	fmt.Println("Normal behavior:")
	a := 42
	fmt.Println(a)
	b := a // These are "value types", and the value was copied
	fmt.Println(a, b)
	a = 429 // So, change a and b stays the same
	fmt.Println(a, b)
	//
	fmt.Println("Using pointers:")
	var c int = 42
	var d *int = &c // d is declared with a pointer that points to an int
	fmt.Println(c, d) // For d, this prints the memory address of the int being pointed to
	fmt.Println(&c, d) // You can see that but printing with the addressof operator
	fmt.Println(c, *d) // To print the value of d, use the dereferencing operator
	c = 429 // So, change c and d changes, too!
	fmt.Println(c, *d)
	*d = 195 // This also works if we change d (note the dereferencing operator)
	fmt.Println(c, *d)
	// Note: there is no pointer arithmatic like in C and C++, except in the unsafe package. But be wary.
	
	fmt.Println("Pointers to structs:")
	var ms *mystruct // assign ms to a pointer to a mystruct type
	ms = &mystruct{foo:42} // create a mystruct pointer to an anonymous object
	fmt.Println(ms)
	//
	var ms2 *mystruct // assign ms to a pointer to a mystruct type
	fmt.Println(ms2)  // right now this is pointing to <nil>
	ms2 = new(mystruct) // create an empty mystruct object with new keyword
	fmt.Println(ms2) // now pointing to a struct with a value 0
	(*ms2).foo = 42 // assigning a value to the dereferenced ms2
	fmt.Println(ms2) // to view the struct
	fmt.Println((*ms2).foo) // to view foo
	// however, we don't need to derefence, because the compiler knows what to do
	ms2.foo = 42 
	fmt.Println(ms2.foo) // to view foo
}

type mystruct struct {
	foo int
}

// PointerAssignmentBehavior is . . .
func PointerAssignmentBehavior() {
	fmt.Println("Pointer Assignment Behavior (Arrays):")
	array1 := [3]int{1,2,3}
	array2 := array1
	fmt.Println(array1, array2)
	array1[1] = 42
	fmt.Println(array1, array2)
	fmt.Println("Pointer Assignment Behavior (Slices):")
	slice1 := []int{1,2,3}
	slice2 := slice1
	fmt.Println(slice1, slice2)
	slice1[1] = 42
	fmt.Println(slice1, slice2)
	fmt.Println("Pointer Assignment Behavior (Maps):")
	map1 := map[string]string {"foo":"bar", "baz": "biz"}
	map2 := map1
	fmt.Println(map1, map2)
	map1["foo"] = "another value"
	fmt.Println(map1, map2)
}