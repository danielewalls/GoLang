/* RZFeeser | Alta3 Research
Reading JSON - with structs  */

package main

import (
    // "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    // "strconv"
)

func main() {
    // Open our jsonFile
    jsonFile, err := os.Open("LinWin_Lvl1.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened users.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

}
