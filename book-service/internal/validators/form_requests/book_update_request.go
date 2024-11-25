package form_requests

import (
	"context"
	"fmt"

	"github.com/nibroos/bookman-go/user-service/internal/dtos"
	"github.com/thedevsaddam/govalidator"
)

// BookUpdateRequest handles the validation for the RegisterRequest.
type BookUpdateRequest struct {
	Validator *govalidator.Validator
}

// NewRegisterUpdateRequest creates a new instance of BookUpdateRequest.
func NewBookUpdateRequest() *BookUpdateRequest {
	v := govalidator.New(govalidator.Options{})
	return &BookUpdateRequest{Validator: v}
}

// Validate validates the RegisterRequest.
func (r *BookUpdateRequest) Validate(req *dtos.UpdateBookRequest, ctx context.Context) map[string]string {
	rules := govalidator.MapData{
		"name":        []string{"required", fmt.Sprintf("unique_ig:books,name,%d", req.ID)},
		"description": []string{""},
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
