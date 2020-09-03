package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kakty/taskmgr/models"
	"gorm.io/gorm"
)

func getTask(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var task models.Task
	id := r.URL.Query()["id"]
	fmt.Println(id)
	result := db.Table("tasks").First(&task, 1)
	if result.Error != nil {
		log.Fatal("DB QUERY ERROR")
	}
	fmt.Println(task)
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(task)
	if err != nil {
		log.Fatal(err)
	}
}
