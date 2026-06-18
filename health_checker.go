package main

import (
	"fmt"
	"net/http"
	"time"
)

func check(url string) {
	client := http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(url)

	if err != nil {
		fmt.Println(url, "DOWN")
		return
	}

	fmt.Println(url, resp.Status)
}

func main() {
	check("https://api.github.com")
}
