package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
    "regexp"
    "strings"
    "strconv"
    . "net/http"
    "os"
    "bufio"
)

const (
    DB_NAME = "dmv"
    DB_USER = "root"
    // DB_PASS = "1"
    // DB_HOST = "tcp(127.0.0.1:8889)"
    DB_PASS = ""
    DB_HOST = "tcp(127.0.0.1:3306)"
)

func main() {
    p := fmt.Println
    _ = p

    getLinks()

    file, _ := os.Open("links.txt")
    output := ""
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    for scanner.Scan() {
        line := scanner.Text()
        arr := strings.Split(line, ";")
        url := arr[1]

        resp, err := Head(url)
        if(err != nil){
            output += line + ";"+ fmt.Sprintf("%v", err) +"\n"
            continue
        }
        defer resp.Body.Close()

        if(resp.StatusCode == 200){
            // output += ";\n"
            continue
        }

        output += line + ";"+ strconv.Itoa(resp.StatusCode) +"\n"
        p(url)

        // os.Exit(0)
    }
    saveFile("broken_links.csv", output)
}

func getLinks(){
    dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
    db, err := sql.Open("mysql", dsn)

    // sql := "SELECT nid, body FROM `node_revisions` WHERE body LIKE '%%http%%' LIMIT 0,10";
    sql := "SELECT nr.nid, nr.body, a.dst FROM `node_revisions` nr, url_alias a WHERE CONCAT('node/', nr.nid) = a.src AND body LIKE '%%http%%'";
    // output := "referrer;url;anchor;status\n"
    output := ""

    rows, err := db.Query(sql)
    checkErr(err)
    for rows.Next() {
        var body, dst string
        var nid int
        err = rows.Scan(&nid, &body, &dst)
        checkErr(err)
        // p(body)
        var re = regexp.MustCompile(`\[http[s]?:\/\/.*?\s.*?\]`)
        results := re.FindAllString(body, -1)
        // results := re.FindAllStringSubmatch(body, -1)
        // p(results)
        // os.Exit(0)

        for _, value := range results {
            // _ = key
            value = value[1:len(value)-1]
            arr := strings.Split(value, " ")
            url := arr[0]
            value = strings.Join(arr[1:], " ")
            value = strings.TrimSpace(value)
            output += "http://www.dmv.com/"+ dst +";"+ url +";"+ value +"\n"
        }
        // qq := []string{"max", "ololo"}
        // _ = qq
    }

    saveFile("links.txt", output)
    db.Close()
}

func saveFile(fname, fileContent string){
    f, err := os.Create(fname)
    checkErr(err)
    byteArray := []byte(fileContent)
    f.Write(byteArray)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
