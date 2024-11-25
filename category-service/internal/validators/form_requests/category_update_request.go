package form_requests

import (
	"context"
	"fmt"

	"github.com/nibroos/bookman-go/user-service/internal/dtos"
	"github.com/thedevsaddam/govalidator"
)

// CategoryUpdateRequest handles the validation for the RegisterRequest.
type CategoryUpdateRequest struct {
	Validator *govalidator.Validator
}

// NewRegisterUpdateRequest creates a new instance of CategoryUpdateRequest.
func NewCategoryUpdateRequest() *CategoryUpdateRequest {
	v := govalidator.New(govalidator.Options{})
	return &CategoryUpdateRequest{Validator: v}
}

// Validate validates the RegisterRequest.
func (r *CategoryUpdateRequest) Validate(req *dtos.UpdateCategoryRequest, ctx context.Context) map[string]string {
	rules := govalidator.MapData{
		"name":        []string{"required", fmt.Sprintf("unique_ig:categories,name,%d", req.ID)},
		"description": []string{""},
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
