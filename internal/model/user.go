package model

type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	IsActive bool   `json:"is_active,omitempty"`
}
