/* RZFeeser | Alta3 Research
Troubleshooting - Writing out JSON & exported identifiers

Within the ColorGroup struct, names are created with lowercase.
Go visibility rules for struct files are slightly ammended, but still remain. (https://pkg.go.dev/encoding/json#Marshal)

Here we are requesting our struct interact with the json package. Therefore, any names that do not follow the rules of "Exported identifiers", will not be found by the json package.

Rules for Exported identifiers include the first character being capitalized, and being declared in the package block (or a field name or method name). (https://golang.org/ref/spec#Exported_identifiers)*/

package main

import (
    "encoding/json"
    "fmt"
    "os"
)

func main() {
    type Person struct {
        Fn string
        Ln string
    }
    // creating a struct with some lowercase names
    type ColorGroup struct {
        ID     int      // uppercase - "OK" at function level because it is a field name
        name   string   // lowercase
        colors []string // lowercase
        p      Person `json:"Person"`   // lowercase (struct, not JSON tag)
    }

    per := Person{Fn: "RZ",
        Ln: "Feeser",
    }

    // group is the TYPE ColorGroup
    group := ColorGroup{
        ID:     24601,
        name:   "Greens",
        colors: []string{"Moss", "Shamrock", "Lime", "Hunter"},
        p:      per,
    }
    
    // serialize values from struct to json format
    // our struct "group" is defined in the main package
    b, err := json.Marshal(group) // requesting interaction between json package and main package
    // result is ONLY the "exported" (capitalized) names will be found!
    
    if err != nil {
        fmt.Println("error:", err)
    }
    
    // print to standard out
    os.Stdout.Write(b) // the "ONLY" thing that the json package could find was ColorGroup.ID!
}
