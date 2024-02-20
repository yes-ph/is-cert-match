package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		certificateFile string
		privateKeyFile  string
	)

	if len(os.Args) >= 2 {
		certificateFile = os.Args[1]
	}

	if len(os.Args) >= 3 {
		privateKeyFile = os.Args[2]
	}

	matched, err := certificateFilesMatch(certificateFile, privateKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	if matched {
		fmt.Printf("\033[1;32m%s\033[0m", "Matched\n")
		return
	}

	fmt.Printf("\033[1;31m%s\033[0m", "Does not match\n")
}
