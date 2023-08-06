package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Arpit-Mohapatra/cyoa"
)

func main() {
	filename := flag.String("file", "../../gopher.json", "JSON file with the CYOA story")
	port := flag.Int("port", 3000, "port to start CYOA application on")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting serer at port:%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}