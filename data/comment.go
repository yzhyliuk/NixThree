package data

import (
	"NixTwo/dataSources/mysql"
	"fmt"
)

var (
	dbname = "blogbase"
)

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

//Save : saves post to database
func (c Comment) Save() {
	result := mysql.DataBase.Create(&c)
	if result.Error != nil {
		fmt.Println("Can't add an post to database")
	}
}
