package main

import (
	"fmt"
    "time"
    _ "runtime"
)


func f(left, right chan int) {
    left <- 1+ <- right
} 
 
func main() {
    // runtime.GOMAXPROCS(runtime.NumCPU())
    start := time.Now()

    const n = 100000
    leftmost := make(chan int)
    right := leftmost
    left := leftmost
    
    for i := 0; i < n; i++ {
        go f(left, right)
    }
    
    go func(c chan int) {
        c <- 1
    }(right)
    
    // time.Sleep(time.Second * 3);
    fmt.Println(<-leftmost)
    fmt.Println(time.Since(start))
}
