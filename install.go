package main

import (
	"bitbucket.org/kardianos/service"
	"log"
)

// Running and installing as a service
const (
	serviceName = "GP-Janitor"
	displayName = "Group Policy Janitor"
	description = "This is a home-made service that keeps some extra windows " +
		"group policy stuff out of this computer"
)

var serv, err = service.NewService(serviceName, displayName, description)

type operation func() error

func register() {
	log.Printf("Registering as service with service name %q...", serviceName)
	serviceOperation(serv.Install)
	log.Printf("NOTE!: To allow changes in your own registry tree, " +
		"you must change the service to be run on your own user " +
		"account either with the SVC UI tool, or with " +
		"'sc config GP-Janitor obj= yourUserName password= yourPassword.")

	start()
}

func start() {
	log.Println("Starting service...")
	serviceOperation(serv.Start)
}

func unregister() {
	log.Println("Removing (unregistering) service...")

	serviceOperation(func() error {
		err = serv.Stop()
		err = serv.Remove()
		return err
	})
}

func serviceOperation(fn operation) {
	// Execute operation
	err := fn()

	if err != nil {
		log.Fatal("Failed: ", err)
	}

	log.Println("...done")
}
