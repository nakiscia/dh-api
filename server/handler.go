package server

import (
	"dh-api/core/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const CannotReadRequestBody = "cannot read request body"
const KeyOrValueCannotBeNull = "key or value cannot be null in request body"
const ValidJsonShouldBeProvided = "valid json should be provided"

type DefaultHandler struct{
	KeyService service.KeyServiceInterface
}

func NewDefaultHandler(keyService service.KeyServiceInterface) DefaultHandler{
	return DefaultHandler{KeyService: keyService}
}

func (h *DefaultHandler) HandleRequests(){
	http.HandleFunc("/key", h.KeyRequestHandler)
}

func (h *DefaultHandler) KeyRequestHandler(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("content-type", "application/json")
	switch request.Method {
	case "GET":
		if key,ok := request.URL.Query()["key"]; ok{
			data, err := h.KeyService.GetKey(key[0])
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`,err.Error())))
				return
			}

			marshal, _ := json.Marshal(map[string]string{
				"value": data,
			})

			writer.WriteHeader(http.StatusOK)
			writer.Write(marshal)
		}
	case "POST":
		body, err := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		if err != nil{
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`,CannotReadRequestBody)))
			return
		}

		var data map[string]string
		err = json.Unmarshal(body,&data)

		if err != nil{
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`,ValidJsonShouldBeProvided)))
			return
		}

		key, keyExist := data["key"]
		value, valueExist := data["value"]

		if !keyExist || !valueExist {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`,KeyOrValueCannotBeNull)))
			return
		}

		err = h.KeyService.SetKey(key, value)

		if err != nil{
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`,err)))
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write(body)
	}
}
