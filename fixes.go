package main

func timerBasedFixes() {
	// Add other time interval based fixes here
	go keepRegistryString(
		"Fix IE home page",
		RegistryKey{"HKCU", `Software\Microsoft\Internet Explorer\Main`, "Start Page"},
		"about:Tabs")

	go keepRegistryDword(
		"Disable WSUS",
		RegistryKey{"HKLM", `SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate\AU`, "UseWUServer"},
		"")
}
