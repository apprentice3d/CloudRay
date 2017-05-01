package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func createMatchHandler(formater *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formater.JSON(w, http.StatusOK,
			struct{ Test string }{"This is a test string"})
	}
}
