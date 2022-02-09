/* Alta3 Research | RZFeeser
   Functions - single function   */ 
   
   package main

   import (
	   "fmt"
   )
   
	 //name of funct   //input                      //output
   func Rollchar(firstName string, lastName string) (string) {
	 fullname := firstName + " the " + lastName 
	 return fullname
   }
   
   func main() {
	   fmt.Println("Welcome to the Character Generator")
   
	   playerChar := Rollchar("Gandalf", "Grey")
	   fmt.Println(playerChar)
   }
   