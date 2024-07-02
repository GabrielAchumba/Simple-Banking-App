package models

type AccountModel struct {
	ID      int     `json:"id"`
	Number  string  `json:"number"`
	Balance float64 `json:"balance"`
}
