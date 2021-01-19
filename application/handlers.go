package application

import (
	"NixTwo/data"
	"context"
	"fmt"
	"net/http"
)

type keyObj struct{}

func (w *webApp) Post(rw http.ResponseWriter, r *http.Request) {
	//getting id from path URL
	id := getID(r.URL.String(), "/posts/")
	if id <= 0 {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Path parameter 'id' should be more than 0"))
		return
	}
	//Switch statment for get and delete methods
	switch r.Method {
	//Hndling Get request
	case http.MethodGet:
		pstEx := new(data.Post)
		post, err := w.generalService.Recive(pstEx, id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}
		newContext := context.WithValue(r.Context(), keyObj{}, post)
		r = r.WithContext(newContext)
		w.returnHandler(rw, r)
		break
	//Handling Delete request
	case http.MethodDelete:
		pstEx := new(data.Post)
		err := w.generalService.Delete(pstEx, id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}
		rw.WriteHeader(http.StatusOK)
		break
	//Handling Put request for post updating
	case http.MethodPut:
		//getting current post with given id
		currentPost := new(data.Post)
		_, err := w.generalService.Recive(currentPost, id)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			msg := fmt.Sprintf("Post with id %d dosen't exist in database", id)
			rw.Write([]byte(msg))
			return
		}
		post := new(data.Post)
		//parsing object from request body
		err = parseObj(post, rw, r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Can't parse Post from request body"))
			return
		}
		post.ID = id
		//merging recived post with one in our database
		post.Merge(currentPost)
		//updating post
		err = w.generalService.Update(post)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			msg := fmt.Sprintf("Uable to update Post: %s", err.Error())
			rw.Write([]byte(msg))
			return
		}
		break
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
		break
	}
}

func (w *webApp) createPost(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		obj := new(data.Post)
		err := parseObj(obj, rw, r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Can't parse Post from request body"))
			return
		}
		err = w.generalService.Create(obj)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			msg := fmt.Sprintf("Uable to create new Post: %s", err.Error())
			rw.Write([]byte(msg))
			return
		}
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (w *webApp) updatePost(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {

	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (w *webApp) deletePost(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {

	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

//handler that returns XML or JSON presentation of data depending of Content-type of request
func (w *webApp) returnHandler(rw http.ResponseWriter, r *http.Request) {
	//getting object from context
	obj := r.Context().Value(keyObj{})
	//checkong Content-Type header and calling return function
	switch r.Header.Get("Content-Type") {
	case "application/xml":
		returnXML(rw, obj)
		break
	case "application/json":
		returnJSON(rw, obj)
		break
	default:
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Request Content-Type should be application/xml or application/json"))
		break
	}
}
