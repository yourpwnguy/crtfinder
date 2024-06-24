package runner

import (
	"fmt"
	"os"
)

// Function for printing banner
func Banner() {
	banner := `
             __  _____           __         
  __________/ /_/ __(_)___  ____/ /__  _____
 / ___/ ___/ __/ /_/ / __ \/ __  / _ \/ ___/
/ /__/ /  / /_/ __/ / / / / /_/ /  __/ /    
\___/_/   \__/_/ /_/_/ /_/\__,_/\___/_/     
	
   Developed by github.com/iaakanshff ` + "(" + g.Bold(g.Green("v1.0")) + ")"

	fmt.Fprintln(os.Stderr, banner + "\n")
}