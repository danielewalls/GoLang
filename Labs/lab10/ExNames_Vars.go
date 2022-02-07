/* Alta3 Research | RZFeeser
   Multiple functions   */

   package main

   import "fmt"
   
   //               (param, param, param) return
   func shippingCalc(x int, y int, z int) int {    // the return type is int
	   return x + y + z 
   }
   
   func main() {
	   fmt.Println("Enter Height: 42")    // we'll learn about inputting data later
	   fmt.Println("Enter Width: 13")
	   fmt.Println("Enter Length: 20")
	   fmt.Println("Total: ", shippingCalc(42, 13, 20))   // print the results of the function
   }
   