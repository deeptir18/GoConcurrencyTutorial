package main

import (
    "fmt"
    "time"
)

func main() {
    // What happens if I change this to a buffered channel?
    ch := make(chan int, 0)

    // Create a goroutine that sends out 5 pieces of data
    go func(ch chan int) {
        fmt.Println("Func goroutine begins sending data")
        for i := 1; i <= 5; i++ {
            ch <- i
            fmt.Println("Func goroutine sends data: ", i)
        }
        fmt.Println("Func goroutine ends sending data")
        close(ch)
    }(ch)

    // Wait for two seconds
    fmt.Println("Main goroutine sleeps 2 seconds")
    time.Sleep(time.Second * 2)

    // Receive data
    fmt.Println("Main goroutine begins receiving data")
    for d := range ch {
        fmt.Println("Main goroutine received data:", d)
    }
}