module github.com/goro9/korat/cmd/korat

go 1.15

replace github.com/goro9/korat/pkg v0.0.0 => ../../pkg

require (
	github.com/goro9/korat/pkg v0.0.0
	github.com/muka/go-bluetooth v0.0.0-20201211051136-07f31c601d33 // indirect
	github.com/sirupsen/logrus v1.7.0 // indirect
)
