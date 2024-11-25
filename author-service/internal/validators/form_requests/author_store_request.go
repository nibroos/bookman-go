package form_requests

import (
	"context"

	"github.com/nibroos/bookman-go/user-service/internal/dtos"
	"github.com/thedevsaddam/govalidator"
)

// AuthorStoreRequest handles the validation for the RegisterRequest.
type AuthorStoreRequest struct {
	Validator *govalidator.Validator
}

// NewRegisterStoreRequest creates a new instance of AuthorStoreRequest.
func NewAuthorStoreRequest() *AuthorStoreRequest {
	v := govalidator.New(govalidator.Options{})
	return &AuthorStoreRequest{Validator: v}
}

// Validate validates the RegisterRequest.
func (r *AuthorStoreRequest) Validate(req *dtos.CreateAuthorRequest, ctx context.Context) map[string]string {
	// utils.DD(req)
	rules := govalidator.MapData{
		"name":        []string{"required", "unique:authors,name"},
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
