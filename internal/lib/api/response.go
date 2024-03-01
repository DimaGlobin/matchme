package api

import (
	"net/http"

	"github.com/go-chi/render"
)

func Respond(w http.ResponseWriter, r *http.Request, code int, v interface{}) {
	render.Status(r, code)
	render.JSON(w, r, v)
}
