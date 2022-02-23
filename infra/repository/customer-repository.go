package repository

import "github.com/tarso1156/clean-go/entity"

type CustomerRepository interface {
	Insert(customer entity.Customer) error
}
