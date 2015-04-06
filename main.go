package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	flagPort = flag.String("port", "8080", "Port to listen on")
	flagHost = flag.String("host", "", "Hostname to listen on")
	flagRoot = flag.String("root", ".", "Directory to serve")
)

func main() {
	flag.Parse()

	address := *flagHost + ":" + *flagPort
	root := http.Dir(*flagRoot)
	log.Printf("Server started from %s at http://%s", *flagRoot, address)
	http.ListenAndServe(address, http.FileServer(root))
}
