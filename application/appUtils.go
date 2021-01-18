package application

import (
	"NixTwo/data/comment"
	"NixTwo/data/post"
	"NixTwo/dataSources/mysql"
	"fmt"
)

//getPostSourceString : returns source URL string for posts of user with given ID
func (pa *parseApp) getPostSourceString(userID int) string{
	return fmt.Sprintf("%s/posts?userId=%d",pa.sourceUrl, userID)
}

//getCommentsSourceString : returns source URL string for comments to post with given ID
func (pa *parseApp) getCommentsSourceString(postID int) string{
	return fmt.Sprintf("%s/comments?postId=%d",pa.sourceUrl, postID)
}

func dbsetup() {
	err := mysql.InitDBConnection("blogbase")
	if err != nil {
		fmt.Println("Can't connect to database")
		return
	}
	err = mysql.DataBase.AutoMigrate(&post.Post{})
	if err != nil {
		fmt.Println(err.Error())
	}
	err = mysql.DataBase.AutoMigrate(&comment.Comment{})
	if err != nil {
		fmt.Println(err.Error())
	}
}


