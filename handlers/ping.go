package handlers

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

func PingHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "pong")
}

