package main

import (
	"fmt"
	"github.com/genstackio/gopc"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Syntax: gopc <apiKey> <env> <action>")
		os.Exit(1)
	}
	apiKey := os.Args[1]
	env := os.Args[2]
	action := os.Args[3]
	switch action {
	case "test-credentials":
		ok, err := gopc.TestApiCredentials(apiKey, env)

		if !ok {
			fmt.Println("NOK - Bad credentials:" + err.Error())
			os.Exit(3)
		}
		fmt.Println("OK")
	default:
		fmt.Println("Error: unrecognized action '" + action + "'")
		os.Exit(2)
	}
}
