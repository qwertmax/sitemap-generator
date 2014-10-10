package main

import (
	"./mylib"
	"fmt"
	_ "math/rand"
	_ "reflect"
	_ "time"
)

var p = fmt.Println
var s = fmt.Sprintf

func fanIn(input1, input2 <-chan string) <-chan string {
    c := make(chan string)

    go func() { for { c <- <-input1 } }()
    go func() { for { c <- <-input2 } }()

    return c
}

func main(){
	c := fanIn(mylib.Boring("Joe"), mylib.Boring("Ann"))
	for i := 0; i < 10; i++ {
	    fmt.Println(<-c)
	}
	fmt.Println("You're boring: I'm leaving.")
}
