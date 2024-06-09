package validation

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ErrorResponse represents a validation error response
type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidateStruct validates a struct based on the tags and returns an error response if validation fails
func ValidateStruct(s interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.Field()
			element.Message = "Field validation for '" + err.Field() + "' failed on the '" + err.Tag() + "' tag"
			errors = append(errors, &element)
		}
	}
	return errors
}

// WriteJSONResponse writes a JSON response to the client
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
