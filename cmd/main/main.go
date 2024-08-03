package main

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"log"
	"os"
	"strings"
)

func main() {
	rcKey := "RC"
	rpKey := "RP"

	rc := flag.String("rc", "", fmt.Sprintf("Replacement string for '%s'", rcKey))
	rp := flag.String("rp", "", fmt.Sprintf("Replacement string for '%s'", rpKey))
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatalln("Input file name is the required argument")
	}

	if *rc == "" || *rp == "" {
		log.Fatalln("-rc and -rp flags are required")
	}

	inputFile := flag.Arg(0)

	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	text := string(content)
	text = strings.ReplaceAll(text, rcKey, *rc)
	text = strings.ReplaceAll(text, rpKey, *rp)

	err = clipboard.WriteAll(text)
	if err != nil {
		log.Fatalf("Error writing to clipboard: %v", err)
	}

	log.Printf("Replacements made and result copied to clipboard")
}
