package runner

import (
	"bufio"
	"os"
	"strings"
)

// Function to read multiple domains
func parseDomain(domain string) []string {
	if strings.Contains(domain, ",") {
		return strings.Split(domain, ",")
	}
	return []string{domain}
}

// Function to read domains from a file
func parseFile(fn string) ([]string, error) {
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var domains []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domains = append(domains, strings.TrimSpace(scanner.Text()))
	}

	return domains, nil
}