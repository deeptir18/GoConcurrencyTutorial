package main

import "sync"
import "fmt"

func main() {
    counter := 0
    var wg sync.WaitGroup

    for i := 0; i < 10000; i++ {
        wg.Add(1)
        go func(i int) {
            // we're not using a mutex here... what goes wrong?
            counter += 1
            wg.Done()
        }(i)
    }

    wg.Wait()
    fmt.Printf("counter is: %d\n", counter)
}
