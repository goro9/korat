package main

import (
	server "github.com/goro9/korat/pkg/server"
)

func main() {
	// TODO: set from command line option by cobra
	adapterID := "hci0"
	server.Run(adapterID)
}
