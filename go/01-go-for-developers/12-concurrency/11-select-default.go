package main

import (
	"fmt"
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot(logChan)
		case <-saveAfter:
			saveSnapshot(logChan)
		default:
			waitForData(logChan)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}


func main() {
	snapshotInterval := 2 * time.Second
	saveAfterDuration := 10 * time.Second

	snapshotTicker := time.NewTicker(snapshotInterval).C
	saveAfter := time.NewTimer(saveAfterDuration).C

	logChan := make(chan string)

	go saveBackups(snapshotTicker, saveAfter, logChan)

	for log := range logChan {
		fmt.Println(log)
	}

	fmt.Println("Main finished")
}