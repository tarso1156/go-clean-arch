package repository

import (
	"database/sql"

	"github.com/tarso1156/clean-go/entity"
)

type CustomerMySqlRepository struct {
	Db *sql.DB
}

func (c CustomerMySqlRepository) Insert(customer entity.Customer) error {
	stmt, err := c.Db.Prepare(
		`
INSERT INTO customers(id, name, age, avg_score)
VALUES
(?,?,?,?)
`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		customer.Id,
		customer.Name,
		customer.Age,
		customer.AvgScore,
	)
	if err != nil {
		return err
	}
	return nil
}
