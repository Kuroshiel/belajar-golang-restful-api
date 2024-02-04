package helper

// Category Repository Implementation Golang RESTful API

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
