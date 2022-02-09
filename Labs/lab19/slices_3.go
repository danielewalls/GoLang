/* RZFeeser | Alta3 Research
   Exploring slices - length vs capacity  */

   package main

   import "fmt"
   
   func main() {
   
	   numbers := make([]int, 3, 5)
	   fmt.Println("numbers =", numbers)
	   fmt.Println("length =", len(numbers))
	   fmt.Println("capacity =", cap(numbers))
   
	   //This line will cause a runtime error index out of range [4] with length 3
	   //numbers[4] = 5
   
	   //Increasing the length from 3 to 5
	   numbers = numbers[0:5]
   
	   numbers[4] = 5
   
	   fmt.Println("Increasing length from 3 to 5")
	   fmt.Println("numbers =", numbers)
	   fmt.Println("length =", len(numbers))
	   fmt.Println("capacity =", cap(numbers))
   
	   // just use append in most situations
	   numbers = append(numbers, 66, 777, 8888, 345)
	   fmt.Println("numbers =", numbers)
	   fmt.Println("length =", len(numbers))
	   fmt.Println("capacity =", cap(numbers))
	   fmt.Println("numbers =", numbers)
   }
   