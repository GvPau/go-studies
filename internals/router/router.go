package router

import (
	"encoding/json"
	"net/http"
	"studies-1/internals/books"
	"studies-1/internals/middlewares"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.JsonMiddleware)
	r.Use(middlewares.Logging)
	//r.Use(middlewares.Authentication)
	r.Use(middlewares.RateLimitMiddlware)

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/api", APIHandler).Methods("GET")

	r.HandleFunc("/api/v1/books/filter/{field}/{value}", books.GetFilteredBooksHandler).Methods("GET")
	r.HandleFunc("/api/v1/books/{id}", books.GetBookHandler).Methods("GET")
	r.HandleFunc("/api/v1/books", books.GetBooksHandler).Methods("GET")
	r.HandleFunc("/api/v1/books", books.CreateBookHandler).Methods("POST")

	return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "This is the API endpoint"}
	json.NewEncoder(w).Encode(response)
}
