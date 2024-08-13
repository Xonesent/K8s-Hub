package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate // nolint:gochecknoglobals // Needed for proper validator work

func init() {
	validate = validator.New()
}
