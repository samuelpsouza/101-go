package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main(){
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url += "http://" + url
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		dst := os.Stdout
		b, err := io.Copy(dst, resp.Body)
		statusCode := resp.Status
	
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%v\n", statusCode)
		fmt.Printf("%s", b)
	}
}