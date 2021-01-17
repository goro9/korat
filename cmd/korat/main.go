package main

import (
	service_example "github.com/goro9/korat/examples/service"
)

func main() {
	adapterID := "hci0"
	args := []string{"service", "server"}
	service_example.Run(adapterID, args[0], args[1])
}
