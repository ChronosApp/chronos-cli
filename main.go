package main

import (
	"fmt"
	"github.com/levigross/grequests"
	"github.com/mgutz/ansi"
	"log"
	"os"
	"time"
)

type Message struct {
	content string
	ts      time.Time
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	remote, err := ParseConfig()
	if err != nil {
		log.Fatal(ansi.Color("Problem reading config file: "+err.Error(), "red+b"))
	}

	ro := grequests.RequestOptions{
		JSON: fmt.Sprintf(`{"ts", "%d", "content": "%s" }`, time.Now().Unix(), os.Args[1]),
	}
	resp, err := grequests.Post(remote.URL, &ro)
	if err != nil {
		log.Fatal(ansi.Color("Remote URL is unreachable: "+err.Error(), "red+b"))
	}
	if resp.StatusCode != 200 {
		log.Fatal(ansi.Color("Something went wrong: ", "red+b"))
	}

	log.Printf(ansi.Color("Ok", "green+b"))
}
