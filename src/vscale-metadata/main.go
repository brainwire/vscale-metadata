package main

import (
	"flag"
	"os"

	"vscale-metadata/server"
)

var (
	port             int
	host             string
	apikey           string
)

const (
	defaultPort             = 8080
	defaultHost             = "0.0.0.0"
	defaultApiKey						= "99136b9b5895526304be641ef859760609d6ef49e2846c6079370696637d5fac"
)

func main() {
	flag.IntVar(&port, "p", defaultPort, "server port")
	flag.StringVar(&host, "h", defaultHost, "server listen addresse")
	flag.StringVar(&apikey, "k", defaultApiKey, "vscale api token")

	flag.Parse()
	os.Exit(run())
}

func run() int {
	config := &server.Config{
		Port:          port,
		Host:          host,
		ApiKey:				 apikey,
	}

	srv := server.NewServer(config)
	if err := srv.Serve(); err != nil {
		return 1
	}
	return 0
}
