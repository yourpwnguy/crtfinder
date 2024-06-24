package runner

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/iaakanshff/gostyle"
)

// Struct to hold our cli arguments
type Options struct {
	Domains   []string // Domains list
	File      string   // Input file
	Recursive int     // Recursive approach
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
	flag.IntVar(&options.Recursive, "r", 5, "")
	flag.BoolVar(&options.Version, "v", false, "")

	// Customize usage message
	flag.Usage = func() {
		h := "\nUsage: crtfinder [options]\n\n"
		h += "Options: [flag] [argument] [Description]\n\n"
		h += "  -d string[]\tDomains to find subdomains for ( can be comma separated )\n"
		h += "  -dL FILE\tInput file containing a list of domains\n"
		h += "  -r int\tFor recursively finding subdomains (default time gap between requests: 5s)\n"
		h += "  -o string\tOutput file to store the subdomains\n"
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

	// Returning the options struct
	return options, nil
}