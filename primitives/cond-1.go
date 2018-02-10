package main

import (
    "sync"
    "fmt"
    "time"
    "math/rand"
)

func main() {
    var buf []int
    var mu sync.Mutex
    go producer(&buf, &mu)

    for {
        // we're using up 100% CPU here...
        mu.Lock()
        for _, v := range buf {
            process(v)
        }
        buf = nil // empty slice
        mu.Unlock()
    }
}

func process(v int) {
    fmt.Printf("processing %d\n", v)
}

func producer(buf *[]int, mu *sync.Mutex) {
    for {
        time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
        mu.Lock()
        *buf = append(*buf, rand.Intn(1000))
        mu.Unlock()
    }
}
