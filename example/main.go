package main

import (
	"encoding/json"
	"fmt"

	"github.com/jonagold-lab/go-apple/apple"
)

func main() {
	resp, err := apple.GetAppByID(1222530780)
	if err != nil {
		panic(err)
	}
	res, _ := json.Marshal(&resp)
	fmt.Println(string(res))
}
