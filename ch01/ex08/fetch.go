package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const HTTP = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, HTTP) {
			url = HTTP + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b := bytes.NewBuffer([]byte(""))
		if _, err := io.Copy(b, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "copy: %v\n", err)
			os.Exit(1)
		}
		resp.Body.Close()
		fmt.Printf("%s", b)
	}
}
