package main

import (
	"fmt"

	"github.com/jorgevgut/goworkshop/cache"
	"github.com/jorgevgut/goworkshop/request"
)

func main() {
	var storeImpl cache.StoreManager   // nolint
	storeImpl = cache.NewSimpleStore() // replace constructor for different impls
	var reqManager request.Manager     // nolint
	reqManager = request.Instance()    // get an initialized instance
	reqManager.UseCache(storeImpl)

	// build http request
	opts := request.Options{
		URL:    "https://api.github.com/users/jorgevgut/orgs",
		Method: "GET",
		Headers: map[string]string{
			"Content-Type": "application/json;",
		},
	}

	responseChan := reqManager.Request(opts) // github

	go func() {
		<-responseChan
		fmt.Println("GITHUB @ REQQ>>>>>>>")
	}()

	fmt.Println(reqManager)

	// block execution preventing program to end
	c := make(chan struct{})
	<-c
}
