/* Packages
   Every Go program is made up of packages.
   Programs start running in the package main.
   This program uses the packages with import paths "fmt" and "math/rand"
*/

// Programs start running in package "main"
package main

// This program uses the packages with import paths "fmt" and "math/rand"
import (
    "fmt"
    "math/rand"
)

// Creating a function main()
func main() {
    fmt.Println("Looks like you're ready to learn GoLang with Alta3 Research")
    fmt.Println("My favorite number is 42")
    fmt.Println("However, I grab random numbers as well", rand.Intn(10))
}
