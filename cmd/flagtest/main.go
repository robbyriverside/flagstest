package main

import (
	"log"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Verbose bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
	Server  string `short:"s" long:"server" description:"Local archive server URL" default:""`
	Docker  bool   `short:"d" long:"docker" description:"Running in docker"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Verbose: %v", opts.Verbose)
}
