package main

import (
	"fmt"
	"sync"
	"time"
)

func testWaitGroups(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
	    testWaitGroups("go is hard myaaannnn!")
	    wg.Done()
	}()

  	go func() {
    	     for i:=0;i<10;i++ {
	        time.Sleep(time.Millisecond * 100)
    	     }
    	     testWaitGroups("concurrency != parallelism")
             wg.Done()
	}()
	wg.Wait()
}
