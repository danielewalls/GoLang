/* RZFeeser | Alta3 Research
Writing out JSON              */

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
    type ColorGroup struct {
        ID     int
        Name   string
        Colors []string
        P      Person `json:"Person"`
    }

    per := Person{Fn: "RZ",
        Ln: "Feeser",
    }

    group := ColorGroup{
        ID:     24601,
        Name:   "Greens",
        Colors: []string{"Moss", "Shamrock", "Lime", "Hunter"},
        P:      per,
    }
    
    // serialize values from struct to json format
    b, err := json.Marshal(group)
    
    if err != nil {
        fmt.Println("error:", err)
    }
    
    // print to standard out
    os.Stdout.Write(b)
}
