package main

import (
  "os"
	"fmt"
	"log"

	"npkg.dev/pkl.v1"
)

func main() {
	// Open the PKL file
	file, err := os.Open("config.pkl")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Create a new pickle decoder
	decoder := pkl.NewDecoder(file)

	// Decode the PKL file
	var data interface{}
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatalf("Failed to decode PKL file: %v", err)
	}

	// Print the decoded data
	fmt.Printf("Decoded data: %+v\n", data)
}
