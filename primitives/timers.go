package main

import "time"
import "fmt"

func main() {
	timeChan := time.NewTimer(time.Second).C // timer duration is 1 second

	tickChan := time.NewTicker(time.Millisecond * 400).C // ticker duration is 400 ms

	doneChan := make(chan bool)
	start := time.Now()

	// convenient way to stop in x seconds
	go func() {
		time.Sleep(time.Second * 2)
		doneChan <- true
	}()

	for {
		select {
		case <-timeChan:
			fmt.Println("Timer expired: ", time.Now().Sub(start))
		case <-tickChan:
			fmt.Println("Ticker expired: ", time.Now().Sub(start))
		case <-doneChan:
			fmt.Println("Done")
			return
		}
	}
}
