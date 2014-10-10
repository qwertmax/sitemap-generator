package mylib

import (
    "fmt"
    "math/rand"
    "time"
)

var p = fmt.Println
var s = fmt.Sprintf

func Boring(msg string) <-chan string { // Return receive-only channel of strings
    c := make(chan string)
    go func() { // We launch goroutine from inside the function
        for i := 0; ; i++ {
            t := time.Duration(rand.Intn(1000)) * time.Millisecond
            time.Sleep(t)
            c <- fmt.Sprintf("%s %d time: %d", msg, i, t/1000000)
            // time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
            // if(msg == "Joe"){
            // 	time.Sleep(1 * time.Second)
            // }else{
            // 	time.Sleep(3 * time.Second)
            // }
            // time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
            // time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
        }
    }()
    return c
}
