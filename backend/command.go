package backend

import (
	"errors"
	"flag"
	"fmt"

	"github.com/STNS/STNS/stns"
)

type Backend interface {
	Migrate() error
	Delete() error
	DriverName() string
}

func GetInstance(config *stns.Config) Backend {
	var b Backend = nil
	switch config.Backend.Driver {
	case "mysql":
		b = &Mysql{config}
	}
	return b
}

func SubCommandRun(b Backend) error {
	if len(flag.Args()) > 1 {
		if b == nil {
			return errors.New("unknown backend driver")
		}

		switch flag.Args()[1] {
		case "init":
			if err := b.Migrate(); err != nil {
				return err
			} else {
				fmt.Println("backend driver " + b.DriverName() + " init successful")
				return nil
			}
		case "delete":
			if err := b.Delete(); err != nil {
				return err
			} else {
				fmt.Println("backend driver " + b.DriverName() + " delete successful")
				return nil
			}
		}
	}
	return errors.New(usageTemplate)
}

var usageTemplate = `
Usage:
	stns backend [arguments]

The commands are:
	init	initialize backend
	delete  remove all of the information
`