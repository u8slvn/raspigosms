package main

import (
	"flag"

	"github.com/u8slvn/raspigosms/app"
	"github.com/u8slvn/raspigosms/web"
)

var (
	HTTPAddr = flag.String("http", "127.0.0.1:8000", "Address to listen for HTTP requests on")
)

func main() {
	// Parse the command-line flags.
	flag.Parse()

	httpAddr := *HTTPAddr

	// Start the dispatcher.
	app.StartDispatcher()

	web.StartServer(httpAddr)
}