package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/robfig/cron"
)

type Job struct {
	ID             int64         `json:"id,omitempty"`
	CronSpec       string        `json:"cronSpec" validate:"required,cron"`
	Name           string        `json:"name" validate:"required"`
	EndPoint       string        `json:"endPoint" validate:"required"`
	Service        string        `json:"service" validate:"required"`
	ExecTime       time.Time     `json:"execTime" validate:"required"`
	RepeatInterval time.Duration `json:"repeatInterval" validate:"required"`
	Cancelled      chan struct{} `json:"_"`
	CreatedBy      string        `json:"createdAt,omitempty"`
	CreatedAt      string        `json:"createdBy" validate:"required"`
	ModifiedBy     string        `json:"modifiedBy,omitempty"`
	ModifiedAt     string        `json:"modifiedAt,omitempty"`
	Email          string        `json:"email", validate:"required,email"`
}

type ValidationErrorResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	ErrorCode    string `json:"errorCode,omitempty"`
	// Namespace       string `json:"namespace,omitempty"` // can differ when a custom TagNameFunc is registered or
	// Field           string `json:"field,omitempty"`     // by passing alt name to ReportError like below
	// StructNamespace string `json:"structNamespace,omitempty"`
	// StructField     string `json:"structField,omitempty"`
	// Tag             string `json:"tag,omitempty"`
	// ActualTag       string `json:"actualTag,omitempty"`
	// Kind            string `json:"kind,omitempty"`
	// Type            string `json:"type,omitempty"`
	// Value           string `json:"value,omitempty"`
	// Param           string `json:"param,omitempty"`
}

func main() {
	validate := validator.New()
	// register function to get tag name from json tags.
	validate.RegisterValidation("cron", cronValidator)
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	// validate.RegisterStructValidation(JobLevelValidation, Job{})

	jobs := []Job{
		{CronSpec: "0 0 * * * *"},    // Every day at midnight
		{CronSpec: "*/5 * * * *"},    // Every 5 minutes
		{CronSpec: "30 8 * * 1-5"},   // At 8:30 AM, Monday through Friday
		{CronSpec: "15 14 1 * *"},    // At 2:15 PM on the 1st of every month
		{CronSpec: "0 0 1 1 *"},      // At midnight on January 1st
		{CronSpec: "0 9-17 * * 1-5"}, // Every hour from 9 AM to 5 PM on weekdays
		{CronSpec: "0 0 1 */3 *"},    // At midnight on the 1st of every 3rd month
		{CronSpec: "0 0 28 */2 *"},   // At midnight on the 28th of every 2nd month
		{CronSpec: "0 0 29 */2 *"},   // At midnight on the 29th of every 2nd month
		{CronSpec: "0 0 30 */2 *"},   // At midnight on the 30th of every 2nd month
		{CronSpec: "0 0 31 */2 *"},   // At midnight on the 31th of every 2nd month
		{CronSpec: "0 0 32 */3 *"},   // At midnight on the 32nd of every 3nd month
		// {CronSpec: "0 0 * * * *"},
		{CronSpec: ""},
		{CronSpec: "* * *"},       // Too few fields (needs 5)
		{CronSpec: "60 0 * * *"},  // Invalid minute (60 is out of range)
		{CronSpec: "0 24 * * *"},  // Invalid hour (24 is out of range)
		{CronSpec: "0 0 32 * *"},  // Invalid day of month (32)
		{CronSpec: "0 0 * 13 *"},  // Invalid month (13)
		{CronSpec: "0 0 * * 7"},   // Day of week (7 is invalid for 0-6 range in robfig cron)
		{CronSpec: "hello world"}, // Nonsense string
		{CronSpec: "@daily"},      // Only valid in `cron.Parse`, not `ParseStandard`
	}

	// returns InvalidValidationError for bad validation input, nil or ValidationErrors ( []FieldError )
	for _, job := range jobs {
		err := validate.Struct(job)
		if err != nil {
			// fmt.Printf("Job %v invalid: %v\n", job, err)
			// // this check is only needed when your code could produce
			// // an invalid value for validation such as interface with nil
			// // value most including myself do not usually have code like this.
			var invalidValidationError *validator.InvalidValidationError
			if errors.As(err, &invalidValidationError) {
				fmt.Println(err)
				return
			}

			var validateErrs validator.ValidationErrors
			if errors.As(err, &validateErrs) {
				if len(validateErrs) > 0 {
					// err1 := validateErrs[0]
					// e := ValidationErrorResponse{
					// 	ErrorCode:    "UNPROCESSIBLE_ENTITY",
					// 	ErrorMessage: err.Field() + " " + err.ActualTag(),
					// }
					fmt.Printf("Job %v is invalid❌\n", job)
				}
			}
		} else {
			fmt.Printf("Job %v valid ✅\n", job)
		}
	}
}

// JobLevelValidation contains custom struct level validations that don't always
// make sense at the field validation level. For Example this function validates that either
// EndPoint or RepeatInterval exist; could have done that with a custom field validation but then
// would have had to add it to both fields duplicating the logic + overhead, this way it's
// only validated once.
//
// NOTE: you may ask why wouldn't I just do this outside of validator, because doing this way
// hooks right into validator and you can combine with validation tags and still have a
// common error output format.
// func JobLevelValidation(sl validator.StructLevel) {

// 	job := sl.Current().Interface().(Job)

// 	if len(job.EndPoint) == 0 || len(job.Service) == 0 {
// 		sl.ReportError(job.EndPoint, "endPoint", "endPoint", "EndPoint", "")
// 		sl.ReportError(job.Service, "service", "service", "Service", "")
// 	}

// 	// plus can do more, even with different tag than "fnameorlname"
// }

func cronValidator(fl validator.FieldLevel) bool {
	_, err := cron.ParseStandard(fl.Field().String())
	return err == nil
}
