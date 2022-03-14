package dto

type AccountExistanceDTO struct {
	Exist bool `json:"exist"`
}

type AccountBalanceDTO struct {
	Balance int `json:"balance"`
}

type TransactionsReqDTO struct {
	Account string `json:"account"`
}
