package main

import (
	"log"
	"net/http"
)

type DefaultHandler struct{}

func (h *DefaultHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Printf("GET %s\n", request.URL.Path)
	writer.Write([]byte("ok"))
}
