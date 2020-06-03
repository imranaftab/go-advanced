package controllers

import "net/http"

// RegisterController Registers a map of routes to the appropriate controller
func RegisterController() {
	pc := newPersonController()

	http.Handle("/person", *pc)
	http.Handle("/person/", *pc)

}
