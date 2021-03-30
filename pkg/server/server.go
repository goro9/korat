package server

import (
	"time"

	"github.com/muka/go-bluetooth/api/service"
	"github.com/muka/go-bluetooth/bluez/profile/agent"
	"github.com/muka/go-bluetooth/bluez/profile/gatt"
	log "github.com/sirupsen/logrus"
)

type SetOptions func(*service.App) error

var setOptions SetOptions

func AddSetter(s SetOptions) {
	setOptions = s
}

func serve(adapterID, uuid string) error {

	options := service.AppOptions{
		AdapterID:  adapterID,
		AgentCaps:  agent.CapNoInputNoOutput,
		UUIDSuffix: uuid[8:],
		UUID:       uuid[:4],
	}

	app, err := service.NewApp(options)
	if err != nil {
		return err
	}
	defer app.Close()

	adv := app.GetAdvertisement()
	adv.AddServiceUUID(uuid)

	app.SetName("korat_test")

	log.Infof("HW address %s", app.Adapter().Properties.Address)

	if !app.Adapter().Properties.Powered {
		err = app.Adapter().SetPowered(true)
		if err != nil {
			log.Fatalf("Failed to power the adapter: %s", err)
		}
	}

	rpiIotSvc, err := app.NewService("0000")
	if err != nil {
		return err
	}

	err = app.AddService(rpiIotSvc)
	if err != nil {
		return err
	}

	setupChar, err := rpiIotSvc.NewChar("0010")
	if err != nil {
		return err
	}

	// TODO: secure write and read
	setupChar.Properties.Flags = []string{
		gatt.FlagCharacteristicRead,
		gatt.FlagCharacteristicWrite,
	}

	setupChar.OnRead(service.CharReadCallback(func(c *service.Char, options map[string]interface{}) ([]byte, error) {
		log.Warnf("GOT READ REQUEST")
		return []byte{42}, nil
	}))

	setupChar.OnWrite(service.CharWriteCallback(func(c *service.Char, value []byte) ([]byte, error) {
		log.Warnf("GOT WRITE REQUEST")
		log.Warnf(string(value))
		return value, nil
	}))

	err = rpiIotSvc.AddChar(setupChar)
	if err != nil {
		return err
	}

	log.Infof("Exposed service %s", rpiIotSvc.Properties.UUID)

	if setOptions != nil {
		if err := setOptions(app); err != nil {
			return err
		}
	}

	err = app.Run()
	if err != nil {
		return err
	}

	timeout := uint32(6 * 3600) // 6h
	log.Infof("Advertising for %ds...", timeout)
	cancel, err := app.Advertise(timeout)
	if err != nil {
		return err
	}

	defer cancel()

	wait := make(chan bool)
	go func() {
		time.Sleep(time.Duration(timeout) * time.Second)
		wait <- true
	}()

	<-wait

	return nil
}
