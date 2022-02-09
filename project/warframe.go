package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)


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

// Read in json file and display all data within that file - no return for now
func DisplayAllData() {

}

// Main kiosk processing loop
func main(){

	var SelectedItem string

	// Infinite loop for kiosk app main menu
	for {
		clearScreen()
		SelectedItem = DisplayMainMenu()

		switch SelectedItem {
		case "1":
			fmt.Printf("\n%v is not implemented yet......", SelectedItem)
			time.Sleep(1 * time.Second)
			clearScreen()

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