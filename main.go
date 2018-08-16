package main

import (
	"time"
	"fmt"
)

func main() {
    ticker := time.NewTicker(time.Millisecond * 5000)
    defer ticker.Stop()
    count := 0
    for {
        select {
        case <-ticker.C:
            count++
            doPeriodically()
        }
    }
}

func doPeriodically() {
    fmt.Println("hoge")
}