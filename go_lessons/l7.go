package main

import (
	"fmt"
    "time"
    "math/rand"
    _ "runtime"
)

var (
    Web = fakeSearch("web")
    Image = fakeSearch("image")
    Video = fakeSearch("video")
)
 
type Result func(result string) string
type Search func(query string) Result
 
func fakeSearch(kind string) Search {
    return func(query string) Result {
        time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
        return Result(fmt.Sprintf("%s result for %q\n", kind, query))
    }
}
 
// Replication in order to avoid timeouts
func First(query string, replicas ...Search) Result {
    c := make(chan Result)
    searchReplica := func(i, int) { c <- replicas[i](query) }
    for i := range replicas {
        go searchReplica(i)
    }
    return <-c
}
 
// Main search function
func Google(query sting) (results []Result) {
    c := make(chan Result)
 
    go func() { c <- First(query, Web1, Web2) }()
    go func() { c <- First(query, Image1, Image2) }()
    go func() { c <- First(query, Video1, Video2) }()
 
    timeout := time.After(80 * time.Millisecond)
 
    for i := 0; i < 3; i++ {
        select {
            case result := <-c:
                results = append(results, result)
            case <-timeout:
                fmt.Println("timed out.")
                return
        }
    }
}

 
func main() {
    // runtime.GOMAXPROCS(runtime.NumCPU())
    start := time.Now()

    Google("test")

    fmt.Println(time.Since(start))
}
