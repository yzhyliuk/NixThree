package application

import (
	"NixTwo/data/comment"
	"encoding/json"
	"net/http"
	"fmt"
)

func (pa *parseApp) getComments(postID int)  {
	defer pa.wg.Done()
	url := pa.getCommentsSourceString(postID)
	requestComment, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: can't get post #%d from url: %s \n", postID, url)
		return
	}
	var commentsArray []comment.Comment
	err = json.NewDecoder(requestComment.Body).Decode(&commentsArray)

	if err != nil {
		fmt.Println("Unable to unmarshal response body")
		return
	}
	for _, comm := range commentsArray {
		//Todo: logic of saving comments
		fmt.Println(comm.Body)
	}
	pa.wg.Done()
}
