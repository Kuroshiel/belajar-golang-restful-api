package web

// Category Service Golang RESTful API

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required, max = 200, min = 1"`
}
