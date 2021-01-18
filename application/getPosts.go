package application

import (
	"NixTwo/data/post"
	"encoding/json"
	"fmt"
	"net/http"
)

func (pa *parseApp) getPostByUserID(userID int) {
	url := pa.getPostSourceString(userID)
	requestPost, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: can't get post of user #%d from url: %s \n", userID, url)
		return
	}
	var postArray []post.Post
	err = json.NewDecoder(requestPost.Body).Decode(&postArray)
	if err != nil {
		fmt.Println("Unable to unmarshal response body")
		return
	}
	for _,pst := range postArray {
		pa.wg.Add(1)
		go pa.getComments(pst.ID)
		obj := pst
		go func() {
			obj.Save()
			pa.wg.Done()
		}()
	}
}
