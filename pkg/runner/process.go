package runner

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

// Struct to hold the specific json fields & values when unmarshaling
type ResponseItem struct {
	CommonName string `json:"common_name"`
	NameValue  string `json:"name_value"`
}

// Function for making the request and fetching the response items
func makeRequest(domain string) ([]ResponseItem, error) {

	// // Set up the HTTP client with a custom Transport that skips TLS verification
	client := &http.Client{}

	// Target Url ( crt.sh )
	crtSite := fmt.Sprintf("https://crt.sh?q=%%25.%s&output=json", domain)

	// Preparing the http request
	req, err := http.NewRequest("GET", crtSite, nil)
	if err != nil {
		return nil, fmt.Errorf(Errfix + "Failed to prepare the request")
	}
	req.Header.Set("User-Agent", "Go-http-client/1.1")

	// Making the request and fetching the response
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(Errfix + "Failed to make http request")
	}
	defer resp.Body.Close()

	// Reading the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(Errfix + "Failed to read respone body")
	}

	// Slice to hold the json fields values
	var responseItems []ResponseItem
	if err := json.Unmarshal(body, &responseItems); err != nil {
		if strings.Contains(string(body), "429") {
			return nil, fmt.Errorf(Errfix + "Please wait, Too many requests")
		}
		return nil, fmt.Errorf(Errfix + "Failed to marshal response body")
	}

	return responseItems, nil
}

// Function for processing the subdomains
func ProcessDomain(domain string, options *Options) ([]string, error) {

	// Logging Message
	fmt.Fprintln(os.Stderr, Succfix+"Enumerating subdomains for", "\""+g.Bold(g.Yellow(domain))+"\"")

	// Function for making the request
	responseItems, err := makeRequest(domain)
	if err != nil {
		return nil, err
	}

	// Extracting the subdomains using the regex from the body
	subdomainSet := extractsubDomains(responseItems, domain)

	// If recursive flag is set
	if options.Recursive {
		fmt.Fprintln(os.Stderr, Succfix+"Recursive approach:", g.BrGreen("ON"), "("+g.BrRed(fmt.Sprintf("Delay: %v", options.Delay))+")")
		var wg sync.WaitGroup
		subdomainChan := make(chan string)

		for subdomain := range subdomainSet {
			// We will only go recursive on wildcard subdomains
			if strings.HasPrefix(subdomain, "*.") {
				wg.Add(1)
				go func(subdomain string) {
					defer wg.Done()

					// Make request and fetch the response items
					responseItems, err := makeRequest(subdomain)
					if err != nil {
						fmt.Fprintf(os.Stderr, Errfix+"Error fetching subdomains for %s:\n%v\n", subdomain, err)
						return
					}

					// Extract the found subdomains from the response items
					recvSubdomainSet := extractsubDomains(responseItems, domain)
					for subdomain := range recvSubdomainSet {
						subdomainChan <- subdomain
					}
				}(strings.TrimPrefix(subdomain, "*."))
				time.Sleep(time.Duration(options.Delay) * time.Second)
			}
		}

		go func() {
			wg.Wait()
			close(subdomainChan)
		}()

		// Iterating over the subdomains received from chan
		for subdomain := range subdomainChan {
			subdomainSet[subdomain] = struct{}{}
		}

	}

	// Deduplication process
	// Append all the unique subdomains into the slice
	subdomainMap := make(map[string]struct{})
	for subdomain := range subdomainSet {
		subdomainMap[strings.TrimPrefix(subdomain, "*.")] = struct{}{}
	}

	// Append all the unique subdomains into the slice
	var subdomains []string
	for subdomain := range subdomainMap {
		subdomains = append(subdomains, subdomain)
	}

	// Sort the slice`
	sort.Strings(subdomains)

	// Returning the subdomains
	return subdomains, nil
}

// Function for extracting the subdomains from the given body
func extractsubDomains(responseItems []ResponseItem, domain string) map[string]struct{} {

	// A map to hold the found subs & de duplication purposes
	subdomainSet := make(map[string]struct{})
	subdomainRegex := regexp.MustCompile(`(?:\*\.)?(?:[a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}`)

	for _, item := range responseItems {
		names := append(strings.Split(item.CommonName, "\n"), strings.Split(item.NameValue, "\n")...)
		for _, name := range names {
			name = strings.TrimSpace(name)
			matches := subdomainRegex.FindAllString(name, -1)

			for _, match := range matches {
				if match != "" && strings.HasSuffix(match, "."+domain) {
					subdomainSet[strings.ToLower(match)] = struct{}{}
				}
			}
		}
	}
	return subdomainSet
}
