package main

import (
	"fmt"
	"os"
	"time"

	"github.com/iaakanshff/crtfinder/pkg/runner"
)

func main() {

	// Parsing the command line arguments
	options, err := runner.ParseOptions()
	if err != nil {
		return
	}

	// Printing the banner here
	runner.Banner()

	// Processing each domain given
	for _, domain := range options.Domains {

		// Getting the current time
		currTime := time.Now()

		// Here we process the domain
		subdomains, err := runner.ProcessDomain(domain, &options.Recursive)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		// Compute the duration
		duration := time.Since(currTime)

		// Extract seconds and milliseconds
		seconds := int(duration.Seconds())
		milliseconds := int(duration.Milliseconds()) % 1000

		// If output file is given we make the file and exit
		if options.Output != "" {
			if err = runner.SaveOutput(options.Output, subdomains); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		} else {
			// Otherwise we print the found subdomains
			for _, subdomain := range subdomains {
				fmt.Println(subdomain)
			}
		}
		fmt.Fprintf(os.Stderr, runner.Succfix+"Found %d subdomains for %s in %d seconds %d milliseconds\n", len(subdomains), domain, seconds, milliseconds)
	}
}