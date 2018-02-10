package main

import (
    "fmt"
    "time"
)

func main() {
    values := [5]int{10, 20, 30, 40, 50}

    fmt.Println("Don't do this:")

    // Bad example - this isn't going to work
    for _, val := range values {
        go func() {
            fmt.Println(val)
        }()
    }

    time.Sleep(time.Second)
    fmt.Println("\nDo do this:")

    // Good example - this is going to work!
    for _, val := range values {
        go func(val int) {
            fmt.Println(val)
        }(val)
    }

    time.Sleep(time.Second)
    fmt.Println("\nOr you can do this:")

    // Also a good example!
    for i, _ := range values {
        val := values[i]
        go func() {
            fmt.Println(val)
        }()
    }

    // A better way of making sure that all goroutines
    // finish before the main program exits
    // would be to use a channel (but this is simpler)
    time.Sleep(time.Second)
}