package variables

import (
	"fmt"
	"reflect"
)

// Doctor is . . .
// Notice naming convention, particularly the case of the names
// If the fields should be seen outside of the package they must be upper case
type Doctor struct {
	number     int
	actor      string
	companions []string
}

// Assignments is . . .
func Assignments() {
	fmt.Println("Assignments:")
	// You can assign the variables without the names, but probably not a good idea
	aDoctor := Doctor{
		number:     3,
		actor:       "John Pertwee",
		companions: []string {
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor)
	fmt.Printf("aDoctor.actor is %v\n", aDoctor.actor)
	fmt.Printf("aDoctor.companions is %v\n", aDoctor.companions)
	fmt.Printf("aDoctor.companions[1] is %v\n", aDoctor.companions[1])

	anotherDoctor := struct{name string}{name: "Matt Smith"} // anonymous struct
	fmt.Println(anotherDoctor)

	scarfedDoctor := aDoctor // duplicates on assignment, not by reference
	scarfedDoctor.actor = "Tom Baker"
	fmt.Printf("aDoctor is %v\n", aDoctor)
	fmt.Printf("scarfedDoctor is %v\n", scarfedDoctor)

	errorDoctor := &aDoctor // we can assign by reference, though
	errorDoctor.actor = "Chris Eccleston"
	fmt.Printf("aDoctor is %v\n", aDoctor)
	fmt.Printf("errorDoctor is %v\n", errorDoctor)
}

// Animal is . . .
type Animal struct {
	Name	string `required:"true" max:"100"`
	Origin 	string
}

// Bird is . . .
type Bird struct {
	Animal // notice that this is not a named field, it is just an embedding
	CanFly bool
	SpeedKPH float32
}

// Composition is . . .
func Composition() {
	fmt.Println("Composition:")
	// Go does not have inheritence, but it does have composition
	aBird := Bird{}
	aBird.Name = "Emu"
	aBird.Origin = "Australia"
	aBird.SpeedKPH = 42
	aBird.CanFly = false
	fmt.Println(aBird)
	// However, by type a Bird is NOT an Animal

	// This is the literal sysntax
	aBird = Bird{
		Animal: Animal{ Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly: false,
	}
	fmt.Println(aBird)

	//However, to describe common behavior we will need Interfaces
}

// Tags is . . .
func Tags() {
	fmt.Println("Tags:")
	// Tags are just labels on fields in a struct. This example 
	// presumes a validation framework that uses the tag.
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}