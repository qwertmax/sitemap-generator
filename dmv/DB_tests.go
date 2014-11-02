package main

import (
   _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "dmvDBlib"
    "fmt"
    "strings"
    "os"
)

const DEV = true
const CONSOLE = true
const STATES_INCLUDED = true

const (
    DB_HOST_DEV = "tcp(127.0.0.1:8889)"
    DB_HOST = "tcp(127.0.0.1:3306)"
    DB_NAME = "dmv"
    DB_USER = "root"
    DB_PASS = ""
    DB_PASS_DEV = "1"
)

var p = fmt.Println
var s = fmt.Sprintf

func main() {
    db := dmvDBlib.GetDB()

}

func 

func getTitles(db *sql.DB) map[string]int {
    // sql := "SELECT title FROM node WHERE type IN ('article', 'page') AND title LIKE 'Commercial Driver\\'s Manual in%%'"
    // sql := "SELECT title FROM node WHERE type IN ('article', 'page') ORDER BY nid LIMIT 0,10"
    // sql := "SELECT title FROM node WHERE type IN ('article', 'page') ORDER BY nid"
    // sql := "SELECT title FROM node"
    row, _ := db.Query(sql)
    titles := make(map[string]int)

    states := getStateList(db)

    for row.Next() {
        var title string
        row.Scan(&title)

        if(!STATES_INCLUDED){
            title = removeState(title, states)
        }

        if (titles[title] == 0) {
            titles[title] = 1
        }else{
            titles[title] = titles[title] + 1
        }
    }

    return titles
}

func saveFile(fname, fileContent string){
    f, err := os.Create(fname)
    checkErr(err)
    byteArray := []byte(fileContent)
    f.Write(byteArray)
}
