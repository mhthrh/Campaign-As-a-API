package HttpUtil

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetID(r *http.Request) int {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	return id
}
