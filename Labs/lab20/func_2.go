/* RZFeeser | Alta3 Research
   nil in the error position means "no error" */
   
   package main

   import (
	   "fmt"
	   "errors"
   )
   
   //name of funct   //input                      //output
   func rollchar(firstName string, lastName string) (string, error) {
	   return firstName + " the " + lastName, errors.New("can't work with 42")
   }
   
   func main() {
	   fmt.Println("Welcome to the Character Generator")
   
	   playerChar, err := rollchar("Gandalf", "Grey")
   
	   if err != nil {
			fmt.Println(err)
	   } else {
		   fmt.Println(playerChar, "has been generated.")
	   }
   }
   