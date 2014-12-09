package main

import (
	"log"
	"time"
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
	// TODO: Add one time run & filesystem watcher based fixes when needed
	// use e.g. github.com/go-fsnotify/fsnotify as watcher

	ticker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-ticker.C:
			// Add other time interval based fixes here
			go fixIEConfiguration()

		case <-exit:
			log.Println("Stopping the service")
			ticker.Stop()
			return
		}
	}
}
