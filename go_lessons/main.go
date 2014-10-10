package main

import (
    "fmt"
    "math/rand"
    "time"
)

var p = fmt.Println
var s = fmt.Sprintf

func main() {
    c := make(chan string)
    // p(<- c)
    go boring("hello", c)

    for i := 0; i < 5; i++ {
        p("You say:", <-c)
    }

    p("you are boring I'm leaving....")
    time.Sleep(2 * time.Second)

}

func boring(msg string, c chan string){
    for i := 0; ; i++ {
        c <- s("%s %d", msg, i)
        time.Sleep(time.Duration(rand.Intn(10) * 100) * time.Millisecond)
    }
}
