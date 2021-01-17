package main

import (
	service_example "github.com/goro9/korat/examples/service"
)

func main() {
	// TODO: set from command line option by cobra
	adapterID := "hci0"
	args := []string{"server", ""}
	service_example.Run(adapterID, args[0], args[1])
}
