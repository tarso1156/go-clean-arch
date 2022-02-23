package usecase

import (
	"github.com/google/uuid"
	"github.com/tarso1156/clean-go/dto"
	"github.com/tarso1156/clean-go/entity"
	"github.com/tarso1156/clean-go/repository"
)

type CreateCustomer struct {
	Repository repository.CustomerRepository
}

func (c CreateCustomer) Execute(input dto.CreateCustomerInputDto) (dto.CreateCustomerOutputDto, error) {
	customer := entity.Customer{}
	customer.Id = uuid.New().String()
	customer.Name = input.Name
	customer.Age = input.Age
	customer.AvgScore = input.AvgScore

	err := c.Repository.Insert(customer)
	if err != nil {
		return dto.CreateCustomerOutputDto{}, err
	}

	output := dto.CreateCustomerOutputDto{}
	output.Id = customer.Id
	output.Name = customer.Name
	output.Age = customer.Age
	output.AvgScore = customer.AvgScore

	return output, nil
}
