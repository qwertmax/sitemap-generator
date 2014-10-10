package main

import (
    "fmt"
    _ "math/rand"
    "time"
)

var p = fmt.Println
var s = fmt.Sprintf

func boring(msg string) <-chan string { // Return receive-only channel of strings
    c := make(chan string)
    go func() { // We launch goroutine from inside the function
        for i := 0; ; i++ {
            c <- fmt.Sprintf("%s %d", msg, i)
            // time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
            if(msg == "Joe"){
            	time.Sleep(1 * time.Second)
            }else{
            	time.Sleep(3 * time.Second)
            }
            // time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
        }
    }()
    return c
}

// c := boring("boring!") // Function returning a channel.
// for i := 0; i < 5; i++ {
//     fmt.Printf("You say: %q\n", <-c)
// }
// fmt.Println("You're boring: I'm leaving.")

// More instances
func main() {
    joe := boring("Joe")
    ann := boring("Ann")
    
    for i := 0; i < 10; i++ {
        fmt.Println(<-joe)
        fmt.Println(<-ann)
    }
    fmt.Println("You're boring: I'm leaving.")
}