package main

import (
	"fmt"

	"github.com/prysmaticlabs/prysm/v3/api/client/beacon"
)

func main() {
	host, err := parseHostFromCommandLineArgs("localhost:4000")
	if err != nil {
		panic(err)
	}
	client, err := beacon.NewClient(host)
	if err != nil {
		panic(err)
	}

	fmt.Println("hello world", client)
}

func parseHostFromCommandLineArgs(defaultHost string) (string, error) {
	return defaultHost, nil
}
