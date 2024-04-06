package main

import (
	"fmt"
	"github.com/pato/gonovation/launchpad"
    "os/exec"
    "os"
)

// ranger
func runRanger() {
    // Command to execute
    cmd := exec.Command("kitty", "ranger")

    // Execute the command and detach
    err := cmd.Start()
    if err != nil {
        fmt.Println("Error running the command \"kitty ranger\":", err)
        return
    }

    fmt.Println("Command \"kitty ranger\" has been completed without erros")

}

// kitty
func runKitty() {
    // Command to execute
    cmd := exec.Command("kitty")

    // Execute the command and detach
    err := cmd.Start()
    if err != nil {
        fmt.Println("Error running the command \"kitty\":", err)
        return
    }

    fmt.Println("Command \"kitty\" has been completed without erros")

}

// vivaldi
func runVivaldi() {
    // Command to execute
    cmd := exec.Command("vivaldi")

    // Execute the command and detach
    err := cmd.Start()
    if err != nil {
        fmt.Println("Error running the command \"vivaldi\":", err)
        return
    }

    fmt.Println("Command \"vivaldi\" has been completed without erros")

}


func resetLights(launchpad *gonovation.Launchpad) {
     
    launchpad.Reset()
    
    // Light avaiable buttons
    r_orange := 3
    g_orange := 3

    // Power off button
    launchpad.Led(0, 8, 3, 0)

    // Terminal applications
    launchpad.Led(0, 7, r_orange, g_orange)
    launchpad.Led(1, 7, r_orange, g_orange)

    // Desktop Applications
    launchpad.Led(0, 6, r_orange, g_orange)

}

func main() {
    
    // Setup
	launchpad := gonovation.GetLaunchPad()
    resetLights(launchpad)
	fmt.Println("Initalized")
	
    // Get events
    events := launchpad.Events()
   
    // Loop and check events
    for {
		event := <-events
		x, y, pressed := gonovation.EventInfo(event)
       
		if pressed {
           
           // Light the pressed key
           fmt.Println("You pressed: ", x, y)
     	   launchpad.Led(int(x), int(y), 0, 3)
           
           button := [2] int { int(x), int(y) }   
           
           switch button {
           
           // Reset button
           case [2] int { 0, 8 }:

                fmt.Println("Clearing Screen")
                launchpad.Reset()
                launchpad.Close()
                os.Exit(0)
           case [2] int { 0, 7 }:

                runRanger()
           case [2] int { 1, 7 }:

                runKitty()
           case [2] int { 0, 6 }:

                runVivaldi()
           default:
           }

        } else {

           // Clear all buttons that are not being pressed
           resetLights(launchpad)
        }
	}

	defer launchpad.Close()
}

