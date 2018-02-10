package main

import "fmt"

var done = make(chan bool)
var msgs = make(chan int)

// Spawn two go routines
// When the main thread finishes (i.e. reads from the done channel), all go routines are killed
// Note that this example only has one producer and one consumer, but when there are more than one producer and one consumer, you might want to use a buffered channel so none of the workers block
func main() {
	go produce()
	go consume()
	<-done
}

// Producer: produce messages to send to the consumer function
func produce() {
	for i := 0; i < 10; i++ {
		msgs <- i
	}
	done <- true
}

// Consumer: wait for messages on msg channel
func consume() {
	for {
		msg := <-msgs
		fmt.Println(msg)
	}
}
