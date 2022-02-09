/* RZFeeser | Alta3 Research
   Slices relationship to arrays     */

   package main

   import "fmt"
   
   func main() {
	   array := [5]int{1, 2, 3, 4, 5}
	   slice := array[:]
   
	   //Modifying the slice
	   slice[1] = 7
	   fmt.Println("Modifying Slice")
	   fmt.Println("Array =", array)
	   fmt.Println("Slice =", slice)
   
	   //Modifying the array. Would reflect back in slice too
	   array[1] = 2
	   fmt.Println("Modifying Underlying Array")
	   fmt.Println("Array =", array)
	   fmt.Println("Slice =", slice)
   }
   