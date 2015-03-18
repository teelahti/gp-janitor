package main

import (
	"flag"
	"fmt"
	"os"
)

// Command line options
var flags = struct {
	help        *bool
	register    *bool
	unRegister  *bool
	interactive *bool
}{
	flag.Bool("?", false, "Shows this help"),
	flag.Bool("register", false, "When set, registers this program as a Windows service"),
	flag.Bool("unregister", false, "When set, removes the Windows service registration"),
	flag.Bool("interactive", false, "When set runs in command line mode (does not expect run as service)"),
}

func specialFlags() {
	flag.Parse()

	if *flags.help {
		printHelps()
		os.Exit(0)
	}

	if *flags.register {
		register()
		os.Exit(0)
	}

	if *flags.unRegister {
		unregister()
		os.Exit(0)
	}
}

func printHelps() {
	fmt.Println("gp-janitor - Removes Windows Group policy cruft from this computer")
	fmt.Println("Usage: ")
	flag.PrintDefaults()
}
