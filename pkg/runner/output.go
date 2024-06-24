package runner

import (
	"fmt"
	"os"
)

// Function for saving the subdomains in a file
func SaveOutput(fn string, subdomains []string) error {
	file, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	for _, subdomain := range subdomains {
		file.WriteString(subdomain + "\n")
	}

	// Logging Message
	fmt.Fprintln(os.Stderr, Succfix+"Output saved to:", g.Yellow(fn))

	return nil
}