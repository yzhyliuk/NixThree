package application

import (
	"encoding/json"
	"encoding/xml"
	"io"
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

func parseObjXML(obj interface{}, body io.ReadCloser) error {
	decoder := xml.NewDecoder(body)
	err := decoder.Decode(obj)
	if err != nil {
		return err
	}
	return nil
}
func parseObjJSON(obj interface{}, body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(obj)
	if err != nil {
		return err
	}
	return nil
}
func parseObj(obj interface{}, rw http.ResponseWriter, r *http.Request) error {
	switch r.Header.Get("Content-Type") {
	case "application/xml":
		err := parseObjXML(obj, r.Body)
		if err != nil {
			return err
		}
		break
	case "application/json":
		err := parseObjJSON(obj, r.Body)
		if err != nil {
			return err
		}
		break
	default:
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Request Content-Type should be application/xml or application/json"))
		break
	}
	return nil
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
