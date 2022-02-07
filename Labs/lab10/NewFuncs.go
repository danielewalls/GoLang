/* Alta3 Research | RZFeeser
   Improved design patterns for functions   */

   package main

   import "fmt"
   
   // func shippingCalc(x int, y int, z int) int {    // this line commented out
   func shippingCalc(x, y, z int) int {               // "best practice" 
	   return x + y + z 
   }
   
   func main() {
	   fmt.Println("Enter Height: 42")    // we'll learn about inputting data later
	   fmt.Println("Enter Width: 13")
	   fmt.Println("Enter Length: 20")
	   fmt.Println("Total: ", shippingCalc(42, 13, 20))   // print the results of the function
   }
   