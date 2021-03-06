package routes

import (
	"encoding/json"
	"net/http"
	"notifications/src/controllers"
	"time"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
	Date    string `json:"date"`
}

func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		json.NewEncoder(w).Encode(Response{Message: "Api is running", Date: now.Format(time.RFC822)})
	})

	router.HandleFunc("/producer", controllers.CreateMessageController).Methods("POST")

	return router
}
