package routes

import (
	"encoding/json"
	"net/http"

	"github.com/faztweb/go-postgresql-restapi/db"
	"github.com/faztweb/go-postgresql-restapi/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	db.DB.First(&task, params["id"])
	json.NewEncoder(w).Encode(&task)
}
