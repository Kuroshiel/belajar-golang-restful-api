package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Category Controller Golang RESTful API

type CategoryController interface {
	Create(writer http.ResponseWriter, response *http.Request, Params httprouter.Params)
	Update(writer http.ResponseWriter, response *http.Request, Params httprouter.Params)
	Delete(writer http.ResponseWriter, response *http.Request, Params httprouter.Params)
	FindById(writer http.ResponseWriter, response *http.Request, Params httprouter.Params)
	FindByAll(writer http.ResponseWriter, response *http.Request, Params httprouter.Params)
}
