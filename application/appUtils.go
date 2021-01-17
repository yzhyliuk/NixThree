package application

import "fmt"

//getPostSourceString : returns source URL string for posts of user with given ID
func (pa *parseApp) getPostSourceString(userID int) string{
	return fmt.Sprintf("%s/posts?userId=%d",pa.sourceUrl, userID)
}

//getCommentsSourceString : returns source URL string for comments to post with given ID
func (pa *parseApp) getCommentsSourceString(postID int) string{
	return fmt.Sprintf("%s/comments?userId=%d",pa.sourceUrl, postID)
}


