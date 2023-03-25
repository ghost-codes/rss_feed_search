package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/ghost-codes/sample/matchers"
	"github.com/ghost-codes/sample/search"
)

func main() {
	//perform search for specified term
	var searchTerm string
	for {
		fmt.Println("Please enter search term")
		fmt.Scanln(&searchTerm)
		search.Run(searchTerm)
		fmt.Println("==============================================>")
		fmt.Printf("Search complete for searchTerm: %s\n\n", searchTerm)
	}
}

// init functions are called before main function
func int() {
	//change device for logging to std out
	log.SetOutput(os.Stdout)
}
