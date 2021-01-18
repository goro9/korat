package server

import (
	"time"

	"github.com/muka/go-bluetooth/api/service"
	"github.com/muka/go-bluetooth/bluez/profile/agent"
	"github.com/muka/go-bluetooth/bluez/profile/gatt"
	log "github.com/sirupsen/logrus"
)

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

	app.SetName("korat_test")

	log.Infof("HW address %s", app.Adapter().Properties.Address)

	if !app.Adapter().Properties.Powered {
		err = app.Adapter().SetPowered(true)
		if err != nil {
			log.Fatalf("Failed to power the adapter: %s", err)
		}
	}

	service1, err := app.NewService("2233")
	if err != nil {
		return err
	}

	err = app.AddService(service1)
	if err != nil {
		return err
	}

	char1, err := service1.NewChar("3344")
	if err != nil {
		return err
	}

	char1.Properties.Flags = []string{
		gatt.FlagCharacteristicRead,
		gatt.FlagCharacteristicWrite,
	}

	char1.OnRead(service.CharReadCallback(func(c *service.Char, options map[string]interface{}) ([]byte, error) {
		log.Warnf("GOT READ REQUEST")
		return []byte{42}, nil
	}))

	char1.OnWrite(service.CharWriteCallback(func(c *service.Char, value []byte) ([]byte, error) {
		log.Warnf("GOT WRITE REQUEST")
		log.Warnf(string(value))
		return value, nil
	}))

	err = service1.AddChar(char1)
	if err != nil {
		return err
	}

	descr1, err := char1.NewDescr("4455")
	if err != nil {
		return err
	}

	descr1.Properties.Flags = []string{
		gatt.FlagDescriptorRead,
		gatt.FlagDescriptorWrite,
	}

	descr1.OnRead(service.DescrReadCallback(func(c *service.Descr, options map[string]interface{}) ([]byte, error) {
		log.Warnf("GOT READ REQUEST")
		return []byte{42}, nil
	}))
	descr1.OnWrite(service.DescrWriteCallback(func(d *service.Descr, value []byte) ([]byte, error) {
		log.Warnf("GOT WRITE REQUEST")
		return value, nil
	}))

	err = char1.AddDescr(descr1)
	if err != nil {
		return err
	}

	err = app.Run()
	if err != nil {
		return err
	}

	log.Infof("Exposed service %s", service1.Properties.UUID)

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
