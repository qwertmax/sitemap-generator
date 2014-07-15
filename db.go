package main

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	"fmt"
)

const (
    DB_HOST = "tcp(127.0.0.1:8889)"
    DB_NAME = "dmv"
    DB_USER = /*"root"*/ "root"
    DB_PASS = /*""*/ "1"
)

func main() {
    dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
    db, err := sql.Open("mysql", dsn)

    rows, err := db.Query("SELECT uid, `name` FROM users WHERE uid < 10")
    checkErr(err)
    for rows.Next() {
        var uid int
        var name string
        err = rows.Scan(&uid, &name)
        checkErr(err)
        // fmt.Println(uid, )
        fmt.Println(name)
    }

    db.Close()
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
