package model

type Product struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

type Products struct {
	Products []Product `json:"products"`
}
