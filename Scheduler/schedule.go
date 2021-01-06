// Simple Go routine that runs every 24 hours
package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(86400 * time.Second)
    done := make(chan bool)

    go func() {
      for {
        select {
          case <- done:
            return
          case _= <-ticker.C:
            fmt.Println("Go routine that runs once a day")
        }
      }
    }()
    <-done
}
