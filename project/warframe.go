package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"encoding/json"
    "io/ioutil"
	// "strconv"
)

// Begin Global Structures ------------------------------------------------------------
// Main array of all the warframes Users
type Users struct {
    Users []User `json:"users"`
}

type User struct {
    Gamertag   string `json:"gamertag"`
    Warframe   string `json:"warframe"`
    Primary    string `json:"primary"`
	Secondary  string `json:"secondary"`
	Melee      string `json:"melee"`
    Companion Companion `json:"companion"`
}

type Companion struct {
    Name    string `json:"name"`
    Weapon  string `json:"weapon"`
}

// End Global Structures ------------------------------------------------------------


// Begin Functions ------------------------------------------------------------
// Clears screen with no return
func clearScreen() {

	clrScreen := exec.Command("clear")
	clrScreen.Stdout = os.Stdout
	clrScreen.Run()
}

// Displays main menu and returns user input for validation and processing
func DisplayMainMenu() string {

	var menuChoice string

	fmt.Println("\n\nWelcome to the Warframe User Relay Main Menu")
	fmt.Println("--------------------------------------------\n")
	fmt.Println("   1. Display a Specific User's Frames")
	fmt.Println("   5. Display All User's Frames")
	fmt.Println("\n   10. Add User frame Data")
	fmt.Println("\n   15. Save all data to JSON file")
	fmt.Println("")
	fmt.Println("   99. Exit Menu")
	fmt.Println("\n--------------------------------------------")

	fmt.Printf("\nEnter your Choice: ")
	fmt.Scan(&menuChoice)
	return menuChoice
}

// End Functions ------------------------------------------------------------

// Read in json file and display all data within that file - no return for now
func DisplayAllData() {

	jsonFile, err := os.Open("wf_data.json")
	if err != nil { 
		fmt.Println(err) 
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
        fmt.Println(users.Users[i])
    }

}
// End Functions ------------------------------------------------------------


// Main kiosk processing loop
func main(){

	var SelectedItem string

	// Infinite loop for kiosk app main menu
	for {
		clearScreen()
		SelectedItem = DisplayMainMenu()

		switch SelectedItem {
		case "1":
			DisplayAllData()
			fmt.Println("\nPress the Enter Key to Continue\n")
			fmt.Scanln()

		case "5":
			fmt.Printf("\n%v is not implemented yet......", SelectedItem)
			time.Sleep(1 * time.Second)
			clearScreen()

		case "10":
			fmt.Printf("\n%v is not implemented yet......", SelectedItem)
			time.Sleep(1 * time.Second)
			clearScreen()

		case "15":
			fmt.Printf("\n%v is not implemented yet......", SelectedItem)
			time.Sleep(1 * time.Second)
			clearScreen()
			
		case "99":
			fmt.Println("\n")
			os.Exit(0) 

		default:
			fmt.Printf("\n%v is NOT valid!!!", SelectedItem)
			time.Sleep(1 * time.Second)
			clearScreen()
		}
	}

}