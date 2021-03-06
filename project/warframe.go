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
	fmt.Println("    1. Display a Specific User's Frames")
	fmt.Println("    5. Display All User's Frames")
	fmt.Println("")
	fmt.Println("   10. Read-in Json Data, ready for data manipulation")
	fmt.Println("   11. Add User into JsonData (noSpaces, use '_')")
	fmt.Println("   15. Save all in-memory data to json datafile")
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

func WriteOutFile(sentinUsers Users) {

	file, _ := json.MarshalIndent(sentinUsers, "", " ")
	_ = ioutil.WriteFile(file2readIn, file, 0644)
	fmt.Println("\n" + file2readIn + " was updated successfully, Press the Enter Key to Continue\n")
	fmt.Scanln()
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

func EnterNewUser(sentinUsers Users) Users {

	var newGamerTag string
	var sentinal int = 0
	var newUser User 
	var stdUserinput string


	fmt.Printf("\nEnter new Gamertag: ")
	fmt.Scan(&newGamerTag)

	// Check to see if user already in data set
	for i := 0; i < len(sentinUsers.Users); i++ {
		if (strings.ToUpper(newGamerTag) == strings.ToUpper(sentinUsers.Users[i].Gamertag)) { 
			sentinal = 1 
			break
		}
	}

	// If not in dataset continue on with user input
	if (sentinal == 0) {

		// Brute force all the questions for the new record
		newUser.Gamertag = newGamerTag

		fmt.Printf("\n" + newGamerTag + "'s warframe is called: ")
		fmt.Scan(&stdUserinput)
		newUser.Warframe = stdUserinput

		fmt.Printf("\n" + newGamerTag + "'s Primary weapon is: ")
		fmt.Scan(&stdUserinput)
		newUser.Primary = stdUserinput

		fmt.Printf("\n" + newGamerTag + "'s Secondary weapon is: ")
		fmt.Scan(&stdUserinput)
		newUser.Secondary = stdUserinput

		fmt.Printf("\n" + newGamerTag + "'s Melee weapon is: ")
		fmt.Scan(&stdUserinput)
		newUser.Melee = stdUserinput

		fmt.Printf("\n" + newGamerTag + "'s Companion is: ")
		fmt.Scan(&stdUserinput)
		newUser.Companion.Name = stdUserinput

		fmt.Printf("\n" + newGamerTag + "'s Companion " + stdUserinput + " is assigned a weapon called: ")
		fmt.Scan(&stdUserinput)
		newUser.Companion.Weapon = stdUserinput

		// Apprend the new record to exiting array and return
		sentinUsers.Users = append(sentinUsers.Users, newUser)		
	} else {
		fmt.Println("\n\nGamertag: " + newGamerTag + " ALREADY exists in the warframe datafile... Returning to main menu...")
		time.Sleep(2 * time.Second)
	}

	//regardless of error or new user return the dataset (updated or not)
	return sentinUsers

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
		case "1": // Display single user
			DisplaySingleUser()
			fmt.Println("\nPress the Enter Key to Continue\n")
			fmt.Scanln()

		case "5": // Display all users
			DisplayAllData()
			fmt.Println("\nPress the Enter Key to Continue\n")
			fmt.Scanln()

		case "10":  //Read in raw data file into memory
			clearScreen()
			allusers = ReadinFile()
			fmt.Println("\n\nAll data successfully read, all set for data manipulation....returning to main menu")
			time.Sleep(2 * time.Second)

		case "11": // Add new user to the array
			if (len(allusers.Users) > 0) {
				allusers = EnterNewUser(allusers)
			} else {
				fmt.Println("\nERROR: NO DATA found in memory please select Option 10 to read in data")
				time.Sleep(3 * time.Second)
			}
			clearScreen()

		case "15": // write out data to the file
			if (len(allusers.Users) > 0) {
				WriteOutFile(allusers)
			} else {
				fmt.Println("\nERROR: NO DATA found in memory please select Option 10 to read in data")
				time.Sleep(3 * time.Second)
			}
			clearScreen()
			
		case "99":  //exit the kiosk
			fmt.Println("\n")
			os.Exit(0) 

		default:  //throw up warning if switch falls through
			fmt.Printf("\n%v is NOT valid!!!", SelectedItem)
			time.Sleep(1 * time.Second)
			clearScreen()
		}
	}
}