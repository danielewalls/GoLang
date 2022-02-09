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

// Begin Global Structures  and vars ------------------------------------------------------------

var file2readIn string = "wf_data.json"

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
	fmt.Println("\n   10. Readin Json Data, ready for data manipulation")
	fmt.Println("\n   11. Add User into JsonData")

	fmt.Println("\n   15. Simple Readin Marshal, unmarshal, save to wfSSimpleave.json")
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

	jsonFile, err := os.Open(file2readIn)
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

	clearScreen()
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
	var sentinal int = 0
	var users Users

	clearScreen()
	fmt.Printf("\nEnter full Gamertag (Case Insensitive): ")
	fmt.Scan(&inputGamertag)

	users = ReadinFile()
	for i := 0; i < len(users.Users); i++ {
		if (strings.ToUpper(inputGamertag) == strings.ToUpper(users.Users[i].Gamertag)) {
			sentinal = 1
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
	if (sentinal == 0) { fmt.Println("\n\nGamertag: " + inputGamertag + " was not found in the warframe datafile") }
}

func WriteOutFile(sentinUsers Users) {

	file, _ := json.MarshalIndent(sentinUsers, "", " ")
	_ = ioutil.WriteFile(file2readIn, file, 0644)
	fmt.Println("\n" + file2readIn + " was updated successfully, Press the Enter Key to Continue\n")
	fmt.Scanln()

}
// End Functions ------------------------------------------------------------


// Main kiosk processing loop
func main(){

	var SelectedItem string
	var allusers Users


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
			clearScreen()
			allusers = ReadinFile()
			fmt.Println("\n\nAll data successfully read, all set for data manipulation....returning to main menu")
			time.Sleep(2 * time.Second)

		case "11":
			fmt.Printf("\n%v is not implemented yet......", SelectedItem)
			time.Sleep(1 * time.Second)
			clearScreen()

		case "15":
			if (len(allusers.Users) > 0) {
				WriteOutFile(allusers)
			} else {
				fmt.Println("\nERROR: NO DATA found in mememory please select Option 10 to read in data")
				time.Sleep(3 * time.Second)
			}
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