package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func relations() abiList {
	file, err := os.ReadFile("relations.json")
	if err != nil {
		log.Fatalf("Error happened in reading file. Err: %s", err)
	}

	data := make(map[string]string)

	// Unmarshal the JSON into our map
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}

	// This will hold our abiList values
	var list abiList
	for address, abiFile := range data {
		list = append(list, addressAndABI{addr: Address(strings.ToLower(address)), abi: abiFile})
	}

	return list
}
