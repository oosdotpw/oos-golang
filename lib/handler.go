package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type Json map[string]interface{}

type HandlerInterface interface {
	Post() int
}

type Handler struct {
	Writer  *http.ResponseWriter
	Request *http.Request

	Header    *http.Header
	PostValue map[string]string
}

func (h *Handler) Init() {
	h.Header = &h.Request.Header

	h.PostValue = make(map[string]string)
	for key := range h.Request.PostForm {
		h.PostValue[key] = h.Request.PostForm.Get(key)
	}
}

func (h *Handler) Filter(field, regex, errmsg string) bool {
	matched, _ := regexp.MatchString(regex, h.PostValue[field])
	if !matched {
		h.Error(errmsg)
		return true
	}
	return false
}

func (h *Handler) Error(msg string) int {
	var data Json = make(Json)

	data["erros_msg"] = msg

	return h.Result(data, true)
}

func (h *Handler) Result(data Json, err bool) int {
	h.Header.Set("Content-Type", "application/json")
	h.Header.Set("Access-Control-Allow-Origin", "*")

	if data == nil {
		data = make(Json)
	}

	data["error"] = err

	b, _ := json.Marshal(data)

	fmt.Fprint(*h.Writer, string(b))
	return 200
}
