package form_requests

import (
	"context"

	"github.com/nibroos/bookman-go/user-service/internal/dtos"
	"github.com/thedevsaddam/govalidator"
)

// CategoryStoreRequest handles the validation for the RegisterRequest.
type CategoryStoreRequest struct {
	Validator *govalidator.Validator
}

// NewRegisterStoreRequest creates a new instance of CategoryStoreRequest.
func NewCategoryStoreRequest() *CategoryStoreRequest {
	v := govalidator.New(govalidator.Options{})
	return &CategoryStoreRequest{Validator: v}
}

// Validate validates the RegisterRequest.
func (r *CategoryStoreRequest) Validate(req *dtos.CreateCategoryRequest, ctx context.Context) map[string]string {
	// utils.DD(req)
	rules := govalidator.MapData{
		"name":        []string{"required", "unique:categories,name"},
		"description": []string{},
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
