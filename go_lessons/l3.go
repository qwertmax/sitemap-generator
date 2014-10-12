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

// func fanIn(input1, input2 <-chan string) <-chan string {
func fanIn(set []<-chan string) <-chan string {
	c := make(chan string)

	for i := range set {
		go func(in <-chan string) {
			for {
				c <- <- in
			}
		}(set[i])
	}

	return c
}

func main() {
	set := []<-chan string{mylib.Boring("Joe"), mylib.Boring("Ann"), mylib.Boring("Max"), mylib.Boring("Nati"), mylib.Boring("Evgen")}
	c := fanIn(set)
	for i := 0; i < 30; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're boring: I'm leaving.")
}
