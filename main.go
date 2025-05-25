package main

import (
	"fmt"

	"github.com/cli/go-gh/v2/pkg/api"
)

func main() {
	fmt.Println("aloha world, this is the gh-salute extension!")
	client, err := api.DefaultRESTClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	response := struct {
		Login     string
		HtmlURL   string `json:"html_url"`
		Followers int
	}{}
	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("running as %s\n", response.Login)
	fmt.Printf("profile: %s\n", response.HtmlURL)
	fmt.Printf("#followers: %d\n", response.Followers)
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
