package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		fmt.Println(resp.Status)
		resp.Body.Close()
	}
}
