package main

import (
	"fmt"
	"github.com/pato/gonovation/launchpad"
    "os/exec"
)

func runUname() {
    // Command to execute
    cmd := exec.Command("uname", "-a")

    // Execute the command and capture output
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Error running the command \"uname -a\":", err)
        return
    }

    // Print the command output
    fmt.Println(string(output))
    fmt.Println("Command \"uname -a\" has been completed without erros")

}

func main() {
    
    // Setup
	launchpad := gonovation.GetLaunchPad()
    launchpad.Reset();
	fmt.Println("Initalized")
	
    // Get events
    events := launchpad.Events()
    
    // Loop and check events
    for {
		event := <-events
		x, y, pressed := gonovation.EventInfo(event)
        
		if pressed {
           
           // Light the pressed key
           fmt.Println("You pressed: %d, %d", x, y)
     	   launchpad.Led(int(x), int(y), 0, 3)
           
        
           button := [2] int { int(x), int(y) }   
           switch button {
           
           // Reset button
           case [2] int { 0, 8 }:
                fmt.Println("Clearing Screen")
                launchpad.Reset()
           case [2] int {0, 7}:
                runUname()
           default:
           }

        } else {

           // Clear all buttons that are not being pressed
           launchpad.Reset()
        }
	}

	defer launchpad.Close()
}

