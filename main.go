package main

import (
	"fmt"

	"nexflare.com/crypto/api"
)

func main() {
	fmt.Println("Hello World")
	res, err := api.GetRate("BTC")
	if err!= nil {
		fmt.Print(err)
	} else {
		fmt.Println(*res)
	}
}