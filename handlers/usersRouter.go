package handlers

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func UsersRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "/users" {
		switch r.Method {
		case http.MethodGet:
			userGetAll(w, r)
			return
		case http.MethodPost:
			userPostOne(w, r)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
		}
	}
	path = strings.TrimPrefix(path, "/users/")
	if !bson.IsObjectIdHex(path) {
		postError(w, http.StatusNotFound)
	}
	id := bson.ObjectIdHex(path)

	switch r.Method {
	case http.MethodGet:
		userGetOne(w, r, id)
		return
	case http.MethodPut:
		userPutOne(w, r, id)
		return
	case http.MethodPatch:
		userPacthOne(w, r, id)
		return
	case http.MethodDelete:
		userDeleteOne(w, r, id)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
