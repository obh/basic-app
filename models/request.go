package models

type RequestParams struct {
	CreatedAfter  int    `json:"created_after" validate:"min=0,max=10000000"`
	Limit         int    `json:"limit" validate:"min=1,max=50"`
	CreatedBefore int    `json:"created_before" validate:"min=0,max=10000000"`
	Size          int    `json:"size"`
	FilterBy      string `json:"filter_by" validate:"max=250"`
}

type ServiceVersionsCount struct {
	ServiceId int `json:"service_id"`
	Count     int `json:"count"`
}
