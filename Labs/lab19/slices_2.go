/* RZFeeser | Alta3 Research
   Exploring slices - copy and append       */
   
   package main

   import "fmt"
   
   func main() {
   
	   s := make([]string, 3)             // typical usage of "make" - length AND capacity are set to 3
	   fmt.Println("slice s:", s)
   
	   s[0] = "a"
	   s[1] = "b"
	   s[2] = "c"
	   fmt.Println("slice s:", s)         // slice s: ["a", "b", "c"]
	   fmt.Println("get s[2]:", s[2])     // get s[2]: "c"
   
	   fmt.Println("len s:", len(s))
   
	   s = append(s, "d")                 // ["a", "b", "c", "d"]
	   s = append(s, "e", "f")            // ["a", "b", "c", "d", "e", "f"]
	   fmt.Println("append s:", s)        // append s: ["a", "b", "c", "d", "e", "f"]
   
	   c := make([]string, len(s))
	   copy(c, s)
	   fmt.Println("copy c:", c)
   
	   // after the copy, changes to "c" no longer effect "s"
	   // as "c" is a new slice
	   c[2] = "updated"
	   fmt.Println("original slice s:", s)
	   fmt.Println("copy slice c:", c)
   
	   l := s[2:5]
	   fmt.Println("sl1:", l)             // sl1: [c d e]
   
	   l = s[:5]
	   fmt.Println("sl2:", l)             // sl2: [a b c d e]
   
	   l = s[2:]
	   fmt.Println("sl3:", l)             // sl3: [c d e f]
   
   }
   