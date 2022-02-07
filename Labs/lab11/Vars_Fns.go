/* Alta3 Research | RZFeeser
   Variables - Package vs function level */

   package main

   import (
	   "fmt"
   )
   
   // this variable is "global" for the "main" package
   // because it begins with a lowercase letter
   var zzz = "package-level (global) variable\n"
   
   
   func main() {
	   
	   // display our "global" variable
	   fmt.Printf(zzz)
	   
	   // This would be "long hand" declaration
	   var mug string = "filled with tea\n"
	   fmt.Printf(mug)
	   
	   // Either type or expression may be omitted (not both)
	   // These are all examples of type being omitted        
	   var a = "Hello"
	   var b = 23
	   var c = true
	   var d = 2.3
	   
	   // var a = "Goodbye"
	   a = "Goodbye"
   
	   fmt.Printf("The type of a, b, c, d is: %T, %T, %T and %T respectively\n", a, b, c, d)
	   // Prints: The type of a, b, c, d is: string, int, bool and float64 respectively
   }
   