package form_requests

import (
	"context"

	"github.com/nibroos/bookman-go/user-service/internal/dtos"
	"github.com/thedevsaddam/govalidator"
)

// BookStoreRequest handles the validation for the RegisterRequest.
type BookStoreRequest struct {
	Validator *govalidator.Validator
}

// NewRegisterStoreRequest creates a new instance of BookStoreRequest.
func NewBookStoreRequest() *BookStoreRequest {
	v := govalidator.New(govalidator.Options{})
	return &BookStoreRequest{Validator: v}
}

// Validate validates the RegisterRequest.
func (r *BookStoreRequest) Validate(req *dtos.CreateBookRequest, ctx context.Context) map[string]string {
	// utils.DD(req)
	rules := govalidator.MapData{
		"name": []string{"required"},
		// "name":        []string{"required", "unique:authors,name"},
		"description": []string{},
		"author_id":   []string{"required"},
		"category_id": []string{},
	}

	opts := govalidator.Options{
		Data:  req,
		Rules: rules,
	}

	v := govalidator.New(opts)
	mappedErrors := v.ValidateStruct()

	if len(mappedErrors) == 0 {
		return nil
	}

	errors := make(map[string]string)
	for field, err := range mappedErrors {
		errors[field] = err[0]
	}
	return errors
}
