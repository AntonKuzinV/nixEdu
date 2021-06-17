package main

import (
    "fmt"
    "time"
)

func pinger(c chan string) {
    for i := 0; ; i++ {
        c <- "ping"
    }
}

func ponger(c chan string) {
    for i := 0; ; i++ {
        c <- "pong"
    }
}

func printer(c chan string) {
    for {
        msg := <- c
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
}
func sleep (t int){
<- time.After(time.Second * time.Duration(t))
}

func main() {
    sleep(5)
    var c chan string = make(chan string)

    go pinger(c)
    go ponger(c)
    go printer(c)
	
    var input string
    fmt.Scanln(&input)
}