package books

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
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
	// Extract and validate parameters
	field, filterValue, operator, err := extractAndValidateParams(r)
	if err != nil {
		log.Printf("Error: %v", err)
		validation.WriteJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	log.Printf("Received path params request to filter books by %s: %s", field, filterValue)
	log.Printf("Received query params operator: %s", operator)

	// Filter books based on parameters
	filteredBooks, err := filterBooks(field, filterValue, operator)
	if err != nil {
		log.Printf("Error filtering books: %v", err)
		validation.WriteJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	log.Printf("Found %d books matching the filter criteria", len(filteredBooks))
	validation.WriteJSONResponse(w, http.StatusOK, filteredBooks)
}

func extractAndValidateParams(r *http.Request) (string, string, string, error) {
	vars := mux.Vars(r)
	field := vars["field"]
	filterValue := vars["value"]

	query := r.URL.Query()
	operator := query.Get("operator")

	if field == "" || filterValue == "" {
		return "", "", "", fmt.Errorf("field and value path parameters are required")
	}

	return field, filterValue, operator, nil
}

func filterBooks(field, filterValue, operator string) ([]Book, error) {
	caser := cases.Title(language.Und)
	fieldName := caser.String(field)

	filteredBooks := []Book{}
	for _, book := range books {
		fieldValue := reflect.ValueOf(book).FieldByName(fieldName).Interface()

		if fieldName == "Price" {
			bookPrice, ok := fieldValue.(float64)
			if !ok {
				return nil, fmt.Errorf("invalid price field value")
			}

			price, err := strconv.ParseFloat(filterValue, 64)
			if err != nil {
				return nil, fmt.Errorf("error parsing filterValue: %v", err)
			}

			if matchPrice(bookPrice, price, operator) {
				filteredBooks = append(filteredBooks, book)
			}
		} else {
			if fieldValue == filterValue {
				filteredBooks = append(filteredBooks, book)
			}
		}
	}
	return filteredBooks, nil
}

func matchPrice(bookPrice, price float64, operator string) bool {
	switch operator {
	case "gt":
		return bookPrice > price
	case "lt":
		return bookPrice < price
	case "gte":
		return bookPrice >= price
	case "lte":
		return bookPrice <= price
	default:
		log.Printf("Invalid operator: %s", operator)
		return false
	}
}
