package main

import (
   _ "github.com/go-sql-driver/mysql"
	"database/sql"
    "time"
    "strings"
    "os"
    "strconv"
    "fmt"
)

const (
    DB_HOST = "tcp(127.0.0.1:8889)"
    // DB_HOST = "tcp(127.0.0.1:3306)"
    DB_NAME = "dmv"
    DB_USER = "root"
    // DB_PASS = "" /*"1"*/
    DB_PASS = "1"
)
const (
    stdLongYear  = "2006"
    stdZeroMonth = "01"
    stdZeroDay   = "02"
)

func main() {
    p := fmt.Println
    dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
    db, err := sql.Open("mysql", dsn)

    sql := "SELECT n.nid, n.title, n.created, a.dst path, u.`name`, td.name genre FROM node n, url_alias a, users u, term_node tn, term_data td WHERE n.type ='news' AND CONCAT('node/', n.nid) = a.src AND n.nid = tn.nid AND td.tid = tn.tid AND td.vid = 10 AND u.uid = n.uid"
    rows, err := db.Query(sql)
    xml := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\" xmlns:news=\"http://www.google.com/schemas/sitemap-news/0.9\">"
    domain := "www.dmv.com"

    checkErr(err)
    for rows.Next() {
        err = rows.Scan(&nid, &title, &created, &path, &name, &genre)
    }

    xml += "</urlset>"

    db.Close()
    // f, err := os.Create("news.xml")
    // byteArray := []byte(xml)
    // f.Write(byteArray)
    // p(xml)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

