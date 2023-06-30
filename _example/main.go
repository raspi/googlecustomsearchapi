package main

import (
	"fmt"
	"github.com/raspi/googlecustomsearchapi"
	"net/http"
)

func main() {
	httpc := http.DefaultClient

	gs := googlecustomsearchapi.New(httpc, "api-key-here", "engine-id-here")

	custom := map[string]string{
		`siteSearchFilter`: `i`,
		`siteSearch`:       `www.imdb.com`,
	}

	res, err := gs.Search(`the room`, 0, custom)
	if err != nil {
		panic(err)
	}

	for _, r := range res.Items {
		fmt.Printf(`%v %v`+"\n", r.Link, r.Title)
	}

}
