package validator

import (
	"errors"
	"reflect"
	"regexp"
	"scheduler/internal/models"
	"strings"

	"github.com/go-playground/validator/v10"
)

func RequestValidator(job models.Job) error {
	validate := validator.New()
	validate.RegisterValidation("endPoint", endPointValidation)
	// register function to get tag name from json tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	// validate.RegisterStructValidation(JobLevelValidation, models.Job{})

	// returns InvalidValidationError for bad validation input, nil or ValidationErrors ( []FieldError )
	err := validate.Struct(job)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return err
		}

		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			// for _, err := range validateErrs {
			if err := validateErrs[0]; err != nil {
				return errors.New(err.Field() + " " + err.ActualTag())
			}
		}
	}
	return nil
}

func endPointValidation(fl validator.FieldLevel) bool {
	endPoint := fl.Field().String()
	// Regex: starts with /, followed by letters, numbers, /, -, _
	re := regexp.MustCompile(`^\/[a-zA-Z0-9\/\-_]*$`)
	return re.MatchString(endPoint)
}
