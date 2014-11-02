package main

import (
	"./mylib"
	"fmt"
	"math/rand"
	_ "reflect"
	"time"
)

func fanIn(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func(){
        for {
            select {
                case s := <-input1: c <-s
                case s := <-input2: c <-s
            }
        }
    }()
    return c
}

func q() <-chan int {
    c := make(chan int)
    go func(){
        t := time.Duration(rand.Intn(2)) * time.Second
        time.Sleep(t)
        c <- 1
        return
    }()

    return c
}

// Timeout using select
func main() {
    c := mylib.Boring("Joe")
    // q := q()
    for {
        select {
            case s := <-c:
                fmt.Println(s)
            case <-time.After(1 * time.Second):
                fmt.Println("You're too slow!")
                return
            // case <-q:
            //     fmt.Println("Max error")
            //     return
        }
    }
}
 
// Timeout for whole conversation using select
// func main() {
//     c := boring("Joe")
//     timeout := time.After(1 * time.Second)
//     for {
//         select {
//             case s := <-c:
//                 fmt.Println(s)
//             case <-timeout:
//                 fmt.Println("You talk too much!")
//                 return
//         }
//     }
// }
