package personcontroller

import (
	"net/http"
	"regexp"
)

type personController struct {
	personIDPattern *regexp.Regexp
}

func (pc personController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! Welcome to Go with Imran."))
}

func newPersonController() *personController {
	return &personController{
		personIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
