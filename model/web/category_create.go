package web

// Category Service Golang RESTful API

type CategoryCreateRequest struct {
	Name string `validate:"required,max=200,min=1" json:"name"`
}
