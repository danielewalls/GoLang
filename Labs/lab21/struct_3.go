/* RZFeeser | Alta3 Research
   Create a struct named person */

   package main

   import "fmt"
   
   type person struct {
	   name string
	   age  int
   }
   
   func main() {
   
	   fmt.Println(person{"Bob", 20})
   
	   fmt.Println(person{name: "Alice", age: 30})	// fields can be referenced by name
   
	   fmt.Println(person{name: "Fred"}) 	// omitted fields are zero filled
   
	   s := person{age: 50, name: "Okon"}	// passing by name, not by position
	   fmt.Println(s.name)			// "Okon"
	   fmt.Println(s.age)			// 50
   
	   sp := &s				// creates a pointer to s
	   fmt.Println(sp.age)			// 50
   
	   sp.age = 51
	   fmt.Println(sp.age)			// pointer now reflects 51
	   fmt.Println(s.age)			// the struct also reflects 51
   }
   