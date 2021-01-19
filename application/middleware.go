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

func returnXML(rw http.ResponseWriter, obj interface{}) {
	encoder := xml.NewEncoder(rw)
	rw.WriteHeader(http.StatusOK)
	encoder.Encode(obj)
}

func returnJSON(rw http.ResponseWriter, obj interface{}) {
	encoder := json.NewEncoder(rw)
	rw.WriteHeader(http.StatusOK)
	encoder.Encode(obj)
}
