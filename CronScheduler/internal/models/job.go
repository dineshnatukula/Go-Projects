package models

import (
	"github.com/jinzhu/gorm"
)

type ResponseVal struct {
	ErrorCode    int    // 500, 400 etc..
	ErrorMessage string //INTERNAL_SERVER_ERROR, BAD_REQUEST, TECHNICAL_ERROR
}

type Job struct {
	gorm.Model
	ID                 int         `json:"id,omitempty"`
	Name               string      `json:"name" validate:"required"`
	URL                string      `json:"URL" validate:"required,http_url"`
	EndPoint           string      `json:"endPoint" validate:"required,endPoint"`
	HttpMethod         string      `json:"httpMethod" validate:"required,oneof=GET POST PUT DELETE PATCH"`
	ReqBody            string      `json:"reqBody"` // Optional
	Service            string      `json:"service" validate:"required"`
	CronSpec           string      `json:"cronSpec" validate:"required,cron"`
	RetryCount         int         `json:"retrCount"` // Default to 0.
	ResponseValidation ResponseVal `json:"responseValidation"`
	CreatedBy          string      `json:"createdAt,omitempty"`
	CreatedAt          string      `json:"createdBy,omitempty"`
	ModifiedBy         string      `json:"modifiedBy,omitempty"`
	ModifiedAt         string      `json:"modifiedAt,omitempty"`
}

// {
// "id": 1,
// "name": "refilCron",
// "URL": "http://betplacement:8080",
// "endPoint": "/refilCron",
// "httpMethod": "GET"
// "service": "Bet Placement"
// "cronSpec": "*/10 * * * * *"
// }
