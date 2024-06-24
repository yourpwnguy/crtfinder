package runner

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/iaakanshff/gostyle"
)

// Struct to hold our cli arguments
type Options struct {
	Domains   []string // Domains list
	File      string   // Input file
	Recursive bool     // Recursive approach
	Delay     int      // Time gap
	Version   bool     // Version info
	Output    string   // Output file
}

// Some colored messages ( mainly prefix )
var (
	g       = gostyle.New()
	Succfix = "[" + g.Blue("INFO") + "] "
	Errfix  = "[" + g.Red("ERR") + "] "
)

// Function for parsing the cli arguments
func ParseOptions() (Options, error) {

	// Instatiating the options struct
	options := Options{}

	// For domain
	domain := flag.String("d", "", "")

	// For other flags
	flag.StringVar(&options.Output, "o", "", "")
	flag.StringVar(&options.File, "dL", "", "")
	flag.BoolVar(&options.Recursive, "r", false, "")
	flag.BoolVar(&options.Version, "v", false, "")

	// Customize usage message
	flag.Usage = func() {
		h := "\nUsage: crtfinder [options]\n\n"
		h += "Options: [flag] [argument] [Description]\n\n"
		h += "INPUT:\n"
		h += "  -d string[]\tDomains to find subdomains for (comma separated)\n"
		h += "  -dL FILE\tInput file containing a list of domains\n\n"
		h += "FEATURES:\n"
		h += "  -r int\tFor recursively finding subdomains with time gap between requests (default: 5s)\n\n"
		h += "OUTPUT:\n"
		h += "  -o string\tOutput file to store the subdomains\n\n"
		h += "DEBUG:\n"
		h += "  -v none\tCheck current version\n"
		fmt.Fprint(flag.CommandLine.Output(), h)
	}
	flag.Parse()

	// if no flag or argument is given return error
	if flag.NArg() == 0 && flag.NFlag() == 0 {
		flag.Usage()
		return options, fmt.Errorf("")
	}

	if options.Version {
		CheckVersion()
		os.Exit(0)
	}

	// Checking if the domain provided is null
	if strings.TrimSpace(*domain) == "" && strings.TrimSpace(options.File) == "" {
		fmt.Fprintln(os.Stderr, Errfix+"Neither domain or input list is provided")
		return options, fmt.Errorf("")

		// Checking if both the domain and Input file is provided
	} else if strings.TrimSpace(*domain) != "" && strings.TrimSpace(options.File) != "" {
		fmt.Fprintln(os.Stderr, Errfix+"Please use only one of the two flags")
		return options, fmt.Errorf("")
	}

	// Handling input from the file or domain list
	if strings.TrimSpace(options.File) != "" {
		var err error
		options.Domains, err = parseFile(strings.TrimSpace(options.File))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return options, fmt.Errorf("")
		}
	} else {
		// Handling the commas, and getting the slice of domains
		options.Domains = parseDomain(strings.TrimSpace(*domain))
	}

	// For dynamically fetching the -r flag and setting it accordingly
	for index, arg := range os.Args {

		// Fetch the flag from cli arguments
		if arg == "-r" {

			// Check if the index+1 doesn't exceed the bounds
			if index+1 < len(os.Args) {

				// Check if the next argument after -r is not null
				if os.Args[index+1] != "" {

					// Convert string to int
					delay, err := strconv.Atoi(os.Args[index+1])
					if err != nil {
						fmt.Println(Errfix + "The given argument is not an integer")
						return options, fmt.Errorf("")
					}

					// Set the recursivity
					options.Recursive = true

					// Check if delay is less than 0
					if delay > 0 {

						// Set the given delay value
						options.Delay = delay
					} else {
						fmt.Println(Errfix + "Delay should be not be less than 0")
						return options, fmt.Errorf("")
					}
				}

				// If no argument is given after -r, set the default value
			} else {

				// Set the recursivity
				options.Recursive = true // If -r is provided without a value, use the default Delay

				// Set the default delay value
				options.Delay = 5
			}
		}
	}

	// Returning the options struct
	return options, nil
}
