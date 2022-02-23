package dto

type CreateCustomerInputDto struct {
	Name     string  `json:"name"`
	Age      int16   `json:"age"`
	AvgScore float32 `json:"avgScore"`
}

type CreateCustomerOutputDto struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Age      int16   `json:"age"`
	AvgScore float32 `json:"avgScore"`
}
