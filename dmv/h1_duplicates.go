package main

import (
   _ "github.com/go-sql-driver/mysql"
    "database/sql"
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
    db := getDB()

    titles := getTitles(db)
    var output string

    for key, val := range titles {
        if ( val > 1) {
            if(CONSOLE){
                p("Key:", key, "Val: ", val)
            }else{
                output += s("%s:%d\n", key, val)
            }
            
        }
    }

    if(!CONSOLE){
        saveFile("h1_duplicates.csv", output)
    }

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

func saveFile(fname, fileContent string){
    f, err := os.Create(fname)
    checkErr(err)
    byteArray := []byte(fileContent)
    f.Write(byteArray)
}
