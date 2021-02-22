## Go concurrency
### Dealing with multiple things at the same period of time. (not necessary simultaneous)

### Main keywords
1. goroutine
2. channels
3. select
4. sync package

### Channels

* Channels are piplines that connect concurrent executing code
* Channels are typed
* Have send and receive semantics
* data is copied from and to channels

#### Semantics

    channel <- "mudambi"

    blockeduntilreceived := <- channel
    
### SELECT

Makes channel non-blocking by providing ability to wait on multiple channels at the same time.
