package controllers

import (
	"net/http"
	"regexp"
)

// PersonController Controller to handle request for person data
type PersonController struct {
	personIDPattern *regexp.Regexp
}

func (pc PersonController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! Welcome to Go with Imran."))
}

func newPersonController() *PersonController {
	return &PersonController{
		personIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
