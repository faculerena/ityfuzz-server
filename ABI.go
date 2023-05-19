package main

import (
	"encoding/json"
	"log"
	"os"
)

func abiRelations() {

	dir, err := os.ReadDir("./ABI/artifacts")
	if err != nil {
		log.Fatal(err)
	}

	for file := range dir {
		data, err := os.ReadFile("./ABI/artifacts/" + dir[file].Name())
		if err != nil {
			log.Fatal(err)
		}

		var artifact map[string]interface{}
		err = json.Unmarshal(data, &artifact)
		if err != nil {
			log.Fatal(err)
		}

		abi, err := json.Marshal(artifact["abi"])
		if err != nil {
			log.Fatal(err)
		}

		responseBase.Result = string(abi)

		wr, err := os.Create("./ABI/" + dir[file].Name())

		fr := prettyWrite(responseBase)

		wr.Write([]byte(fr))

	}

}

func prettyWrite(p Response) string {
	prettyJSON, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}

	return string(prettyJSON)
}
