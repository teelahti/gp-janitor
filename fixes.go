package main

import (
	"golang.org/x/sys/windows/registry"
)

func timerBasedFixes() {
	// Add other time interval based fixes here

	go keepRegistryValue(
		"Fix IE home page",
		registry.CURRENT_USER,
		`Software\Microsoft\Internet Explorer\Main`,
		"Start Page",
		"about:Tabs")

	go keepRegistryValue(
		"Disable WSUS",
		registry.LOCAL_MACHINE,
		`SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate\AU`,
		"UseWUServer",
		"0")
}
