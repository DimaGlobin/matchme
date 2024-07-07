package api

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err string `json:"err"`
}

type IdResponse struct {
	Id  uint64 `json:"id"`
	Msg string `json:"msg"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func Respond(w http.ResponseWriter, r *http.Request, code int, v interface{}) {
	render.Status(r, code)
	render.JSON(w, r, v)
}
