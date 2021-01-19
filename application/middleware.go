package application

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strconv"
	"strings"
)

func getID(URL string, cut string) int {
	str := strings.Trim(URL, cut)
	id, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return id
}

//Encodes XML to Response Writer
func returnXML(rw http.ResponseWriter, obj interface{}) {
	encoder := xml.NewEncoder(rw)
	rw.WriteHeader(http.StatusOK)
	encoder.Encode(obj)
}

//Encodes json to Response Writer
func returnJSON(rw http.ResponseWriter, obj interface{}) {
	encoder := json.NewEncoder(rw)
	rw.WriteHeader(http.StatusOK)
	encoder.Encode(obj)
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
		return
	}
}
