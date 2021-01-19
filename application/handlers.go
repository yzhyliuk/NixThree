package application

import (
	"NixTwo/data"
	"net/http"
)

func (w *webApp) getPost(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := getID(r.URL.String(), "/posts/")
		if id <= 0 {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Path parameter 'id' should be more than 0"))
			return
		}
		pstEx := new(data.Post)
		post, err := w.generalService.Recive(pstEx, id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}
		switch r.Header.Get("Content-Type") {
		case "application/xml":
			returnXML(rw, post)
			break
		case "application/json":
			returnJSON(rw, post)
			break
		default:
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Request Content-Type should be application/xml or application/json"))
			return
		}

	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
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
