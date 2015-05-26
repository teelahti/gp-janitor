package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
	"strconv"
)

func keepRegistryValue(desc string, root registry.Key, path string, keyname string, targetValue string) {

	key, err := registry.OpenKey(root, path, registry.QUERY_VALUE|registry.SET_VALUE)

	if err != nil {
		log.Printf("%v: Failed to open key: %v", desc, err)
		return
	}
	defer key.Close()

	val, valtype, err := getStringValue(key, keyname)

	if err != nil {
		log.Printf("%v: Failed to get REG value: %v", desc, err)
		return
	}

	if val == targetValue {
		log.Printf("%v: Value %q ok, no need to change.", desc, val)
		return
	}

	log.Printf("%v: Wrong value %q, changing it to %q... ", desc, val, targetValue)

	switch valtype {
	case registry.SZ, registry.EXPAND_SZ:
		err = key.SetStringValue(keyname, targetValue)
	case registry.DWORD:
		targetValueConverted, _ := strconv.ParseUint(targetValue, 10, 32)
		err = key.SetDWordValue(keyname, uint32(targetValueConverted))
	default:
		log.Fatal("Unsupported registry value type %s", valtype)
	}

	if err != nil {
		log.Printf("%v: Changing the value failed: %v", desc, err)
	}
}

func getStringValue(key registry.Key, name string) (string, uint32, error) {
	_, valtype, err := key.GetValue(name, nil)

	if err != nil {
		return "", registry.NONE, err
	}

	switch valtype {
	case registry.SZ, registry.EXPAND_SZ:
		return key.GetStringValue(name)
	case registry.DWORD:
		val, valtype, err := key.GetIntegerValue(name)
		return strconv.FormatUint(val, 10), valtype, err
	default:
		return "", registry.NONE, fmt.Errorf("Unsupported registry value type %s", valtype)
	}
}
