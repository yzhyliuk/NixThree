package post

import (
	"NixTwo/dataSources/mysql"
	"fmt"
)

var (
	dbname = "blogbase"
)

//Post : struct that represents Post
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

//Save : saves post to database
func (p *Post) Save() {
	result := mysql.DataBase.Create(&p)
	if result.Error != nil {
		fmt.Println("Can't add an post to database")
	}
}
