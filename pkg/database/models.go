package database

import (
	"github.com/gocql/gocql"
	"time"
)

type Transaction struct {
	TransactionId   gocql.UUID `json:"transaction_id"`
	AccountId       gocql.UUID `json:"account_id"`
	OperationTypeId int        `json:"operation_type_id"`
	Amount          int        `json:"amount"`
	EventDate       time.Time  `json:"event_date"`
}

type Account struct {
	AccountId      gocql.UUID `json:"account_id"`
	DocumentNumber int        `json:"document_number"`
}
