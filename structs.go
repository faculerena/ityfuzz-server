package main

type Address string

type addressAndABI struct {
	addr Address
	abi  string
}

type abiList []addressAndABI

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}
