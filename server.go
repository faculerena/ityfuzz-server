package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var list abiList

var responseBase = Response{
	Status:  "1",
	Message: "OK",
}

func main() {

	list = relations()
	abiRelations()
	fmt.Println("Relations and ABI files generated. Starting server.")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/abi/", getAbi)

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func getAbi(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	addr := Address(params.Get("address"))

	fmt.Printf("Received request for the ABI of contract %s\n", addr)

	resp := address2ABI(addr)

	if resp == "" {
		fmt.Println("ERROR HERE")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No ABI found for address %s\n", addr)
		return
	}

	fmt.Printf("Succesfuly returned ABI for address %s, sending back to ityfuzz.\n", addr)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, resp)
}

func address2ABI(addr Address) string {
	for _, v := range list {
		if v.addr == addr {
			return seekABIFile(v.abi)
		}
	}
	return ""

}

func seekABIFile(file string) string {

	path := "./ABI/" + file
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error happened in reading file. Err: %s", err)
	}
	var objMap Response
	err = json.Unmarshal(data, &objMap)
	if err != nil {
		fmt.Println("Error while decoding JSON", err)
		return ""
	}

	// Marshal it back to a string
	b, err := json.Marshal(objMap)
	if err != nil {
		fmt.Println("Error while encoding JSON", err)
		return ""
	}

	return string(b)
}
