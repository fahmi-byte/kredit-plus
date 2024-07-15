package service

import (
	"context"
	"database/sql"
	"fmt"
	"kredit-plus/helper"
	"kredit-plus/model/domain"
	"kredit-plus/model/repository"
	"kredit-plus/model/web"
	"time"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	DB                 *sql.DB
}

func NewCustomerService(customerRepository repository.CustomerRepository, DB *sql.DB) *CustomerServiceImpl {
	return &CustomerServiceImpl{CustomerRepository: customerRepository, DB: DB}
}

func (service *CustomerServiceImpl) CheckTenorLimitCustomer(ctx context.Context, customerId int, tenor string, otr float32, pin string, validTenorChan chan bool, errChan chan error) {
	fmt.Println("masuk kedalam cek tenor")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.CustomerRepository.CheckTenorLimit(ctx, tx, customerId, tenor, otr, pin)
	if err != nil {
		fmt.Println("error tenor")
		errChan <- err
	} else {
		validTenorChan <- true
	}

}

func (service *CustomerServiceImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer {
	//TODO implement me
	panic("implement me")
}

func (service *CustomerServiceImpl) CreateCustomer(ctx context.Context, customer web.CustomerRequest) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	date, err := time.Parse("2006-01-02", customer.BirthDate)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	requestCustomer := domain.Customer{
		UserId:       customer.UserId,
		FullName:     customer.FullName,
		BirthDate:    date,
		BirthPlace:   customer.BirthPlace,
		IdentityCard: customer.IdentityCard,
		SelfiePhoto:  customer.SelfiePhoto,
		Salary:       customer.Salary,
		Pin:          customer.Pin,
	}
	service.CustomerRepository.CreateCustomer(ctx, tx, requestCustomer)
}
