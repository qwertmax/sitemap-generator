package main

import (
	"time"
	"fmt"
	// "strconv"
)

const (
    stdLongYear  = "2006"
    stdZeroMonth = "01"
    stdZeroDay   = "02"
)

func main() {
    p := fmt.Println

    // q := time.Unix(1392899576, 0).Format("2006/01/02")
    q := time.Unix(1392619176, 0).Format(stdLongYear +"/"+ stdZeroMonth +"/"+ stdZeroDay)
    p(q)

	// date := 1389199972
	// t := time.Unix(date, 0)
	// p("Time: %v\n", t)
}
