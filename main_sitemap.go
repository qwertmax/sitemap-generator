package main

import (
   _ "github.com/go-sql-driver/mysql"
	"database/sql"
    "time"
    // "strings"
    "os"
    "strconv"
    "fmt"
)

const (
    DB_HOST = "tcp(127.0.0.1:8889)"
    // DB_HOST = "tcp(127.0.0.1:3306)"
    DB_NAME = "dmv"
    DB_USER = "root"
    // DB_PASS = ""
    DB_PASS = "1"

    stdLongYear  = "2006"
    stdZeroMonth = "01"
    stdZeroDay   = "02"

    priority = 0.5
    itemsize = 50000
)

func main() {
    p := fmt.Println
    dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
    db, err := sql.Open("mysql", dsn)

    types := []string{"article", "blog", "cdl_physical", "dmv_office", "directory_entry", "driving_school", "news", "page", "page_state"}
    p(types)

    for typeID := range types {
        p(types[typeID])
    }

    os.Exit(0)

    sql := "SELECT n.nid, n.title, n.changed, a.dst path, n.type node_type FROM node n, url_alias a WHERE CONCAT('node/', n.nid) = a.src AND n.type IN ("+
        "'article',"+
        "'blog',"+
        "'news',"+
        "'page',"+
        "'page_state',"+
        "'directory_entry',"+
        "'cdl_physical',"+
        "'dmv_office',"+
        "'driving_school'"+
    ")"

    rows, err := db.Query(sql)
    xml := ""
    domain := "www.dmv.com"
    var url, path, title, node_type string
    var nid, changed int
    i := 0

    checkErr(err)
    for rows.Next() {
        err = rows.Scan(&nid, &title, &changed, &path, &node_type)
        url = "http://" + domain +"/"+ path
        xml += "<url>"+
            "<loc>"+ url + "</loc>"+
            "<lastmod>"+ TimeFormat(changed) +"</lastmod>"+
            "<changefreq>"+ getChangefreq(node_type) +"</changefreq>"+
            "<priority>"+ Ftoa(priority) +"</priority>"+
        "</url>"

        i++

        if(i % itemsize == 0){
            saveFileStemap(xmlName(i, itemsize), xmlWrap(xml))
            xml = "";
            p(i)
        }
    }
    db.Close()

    saveFileStemap("main-last.xml", xmlWrap(xml))
}

func xmlName(i, itemsize int) string {
    return "main-"+ strconv.Itoa(i-itemsize) +"-"+ strconv.Itoa(i) +".xml"
}

func xmlWrap(xml string) string{
    xml_head := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\" xmlns:news=\"http://www.google.com/schemas/sitemap-news/0.9\">"

    return xml_head + xml +"</urlset>"
}

func saveFileStemap(fname, xml string){
    f, err := os.Create(fname)
    checkErr(err)
    byteArray := []byte(xml)
    f.Write(byteArray)
}

func getChangefreq(node_type string) string{
    var changefreq string
    switch  node_type {
        case "article":
            changefreq = "weekly"
            break
        case "blog":
            changefreq = "daily"
            break
        case "cdl_physical":
            changefreq = "monthly"
            break
        case "dmv_office":
            changefreq = "monthly"
            break
        case "directory_entry":
            changefreq = "monthly"
            break
        case "driving_school":
            changefreq = "monthly"
            break
        case "news":
            changefreq = "daily"
            break
        case "page":
            changefreq = "weekly"
            break
        case "page_state":
            changefreq = "weekly"
            break
    }
    return changefreq
}

func Ftoa(f float64) string {
    return strconv.FormatFloat(f, 'f', 1, 64)  
}

func TimeFormat(t int) string {
    return time.Unix(int64(t), 0).Format(stdLongYear +"-"+ stdZeroMonth +"-"+ stdZeroDay)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

