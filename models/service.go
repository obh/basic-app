package models

import "time"

type Service struct {
	Id           int              `json:"id"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	VersionCount int              `json:"version_count"`
	CreatedOn    time.Time        `json:"created_on"`
	Versions     []ServiceVersion `json:"versions"`
}

type ServiceVersion struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedOn time.Time `json:"created_on"`
}

type CatalogResponse struct {
	Services []Service     `json:"services"`
	Count    int           `json:"count"`
	Size     int           `json:"size"`
	Request  RequestParams `json:"request"`
}
