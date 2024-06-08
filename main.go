package main

import "fmt"

// Function to change the value of an integer using a pointer
func changeValue(val *int) {
    *val = 42
}

// Function to demonstrate pointer to pointer
func changeValueDoublePointer(val **int) {
    **val = 58
}

// Function to modify a struct using a pointer
func changePerson(p *Person) {
    p.age = 25
}

// Person struct
type Person struct {
    name string
    age  int
}

func main() {
    // Basic pointer usage
    num := 10
    fmt.Println("Initial num:", num) // Output: Initial num: 10
    
    // Getting the pointer to num
    ptr := &num
    fmt.Println("Pointer to num:", ptr) // Output: Pointer to num: <memory_address>
    
    // Changing value using the pointer
    changeValue(ptr)
    fmt.Println("After changeValue:", num) // Output: After changeValue: 42
    
    // Pointers with structs
    person := Person{"Alice", 30}
    fmt.Println("Initial person:", person) // Output: Initial person: {Alice 30}
    
    // Getting the pointer to the person struct
    personPtr := &person
    
    // Changing struct field using the pointer
    changePerson(personPtr)
    fmt.Println("After changePerson:", person) // Output: After changePerson: {Alice 25}
    
    // Pointers to pointers
    ptrToPtr := &ptr
    fmt.Println("Pointer to pointer:", ptrToPtr) // Output: Pointer to pointer: <memory_address>
    
    // Changing value using pointer to pointer
    changeValueDoublePointer(ptrToPtr)
    fmt.Println("After changeValueDoublePointer:", num) // Output: After changeValueDoublePointer: 58
    
    // Using new to create a pointer
    numPtr := new(int)
    fmt.Println("Value at new pointer:", *numPtr) // Output: Value at new pointer: 0
    
    // Assigning a value to the newly created pointer
    *numPtr = 100
    fmt.Println("After assigning 100:", *numPtr) // Output: After assigning 100: 100
    
    // Nil pointers
    var nilPtr *int
    if nilPtr == nil {
        fmt.Println("nilPtr is nil") // Output: nilPtr is nil
    }
}





// Explanation:
// Basic Pointer Usage:

// changeValue(val *int) changes the value at the given pointer address.
// The pointer ptr holds the address of num.
// changeValue(ptr) modifies the value of num through the pointer.
// Pointers with Structs:

// Person struct has name and age fields.
// changePerson(p *Person) modifies the age field of the Person struct.
// personPtr holds the address of the person struct, allowing modification through changePerson(personPtr).
// Pointers to Pointers:

// ptrToPtr holds the address of ptr.
// changeValueDoublePointer(val **int) modifies the value through a pointer to a pointer.
// Using new to Create a Pointer:

// numPtr := new(int) allocates memory for an integer and returns its pointer.
// *numPtr is used to access and modify the value at the allocated memory.
// Nil Pointers:

// var nilPtr *int declares a nil pointer.
// if nilPtr == nil checks if the pointer is nil, which it is initially.
