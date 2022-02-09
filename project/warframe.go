package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)


func clearScreen() int {

	// Clear Screen before menu
	clrScreen := exec.Command("clear")
	clrScreen.Stdout = os.Stdout
	clrScreen.Run()
	return 0
}

func DisplayMainMenu() string {

	var menuChoice string

	fmt.Println("\n\nWelcome to the Warframe User Relay Main Menu")
	fmt.Println("--------------------------------------------")
	fmt.Println("1. Display All Current User's Frames")
	fmt.Println("2. Add User frame Data")
	fmt.Println("3. Save all data to JSON file")
	fmt.Println("")
	fmt.Println("9. Exit Menu")
	fmt.Println("--------------------------------------------")
	fmt.Println("\n")

	fmt.Printf("Enter your Choice: ")
	fmt.Scan(&menuChoice)
	return menuChoice
}


func main(){

	var SelectedItem string

	// Infinite loop for kiosk app main menu
	for {
		_ = clearScreen()
		SelectedItem = DisplayMainMenu()
		switch SelectedItem {
		case "9":
			fmt.Println("\n")
			os.Exit(0) 
		default:
			fmt.Printf("\n%v is NOT valid!!!", SelectedItem)
			time.Sleep(1 * time.Second)
			_ = clearScreen()
		}
	}

}