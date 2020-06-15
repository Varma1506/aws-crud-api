package models

type CreateRequestBody struct {
	Id          string  `json:"id"`
	Category    string  `json:"category"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string
}

type UpdateAndDeleteRequestBody struct {
	Id string `json:"id"`
}

type CronEvent struct {
	Resources string `json:"resources"`
}
