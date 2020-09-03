package routes

import (
	"encoding/json"
	"net/http"

	"github.com/kakty/taskmgr/models"
)

//SignIn :
func SignIn(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
