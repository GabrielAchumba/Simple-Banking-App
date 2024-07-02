package models

type AccountModel struct {
	ID        int     `json:"id"`
	Reference string  `json:"reference"`
	Balance   float64 `json:"balance"`
}
