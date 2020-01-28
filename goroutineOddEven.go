package main

import (
	"fmt"
	"time"
)
func printAndReturnNextInt(prev int) int {
    time.Sleep(1 * time.Second) 
    fmt.Println( prev )

return prev +2

}
func main() {
   // channel to send messages
	var odd = make(chan bool) 
	var even = make(chan bool) 
	var data = make(chan int) 

	go func() {
        var i = 0
        for {
            i = printAndReturnNextInt(i)
 if i > 10 { even <- true }
data <-1

        }
    }()
go func() {

        var i = 1
        for {
<- data
            i = printAndReturnNextInt(i)
if i > 10 { odd <- true }


        }
    }()

<- even
<- odd
}
