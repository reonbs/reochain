package utilities

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Utilities interface {
	WriteJSON(response http.ResponseWriter, status int, data interface{}, wrap string) error
	ErrorJson(w http.ResponseWriter, err error)
	GetUrlParams(request *http.Request) map[string]string
}

type util struct{}

func NewUtilities() Utilities {
	return &util{}
}

func (*util) WriteJSON(response http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)

	if err != nil {
		return err
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	response.Write(js)

	return nil
}

func (*util) ErrorJson(w http.ResponseWriter, err error) {
	type JsonError struct {
		Message string `json:"message"`
	}

	theError := JsonError{
		Message: err.Error(),
	}

	NewUtilities().WriteJSON(w, http.StatusBadRequest, theError, "error")
}

func (*util) GetUrlParams(request *http.Request) map[string]string {
	queryParameters := make(map[string]string)

	for Key := range request.URL.Query() {
		values, ok := request.URL.Query()[Key]
		if !ok || len(values[0]) < 1 {
			log.Printf("Url Param 'key' is missing %s or has no value", Key)
			continue
		}
		GetKeyValue(Key, values[0], queryParameters)
	}

	return queryParameters
}

func GetKeyValue(key string, value string, m map[string]string) {
	if key == "" {
		return
	}

	if strings.Contains(key, "[") && strings.Contains(key, "]") {
		arr := strings.Split(key, "")
		operator := strings.Join(arr[strings.Index(key, "[")+1:strings.Index(key, "]")], "")
		column := strings.Split(key, "[")
		m[column[0]+"|"+operator] = value
	} else {
		m[key] = value
	}
}
