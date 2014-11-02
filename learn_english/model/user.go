package user

type User struct {
	Id          int `form:"uid"`
	Title       string `form:"name"`
}
