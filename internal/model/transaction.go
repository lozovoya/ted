package model

import "time"

type Transaction struct {
	ID     string    `json:"id,omitempty"`
	Source string    `json:"source,omitempty"`
	Type   string    `json:"type,omitempty"`
	Dest   string    `json:"dest,omitempty"`
	Amount int       `json:"amount,omitempty"`
	Time   time.Time `json:"time,omitempty"`
}
