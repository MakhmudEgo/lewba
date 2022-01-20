package handler

import (
	"log"
	"net/http"
)

type testHandler struct {
}

func TestHandler() *testHandler {
	return &testHandler{}
}

func (t testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("username")
}
