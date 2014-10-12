package main

import (
	"./mylib"
	"fmt"
	_ "math/rand"
	_ "reflect"
	_ "time"
)

func fanIn(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func() { 
        for {
            select {
                case s := <-input1; c <-s 
                case s := <-input2; c <-s 
            }    
        } 
    }()
    return c
}
 
// Timeout using select
func main() {
    c := boring("Joe")
    for {
        select {
            case s := <-c:
                fmt.Println(s)
            case <-time.After(1 * time.Second):
                fmt.Println("You're too slow!")
                return
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
