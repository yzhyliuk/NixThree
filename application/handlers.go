package application

import (
	"NixTwo/data"
	"context"
	"net/http"
)

type keyObj struct{}

func (w *webApp) getPost(rw http.ResponseWriter, r *http.Request) {
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
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
		break
	}
}

func (w *webApp) createPost(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

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
