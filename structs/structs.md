type structname struct {
    field type
    field type
}

	
    alex := person{"Alex", "Anderson"}
	alex := person{firstName: "Alex", lastName: "Murphy"}
	var alex person
	fmt.Println(alex)
    


Printf is used for interpolation formats
%+v is for printing the field and its value separately

	fmt.Printf("%+v\n", alex)

Structs are updated using the "." property
	
    alex.firstName = "Violet"
	alex.lastName = "Blue"

Multi-line struct declaration reqiure "," at the end 

	jim := person{
		firstName: "Jimi",
		lastName:  "Hendrix",
		contact: contactInfo{
			email:   "jim@google.com",
			zipCode: 98000,
		},
	}

## Struct shorthand property

    type person struct {
	firstName string
	lastName  string
	contactInfo   contactInfo
    }
is exactly the same as

    type person struct {
	firstName string
	lastName  string
	contactInfo
    }

Function parameters are pass by value by default

Turn address into value by using *
Turn value into address by suing &

go will implicitly cast the value type parameters to pointers wrt the function definition

	func (pointerToPerson *person) updateName(newFirstName string) {
		(*pointerToPerson).firstName = newFirstName
	}

 	jimPointer := &jim
	jimPointer.updateName("Jimmy")
is same as

	jim.updateName("Jimmy")

## Slices and Arrays

slice is a datastructure that has 3 elemets in it
1. pointer to the head
2. capacity
3. length

Passing slice as a function patrameter creates a copy of all the 3 elements. However, the pointer to head will pointing to the same address as the actual slice variable

|Value Types| Reference Types|
|-----------|----------------|
|int | slices|
|float | maps|
|string | channels|
|bool | pointers|
|structs | functions|

