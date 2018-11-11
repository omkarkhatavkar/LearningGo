package main

import (
	"fmt"
	"net/http"
)

func getCtype(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "Error", err
	}
	defer resp.Body.Close()
	if resp.Header.Get("Content-Type") == "" {
		return "Error", fmt.Errorf("\nError The content-type is Empty")
	}
	return resp.Header.Get("Content-Type"), nil

}

func main() {
	fmt.Println(getCtype("https://www.facebook.com"))
}
