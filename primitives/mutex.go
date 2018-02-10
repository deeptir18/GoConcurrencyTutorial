package main

import "sync"
import "fmt"

func main() {
    counter := 0
    var mu sync.Mutex
    var wg sync.WaitGroup

    for i := 0; i < 10000; i++ {
        wg.Add(1)
        go func(i int) {
            // what happens if we don't use the mutex?
            mu.Lock()
            defer mu.Unlock()
            counter += 1
            wg.Done()
        }(i)
    }

    wg.Wait()
    fmt.Printf("counter is: %d\n", counter)
}
