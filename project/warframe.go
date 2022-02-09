package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"encoding/json"
	"io/ioutil"
	"strings"
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

	fmt.Println("\n\n  Welcome to the Warframe User Information Kiosk")
	fmt.Println("--------------------------------------------------\n")
	fmt.Println("   1. Display a Specific User's Frames")
	fmt.Println("   5. Display All User's Frames")
	fmt.Println("\n   10. Add User frame Data")
	fmt.Println("\n   15. Save all data to JSON file")
	fmt.Println("")
	fmt.Println("   99. Exit Menu")
	fmt.Println("\n--------------------------------------------------")

	fmt.Printf("\nEnter your Choice: ")
	fmt.Scan(&menuChoice)
	return menuChoice
}

//REad in file and return the list of users
func ReadinFile() Users {

	var warframeUsers Users

	jsonFile, err := os.Open("wf_data.json")
	if err != nil { 
		fmt.Println("ERROR Opening File:")
		fmt.Println(err) 
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &warframeUsers)
	return warframeUsers
}

// display all data within that file
func DisplayAllData() {

	var users Users

	users = ReadinFile()
	for i := 0; i < len(users.Users); i++ {

		// Module needed here if have time for color
		// yellow := color.New(FgYellow).SprintFunc()
		// red := color.New(FgRed).SprintFunc()
		// fmt.Printf("this is a %s and this is %s.\n", yellow("warning"), red("error"))

		fmt.Println("\n--------------------------------------------")
		fmt.Println("Gamer: " + users.Users[i].Gamertag)
		fmt.Println(" Frame:              " + users.Users[i].Warframe)
		fmt.Println("  Primary Weapon:    " + users.Users[i].Primary)
		fmt.Println("  Secondary Weapon:  " + users.Users[i].Secondary)
		fmt.Println("  Melee Weapon:      " + users.Users[i].Melee)
		fmt.Println("   Companion Name:   " + users.Users[i].Companion.Name)
		fmt.Println("   Companion Weapon: " + users.Users[i].Companion.Weapon)
    }
}

// Display asked for user if found
func DisplaySingleUser(){

	var inputGamertag string
	var users Users

	fmt.Printf("\nEnter full Gamertag (Case Insensitive): ")
	fmt.Scan(&inputGamertag)
	clearScreen()

	users = ReadinFile()
	for i := 0; i < len(users.Users); i++ {
		if (strings.ToUpper(inputGamertag) == strings.ToUpper(users.Users[i].Gamertag)) {
			fmt.Println("\n-----------------FOUND----------------------\n")
			fmt.Println("Gamer: " + users.Users[i].Gamertag)
			fmt.Println(" Frame:              " + users.Users[i].Warframe)
			fmt.Println("  Primary Weapon:    " + users.Users[i].Primary)
			fmt.Println("  Secondary Weapon:  " + users.Users[i].Secondary)
			fmt.Println("  Melee Weapon:      " + users.Users[i].Melee)
			fmt.Println("   Companion Name:   " + users.Users[i].Companion.Name)
			fmt.Println("   Companion Weapon: " + users.Users[i].Companion.Weapon)
			fmt.Println("\n--------------------------------------------")
		}
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
			DisplaySingleUser()
			fmt.Println("\nPress the Enter Key to Continue\n")
			fmt.Scanln()

		case "5":
			DisplayAllData()
			fmt.Println("\nPress the Enter Key to Continue\n")
			fmt.Scanln()

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