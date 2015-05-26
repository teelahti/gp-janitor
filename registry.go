package main

import (
	// TODO: Use new official package instead https://godoc.org/golang.org/x/sys/windows/registry
	"github.com/luisiturrios/gowin"
	"log"
	"strconv"
)

type RegistryKey struct{ hkey, path, name string }
type RegistryApi func(desc string, reg RegistryKey, targetValue string)
type RegWriter func(hkey, path, name, value string) error

// Define the public API
var keepRegistryString = keepRegistryValueFactory(gowin.WriteStringReg)
var keepRegistryDword = keepRegistryValueFactory(
	func(hkey, path, name, targetValue string) error {
		val, _ := strconv.ParseUint(targetValue, 10, 32)
		return gowin.WriteDwordReg(hkey, path, name, uint32(val))
	})

// Factory function to avoid code duplication
func keepRegistryValueFactory(regWriter RegWriter) RegistryApi {

	return func(desc string, reg RegistryKey, targetValue string) {
		val, err := gowin.GetReg(reg.hkey, reg.path, reg.name)

		if err != nil {
			log.Printf("%v: Failed to get REG value: %v", desc, err)
			return
		}

		if val == targetValue {
			log.Printf("%v: Value %q ok, no need to change.", desc, val)
			return
		}

		log.Printf("%v: Wrong value %q, changing it to %q... ", desc, val, targetValue)

		err = regWriter(reg.hkey, reg.path, reg.name, targetValue)

		if err != nil {
			log.Printf("%v: Changing the value failed: %v", desc, err)
		}
	}
}
