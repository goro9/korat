package server

import (
	"os"

	"github.com/muka/go-bluetooth/hw"
	log "github.com/sirupsen/logrus"
)

func Run(adapterID string) error {

	log.SetLevel(log.TraceLevel)

	btmgmt := hw.NewBtMgmt(adapterID)
	if len(os.Getenv("DOCKER")) > 0 {
		btmgmt.BinPath = "./bin/docker-btmgmt"
	}

	// set LE mode
	btmgmt.SetPowered(false)
	btmgmt.SetLe(true)
	btmgmt.SetBredr(false)
	btmgmt.SetPowered(true)

	uuid := "b6052f67-891b-41f8-9517-afcb066c7955"
	return serve(adapterID, uuid)
}
