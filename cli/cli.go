package main

import (
	"encoding/json"
	"fmt"
	"github.com/genstackio/gopc"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Syntax: gopc <publicVendorToken> <privateVendorToken> <env> <action>")
		os.Exit(1)
	}
	publicVendorToken := os.Args[1]
	privateVendorToken := os.Args[2]
	env := os.Args[3]
	action := os.Args[4]
	switch action {
	case "create-request":
		c := gopc.Client{}
		c.Init(publicVendorToken, privateVendorToken, env)
		var data url.Values
		r, err := c.CreateRequest(data)
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
			os.Exit(4)
		}
		fmt.Println(r)
	case "get-b2c-balance":
		c := gopc.Client{}
		c.Init(publicVendorToken, privateVendorToken, env)
		r, err := c.GetB2CBalance()
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
			os.Exit(5)
		}
		j, err := json.Marshal(*r)
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
			os.Exit(6)
		}
		fmt.Println(string(j))
	case "test-credentials":
		ok, err := gopc.TestApiCredentials(publicVendorToken, privateVendorToken, env)

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
