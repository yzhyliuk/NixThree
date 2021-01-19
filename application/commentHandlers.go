package application

import (
	"NixTwo/data"
	"context"
	"fmt"
	"net/http"
)

func (w *webApp) Comment(rw http.ResponseWriter, r *http.Request) {
	//getting id from path URL
	id := getID(r.URL.String(), "/comments/")
	if id <= 0 {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Path parameter 'id' should be more than 0"))
		return
	}
	//Switch statment for get and delete methods
	switch r.Method {
	//Hndling Get request
	case http.MethodGet:
		comEx := new(data.Comment)
		comment, err := w.generalService.Recive(comEx, id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}
		newContext := context.WithValue(r.Context(), keyObj{}, comment)
		r = r.WithContext(newContext)
		w.returnHandler(rw, r)
		break
	//Handling Delete request
	case http.MethodDelete:
		commEx := new(data.Comment)
		err := w.generalService.Delete(commEx, id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}
		rw.WriteHeader(http.StatusOK)
		break
	//Handling Put request for comment updating
	case http.MethodPut:
		//getting current commnet with given id
		currentComment := new(data.Comment)
		_, err := w.generalService.Recive(currentComment, id)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			msg := fmt.Sprintf("Post with id %d dosen't exist in database", id)
			rw.Write([]byte(msg))
			return
		}
		comment := new(data.Comment)
		//parsing object from request body
		err = parseObj(comment, rw, r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Can't parse Comment from request body"))
			return
		}
		comment.ID = id
		//merging recived comment with one in our database
		comment.Merge(currentComment)
		//updating comment
		err = w.generalService.Update(comment)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			msg := fmt.Sprintf("Uable to update Comment: %s", err.Error())
			rw.Write([]byte(msg))
			return
		}
		break
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
		break
	}
}

func (w *webApp) createComment(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		obj := new(data.Comment)
		err := parseObj(obj, rw, r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Can't parse Comment from request body"))
			return
		}
		err = w.generalService.Create(obj)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			msg := fmt.Sprintf("Uable to create new Comment: %s", err.Error())
			rw.Write([]byte(msg))
			return
		}
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
