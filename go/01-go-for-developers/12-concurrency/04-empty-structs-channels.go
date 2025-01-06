package main

import (
	"fmt"
	"time"
)

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan // This will block until each token is received
	}
}

// don't touch below this line

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}

func main() {
	numDBs := 5
	dbChan, count := getDBsChannel(numDBs)
	waitForDBs(numDBs, dbChan)
	fmt.Printf("All %v databases are online, count=%d\n", numDBs, *count)
	time.Sleep(1 * time.Second)
}
