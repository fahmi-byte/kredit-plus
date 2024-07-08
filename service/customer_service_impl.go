package service

import (
	"context"
	"database/sql"
	"errors"
	"kredit-plus/helper"
	"kredit-plus/model/domain"
	"kredit-plus/model/repository"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	DB                 *sql.DB
}

func NewCustomerService(customerRepository repository.CustomerRepository, DB *sql.DB) *CustomerServiceImpl {
	return &CustomerServiceImpl{CustomerRepository: customerRepository, DB: DB}
}

func (service *CustomerServiceImpl) CheckTenorLimitCustomer(ctx context.Context, customerId int, tenor string, otr float32, pin string, ch chan bool, errChan chan error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	isValid := service.CustomerRepository.CheckTenorLimit(ctx, tx, customerId, tenor, otr, pin)
	if !isValid {
		errChan <- errors.New("limit tenor is not enough")
	}
	ch <- true
}

func (service *CustomerServiceImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer {
	//TODO implement me
	panic("implement me")
}
