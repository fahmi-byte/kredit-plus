package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"kredit-plus/helper"
	"kredit-plus/model/domain"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{}
}

func (repository *CustomerRepositoryImpl) CheckTenorLimit(ctx context.Context, tx *sql.Tx, customerId int, tenor string, otr float32, pin string) error {
	SQL := fmt.Sprintf("SELECT %s, pin FROM tenor_customers JOIN customers on customers.id = tenor_customers.customer_id WHERE customer_id = $1", tenor)
	rows, err := tx.QueryContext(ctx, SQL, customerId)
	helper.PanicIfError(err)

	var tenorCustomer float32
	var existPin string
	if rows.Next() {
		err := rows.Scan(&tenorCustomer, &existPin)
		helper.PanicIfError(err)
	} else {
		panic("NOT FOUND")
	}
	rows.Close()

	if pin != existPin {
		return errors.New("Wrong Pin!")
	}

	if otr > tenorCustomer {
		return errors.New("Customer limit is not enough!")
	}

	newLimit := tenorCustomer - otr
	SQLUpdate := fmt.Sprintf("UPDATE tenor_customers SET %s = $1 WHERE customer_id = $2", tenor)
	_, err = tx.ExecContext(ctx, SQLUpdate, newLimit, customerId)
	helper.PanicIfError(err)

	return nil
}

func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer {
	//TODO implement me
	panic("implement me")
}

func (repository *CustomerRepositoryImpl) CreateCustomer(ctx context.Context, tx *sql.Tx, customer domain.Customer) {
	hashPin := helper.Hash(customer.Pin)
	var lastCustomerId int
	SQL := "INSERT INTO customers(user_id, full_name, birth_place, birth_date, salary, identity_card, selfie_photo, pin) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	err := tx.QueryRowContext(ctx, SQL, customer.UserId, customer.FullName, customer.BirthPlace, customer.BirthDate, customer.Salary, customer.IdentityCard, customer.SelfiePhoto, hashPin).Scan(&lastCustomerId)
	helper.PanicIfError(err)

	tenor1 := customer.Salary * 0.30
	tenor2 := customer.Salary * 0.50
	tenor3 := customer.Salary * 0.60
	tenor4 := customer.Salary * 0.75

	SQLInsertCustomerTenor := "INSERT INTO tenor_customers (customer_id, tenor_1, tenor_2, tenor_3, tenor_4) VALUES($1, $2, $3, $4, $5)"
	_, err = tx.ExecContext(ctx, SQLInsertCustomerTenor, lastCustomerId, tenor1, tenor2, tenor3, tenor4)
	helper.PanicIfError(err)
}
