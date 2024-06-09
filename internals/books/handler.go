package books

import (
	"encoding/json"
	"net/http"
	"reflect"
	"studies-1/internals/validation"

	"github.com/gorilla/mux"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	errors := validation.ValidateStruct(newBook)
	if errors != nil {
		validation.WriteJSONResponse(w, http.StatusBadRequest, errors)
		return
	}

	books = append(books, newBook)
	validation.WriteJSONResponse(w, http.StatusCreated, newBook)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	for _, book := range books {
		if book.ID == bookID {
			validation.WriteJSONResponse(w, http.StatusOK, book)
			return
		}
	}

	validation.WriteJSONResponse(w, http.StatusNotFound, map[string]string{"error": "Resource not found"})
}

func GetFilteredBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	vars := mux.Vars(r)
	field := vars["field"]
	filterValue := vars["value"]

	caser := cases.Title(language.Und)

	var filteredBooks []Book
	for _, book := range books {
		fieldName := caser.String(field)
		fieldValue := reflect.ValueOf(book).FieldByName(fieldName).Interface()

		if fieldValue == filterValue {
			filteredBooks = append(filteredBooks, book)
		}
	}

	validation.WriteJSONResponse(w, http.StatusOK, filteredBooks)
}
