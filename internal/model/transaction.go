package model

type Transaction struct {
	ID     string `json:"id,omitempty"`
	Source string `json:"source,omitempty"`
	Dest   string `json:"dest,omitempty"`
	Amount string `json:"amount,omitempty"`
}
