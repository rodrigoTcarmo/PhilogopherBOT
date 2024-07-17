package main

import (
	"fmt"

	"github.com/rodrigoTcarmo/PhilogopherBOT/pkg/auth"
	"github.com/rodrigoTcarmo/PhilogopherBOT/pkg/hashtag"
)

func main() {
	fmt.Println("Starting PhilogopherBOT hashtag...")
	client, err := auth.Auth()
	if err != nil {
		fmt.Printf("error while trying to auth: %s", err)
	}

	err = hashtag.Hashtag(client)
	if err != nil {
		fmt.Printf("error while trying to call hashtag func: %s", err)
	}
}
