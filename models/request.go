package models

import (
	"time"
)

type RequestParams struct {
	CreatedAfter  int    `json:"created_after"`
	Limit         int    `json:"limit" validate:"min=1,max=100"`
	CreatedBefore int    `json:"created_before"`
	Size          int    `json:"size"`
	FilterBy      string `json:"filter_by" validate:"max=250"`
}

type ServiceVersionCount struct {
	ServiceId int `json:"service_id"`
	Count     int `json:"count"`
}

type ServiceWithVersions struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"created_on"`
	Versions    int       `json:"versions"`
}
