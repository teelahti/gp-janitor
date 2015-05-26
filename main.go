package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/sys/windows/svc"
)

const (
	serviceName = "GP-Janitor"
	displayName = "Group Policy Janitor"
	description = "This is a home-made service that keeps some extra windows " +
		"group policy stuff out of this computer"
)

func main() {
	// check, whether this is an interactive command line session
	isIntSess, err := svc.IsAnInteractiveSession()
	if err != nil {
		log.Fatal("failed to determine if we are running in an interactive session: ", err)
	}

	if !isIntSess {
		runService(serviceName, false)
		return
	}

	if len(os.Args) < 2 {
		usage("no command specified")
	}

	cmd := strings.ToLower(os.Args[1])
	switch cmd {
	case "debug":
		runService(serviceName, true)
		return
	case "install":
		err = installService(serviceName, displayName, description)
		if err == nil {
			// TODO: Give username and password	as args and remove this message
			log.Print("Installed as a Windows service")
			log.Printf("NOTE!: To allow changes in your own registry tree, " +
				"you must change the service to be run on your own user " +
				"account either with the SVC UI tool, or with " +
				"'sc config GP-Janitor obj= yourUserName password= yourPassword.")

			// Start newly installed service
			err = startService(serviceName)
		}
	case "remove":
		err = removeService(serviceName)
	case "start":
		err = startService(serviceName)
	case "stop":
		err = controlService(serviceName, svc.Stop, svc.Stopped)
	case "pause":
		err = controlService(serviceName, svc.Pause, svc.Paused)
	case "continue":
		err = controlService(serviceName, svc.Continue, svc.Running)
	default:
		usage(fmt.Sprintf("invalid command %s", cmd))
	}
	if err != nil {
		log.Fatalf("failed to %s %s: %v", cmd, serviceName, err)
	}
	return
}
