package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Category Controller Golang RESTful API

type CategoryController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
