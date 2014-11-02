package main

import (
	"github.com/go-martini/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/martini-contrib/auth"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"./model"
)

type Blog struct {
	Id          int `form:"Id"`
	Title       string `form:"Title"`
	Date        string `form:"Date"`
	// Description template.HTML `form:"Description"`
	Author      int `form:"Author"`
}

func main() {
	m := martini.Classic()
	m.Use(DB())
	m.Use(render.Renderer())

	m.Use(auth.Basic("admin", "111"))

    m.Get("/wishes", func(r render.Render, db *sql.DB) {
    	u := user.User{1, "admin"}
    	
    	
        r.HTML(200, "list", GetUsers(db))
    })

	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})

	m.Group("/books", func(r martini.Router) {
		r.Get("/:id", GetBooks)
	    r.Post("/new", NewBook)
	    r.Put("/update/:id", UpdateBook)
	    // r.Delete("/delete/:id", DeleteBook)
	})

	m.Use(martini.Static("assets"))

	m.Run()
}

func GetUsers(db *sql.DB) []user.User {
	var userlist []user.User
	sql := "select uid, `name` FROM users ORDER BY uid ASC LIMIT 0,10"
	row, err := db.Query(sql)
	CheckErr(err)
	for row.Next() {
	    var name string
	    var uid int
	    row.Scan(&uid, &name)
	    u := user.User{uid, name}
	    userlist = append(userlist, u)
	}
	return userlist
}

// DB Returns a martini.Handler
func DB() martini.Handler {
    dsn := "root:1@tcp(127.0.0.1:8889)/tmp?charset=utf8"
    db, _ := sql.Open("mysql", dsn)
    // if err != nil {
    //     panic(err)
    // }

    return func(c martini.Context) {
        c.Map(db)
        defer db.Close()
        c.Next()
    }
}

func GetBooks(p martini.Params) string {
	return "Get Books method with: " + p["id"]
}

func NewBook(p martini.Params) string {
	return "NewBook ha ha ha"
}

func UpdateBook(p martini.Params) string {
	return "UpdateBook params: " + p["id"]
}

func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}
