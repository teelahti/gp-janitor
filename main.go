package main

import (
	"log"
	"time"
)

const (
	timerIntervalSeconds = 30
)

var exit = make(chan bool)

func main() {
	specialFlags()

	// Option A: Run from command line with interactive switch
	if *flags.interactive {
		doWork()
		return
	}

	// Option B: Run as a windows service
	// Launch the main loop through kardianos/service
	err := serv.Run(startWork, stopWork)

	if err != nil {
		serv.Error(err.Error())
		log.Fatal(err)
	}
}

func startWork() error {
	go doWork()
	return nil
}

func stopWork() error {
	exit <- true
	return nil
}

func doWork() {
	log.Printf("Timer set to %d s", timerIntervalSeconds)

	registerFileSystemWatcherFixes()
	oneTimeFixes()

	// Run timer based fixes once to get immediate feedback
	// on interactive mode.
	timerBasedFixes()

	ticker := time.NewTicker(timerIntervalSeconds * time.Second)

	for {
		select {
		case <-ticker.C:
			timerBasedFixes()

		case <-exit:
			log.Println("Stopping the service")
			ticker.Stop()
			return
		}
	}
}

func registerFileSystemWatcherFixes() {
	// TODO: Add filesystem watcher based fixes when needed
	// use e.g. github.com/go-fsnotify/fsnotify as watcher
}

func oneTimeFixes() {
	// TODO: Add one time (on startup) run fixes here
}

func timerBasedFixes() {
	// Add other time interval based fixes here
	reg := RegistryKey{"HKCU", "Software\\Microsoft\\Internet Explorer\\Main", "Start Page"}
	go keepRegistryString(reg, "about:Tabs", "Fix IE home page")
}
