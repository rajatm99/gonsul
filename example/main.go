package main

import (
	"fmt"
	"github.com/rajatm99/gonsul/gonsul"
	"net/http"
	"time"
)

func main() {
	x := gonsul.Gonsul{
		Host: "http://localhost:8500/v1/kv",
		Path: "gonsul",
		File: "./example.json",
		HttpClient: http.Client{
			Timeout: time.Second * 30,
		},
	}
	err := x.InitFromFile()
	fmt.Println(err)
}
