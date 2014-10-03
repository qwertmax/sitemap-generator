package main

import (
   _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
    "strings"
   _ "reflect"
)

const (
    DB_HOST_DEV = "tcp(127.0.0.1:8889)"
    DB_HOST = "tcp(127.0.0.1:3306)"
    DB_NAME = "dmv"
    DB_USER = "root"
    DB_PASS = ""
    DB_PASS_DEV = "1"
)
const DEV = true
var p = fmt.Println

func main() {
    db := getDB()

    // states := getStateList(db)
    // p(states)

    titles := getTitles(db)

    for key, val := range titles {
        // p("Key:", key, "Val: ", val)
        if ( val > 1) {
            p("Key:", key, "Val: ", val)
        }
    }

    // title := "New York qqq California aaaa Texas Washington DC"
    // title = removeState(title, states)
    // p(titles)

}

func getTitles(db *sql.DB) map[string]int {
    // sql := "SELECT title FROM node WHERE type IN ('article', 'page') AND title LIKE 'Commercial Driver\\'s Manual in%%'"
    // sql := "SELECT title FROM node WHERE type IN ('article', 'page') ORDER BY nid LIMIT 0,10"
    sql := "SELECT title FROM node WHERE type IN ('article', 'page') ORDER BY nid"
    // sql := "SELECT title FROM node"
    row, _ := db.Query(sql)
    titles := make(map[string]int)

    states := getStateList(db)

    for row.Next() {
        var title string
        row.Scan(&title)

        title = removeState(title, states)

        if (titles[title] == 0) {
            titles[title] = 1
        }else{
            titles[title] = titles[title] + 1
        }
    }

    return titles
}

func removeState(title string, states []string) string {
    for _, state := range states {
        title = strings.Replace(title, state, "", -1)
        // p(state)
    }
    
    title = strings.Replace(title, "  ", " ", 1)
    title = strings.Trim(title, " ")

    return title
}

func getDB() *sql.DB{
    var dsn string
    if (DEV){
        dsn = DB_USER + ":" + DB_PASS_DEV + "@" + DB_HOST_DEV + "/" + DB_NAME + "?charset=utf8"
    }else{
        dsn = DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
    }
    db, _ := sql.Open("mysql", dsn)

    return db
}

func getStateList(db *sql.DB) []string {
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

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
