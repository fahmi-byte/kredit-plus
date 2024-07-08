package helper

import (
	"kredit-plus/model/domain"
	"kredit-plus/model/web"
)

func ToTransactionResponse(transaction domain.Transaction) web.TransactionResponse {
	return web.TransactionResponse{
		Id:             transaction.Id,
		MerchantID:     transaction.MerchantId,
		CustomerId:     transaction.CustomerId,
		OTR:            transaction.OTR,
		AdminFee:       transaction.AdminFee,
		AssetName:      transaction.AssetName,
		InterestRateId: transaction.InterestRateId,
		InterestAmount: transaction.InterestAmount,
		Tenor:          transaction.Tenor,
	}
}

func ToTransactionsResponse(transactions []domain.Transaction) []web.TransactionResponseList {
	responseCategories := []web.TransactionResponseList{}
	for _, transaction := range transactions {
		responseCategories = append(responseCategories, web.TransactionResponseList{
			Id:                transaction.Id,
			MerchantName:      transaction.MerchantName,
			CustomerName:      transaction.FullName,
			OTR:               transaction.OTR,
			AdminFee:          transaction.AdminFee,
			InstallmentAmount: transaction.InstallmentAmount,
			InterestAmount:    transaction.InterestAmount,
			AssetName:         transaction.AssetName,
			InterestRate:      transaction.InterestRate,
			TransactionDate:   transaction.TransactionDate,
			Tenor:             transaction.Tenor,
		})
	}

	return responseCategories
}

func ToMerchantResponse(merchant domain.Merchant) web.MerchantResponse {
	return web.MerchantResponse{
		Id:           merchant.Id,
		MerchantName: merchant.MerchantName,
		BankAccount:  merchant.BankAccount,
	}
}

func ToTransactionListResponse(transactions []domain.Transaction) []web.TransactionResponseList {
	transactionListResponse := []web.TransactionResponseList{}
	for _, transaction := range transactions {
		transactionListResponse = append(transactionListResponse, web.TransactionResponseList{
			Id:                transaction.Id,
			MerchantName:      transaction.MerchantName,
			CustomerName:      transaction.FullName,
			OTR:               transaction.OTR,
			AdminFee:          transaction.AdminFee,
			AssetName:         transaction.AssetName,
			InterestAmount:    transaction.InterestAmount,
			InstallmentAmount: transaction.InstallmentAmount,
			Tenor:             transaction.Tenor,
		})
	}

	return transactionListResponse
}
