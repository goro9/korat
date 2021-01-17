module github.com/muka/go-bluetooth/_examples

go 1.14

replace github.com/goro9/korat/examples v0.0.0 => ../examples

require (
	github.com/goro9/korat/examples v0.0.0
	github.com/muka/go-bluetooth v0.0.0-20201211051136-07f31c601d33 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v0.0.7
	github.com/spf13/viper v1.6.2
)
