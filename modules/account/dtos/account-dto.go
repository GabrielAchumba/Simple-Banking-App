package dtos

type CreateAccountDTO struct {
	Reference string  `json:"reference"`
	Balance   float64 `json:"balance" binding:"required"`
}
