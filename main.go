package main

import (
	"flag"
	"log"

	"github.com/jbic9832/go-calorie-tracker/api"
)

func main() {
	listendAddr := flag.String("listenAddr", ":3000", "server address")
	flag.Parse()

    server := api.NewServer(*listendAddr)
    log.Fatal(server.Start())
}
