package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

//UseRoutes :
func UseRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/", homePage)
	r.HandleFunc("/getTask", func(w http.ResponseWriter, r *http.Request) {
		getTask(w, r, db)
	})
	r.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		SignIn(w, r, db)
	})
	r.HandleFunc("/welcome", Welcome)
}
