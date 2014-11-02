package dmvDBlib

import (
   _ "github.com/go-sql-driver/mysql"
    "database/sql"
    // "fmt"
    "strings"
    // "os"
)

func GetDB() *sql.DB{
    var dsn string
    if (DEV){
        dsn = DB_USER + ":" + DB_PASS_DEV + "@" + DB_HOST_DEV + "/" + DB_NAME + "?charset=utf8"
    }else{
        dsn = DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
    }
    db, _ := sql.Open("mysql", dsn)

    return db
}

func GetStateList(db *sql.DB) []string {
    sql := "select `name` FROM term_data where vid = 1 ORDER BY tid DESC"
    row, err := db.Query(sql)
    checkErr(err)
    var states []string

    for row.Next() {
        var name string
        row.Scan(&name)
        states = append(states, name)
    }
    return states
}

func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}
