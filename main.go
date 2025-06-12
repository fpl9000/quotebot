package main

import (
	"bufio"
	"embed"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Embed the quotes.txt file into the binary at compile time.
// This directive tells Go to include the quotes.txt file in the executable.
//
//go:embed quotes.txt
var embeddedQuotes embed.FS

// loadEmbeddedQuotes reads the embedded quotes file and returns them as a slice.
func loadEmbeddedQuotes() ([]string, error) {
	// Read the embedded file content.
	content, err := embeddedQuotes.ReadFile("quotes.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded quotes: %w", err)
	}

	// Split the content into lines and filter out empty lines.
	lines := strings.Split(string(content), "\n")
	var quotes []string
	for _, line := range lines {
		// Trim whitespace and only add non-empty lines.
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			quotes = append(quotes, line)
		}
	}

	return quotes, nil
}

// readQuotesFromFile reads quotes from the specified file and returns them as a slice.
func readQuotesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var quotes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Only add non-empty lines to avoid selecting blank lines.
		if len(line) > 0 {
			quotes = append(quotes, line)
		}
	}

	// Check for scanning errors.
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return quotes, nil
}

// showHelp displays usage information and exits.
func showHelp() {
	fmt.Println("quotebot - A random quote generator")
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println("  quotebot [OPTIONS] [FILE]")
	fmt.Println()
	fmt.Println("DESCRIPTION:")
	fmt.Println("  Displays a randomly selected quote. By default, uses built-in quotes")
	fmt.Println("  embedded in the executable. If a file is specified, quotes are read")
	fmt.Println("  from that file instead. Each quote in the file should be on a separate line.")
	fmt.Println()
	fmt.Println("OPTIONS:")
	fmt.Println("  -h, --help    Show this help message and exit")
	fmt.Println()
	fmt.Println("ARGUMENTS:")
	fmt.Println("  FILE          Optional path to a file containing quotes")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  quotebot                 # Use embedded quotes")
	fmt.Println("  quotebot quotes.txt      # Use quotes from quotes.txt")
	fmt.Println("  quotebot -h              # Show this help")
	os.Exit(0)
}

func main() {
	var quotes []string
	var err error

	// Check for help flags first.
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		showHelp()
	}

	// Check if a file path was provided as a command-line argument.
	if len(os.Args) > 1 {
		// Use quotes from the specified file.
		filename := os.Args[1]
		quotes, err = readQuotesFromFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading quotes from %s: %v\n", filename, err)
			os.Exit(1)
		}
	} else {
		// Use embedded quotes.
		quotes, err = loadEmbeddedQuotes()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading embedded quotes: %v\n", err)
			os.Exit(1)
		}
	}

	// Ensure we have at least one quote to select from.
	if len(quotes) == 0 {
		fmt.Fprintf(os.Stderr, "No quotes available\n")
		os.Exit(1)
	}

	// Select a random quote and print it to standard output.
	randomIndex := rand.Intn(len(quotes))
	fmt.Println(quotes[randomIndex])
}
