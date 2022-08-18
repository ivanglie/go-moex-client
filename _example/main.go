package main

import (
	"fmt"

	moex "github.com/ivanglie/go-moex-client"
)

func main() {
	client := moex.NewClient()
	rate, err := client.GetRate(moex.CNYRUB)
	if err != nil {
		panic(err)
	}
	fmt.Println(rate)
}
