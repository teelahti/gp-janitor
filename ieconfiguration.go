package main

import (
	"github.com/luisiturrios/gowin"
	"log"
)

const (
	hkey = "HKCU"
	path = "Software\\Microsoft\\Internet Explorer\\Main"
	name = "Start Page"
)

func fixIEConfiguration() {
	const target = "about:Tabs"

	val, err := ieHomepage()

	if err != nil {
		log.Println("Failed to get REG value: ", err)
		return
	}

	if val == target {
		log.Printf("Homepage %q ok, no need to change.", val)
		return
	}

	log.Printf("Invalid homepage %q, changing it to %q... ", val, target)
	setIeHomepage(target)
}

func ieHomepage() (string, error) {
	return gowin.GetReg(hkey, path, name)
}

func setIeHomepage(value string) error {
	err := gowin.WriteStringReg(hkey, path, name, value)

	if err != nil {
		log.Println("Changing the value failed ", err)
	}

	return err
}
