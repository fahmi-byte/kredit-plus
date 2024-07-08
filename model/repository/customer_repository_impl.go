package repository

import (
	"context"
	"database/sql"
	"fmt"
	"kredit-plus/helper"
	"kredit-plus/model/domain"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{}
}

func (repository *CustomerRepositoryImpl) CheckTenorLimit(ctx context.Context, tx *sql.Tx, customerId int, tenor string, otr float32, pin string) bool {
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

	if otr <= tenorCustomer && pin == existPin {
		newLimit := tenorCustomer - otr
		SQLUpdate := fmt.Sprintf("UPDATE tenor_customers SET %s = $1 WHERE customer_id = $2", tenor)
		_, err := tx.ExecContext(ctx, SQLUpdate, newLimit, customerId)
		helper.PanicIfError(err)

		return true
	}

	return false
}

func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer {
	//TODO implement me
	panic("implement me")
}
