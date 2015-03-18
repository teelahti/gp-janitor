package main

import (
	"github.com/luisiturrios/gowin"
	"log"
)

type RegistryKey struct {
	hkey, path, name string
}

func keepRegistryString(reg RegistryKey, targetValue string, desc string) {

	val, err := gowin.GetReg(reg.hkey, reg.path, reg.name)

	if err != nil {
		log.Println("Failed to get REG value: ", err)
		return
	}

	if val == targetValue {
		log.Printf("%v: Value %q ok, no need to change.", desc, val)
		return
	}

	log.Printf("%v: Wrong value %q, changing it to %q... ", desc, val, targetValue)
	writeRegistryString(reg, targetValue)
}

func writeRegistryString(reg RegistryKey, value string) error {
	err := gowin.WriteStringReg(reg.hkey, reg.path, reg.name, value)

	if err != nil {
		log.Println("Changing the value failed ", err)
	}

	return err
}
