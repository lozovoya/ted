package model

type Account struct {
	ID       string `json:"id,omitempty"`
	Owner    string `json:"owner,omitempty"`
	Balance  int    `json:"balance,omitempty"`
	IsActive bool   `json:"is_active"`
}
