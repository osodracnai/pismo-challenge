package transactions

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/osodracnai/pismo-challenge/pkg/database"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Transactions struct {
	db TransactionDatabase
}
type TransactionDatabase interface {
	InsertTransaction(ctx context.Context, account database.Transaction) error
	GetTransactionByID(ctx context.Context, id string) (*database.Transaction, error)
}

type CreateRequest struct {
	AccountId       string  `json:"account_id"`
	OperationTypeId int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

func New(db TransactionDatabase) *Transactions {
	return &Transactions{db: db}
}

type Accounts struct {
	db TransactionDatabase
}

func (a *Transactions) Create(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "CreateTransaction")
	defer span.Finish()
	var req CreateRequest
	logrus.Debugln("Binding create account request")
	if err := c.BindJSON(&req); err != nil {
		ext.Error.Set(span, true)
		span.SetTag("error", fmt.Sprintf("binding create transaction request: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accId, err := gocql.ParseUUID(req.AccountId)
	if err != nil {
		ext.Error.Set(span, true)
		span.SetTag("error", fmt.Sprintf("binding create transaction request: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = a.db.InsertTransaction(ctx, database.Transaction{
		TransactionId:   gocql.MustRandomUUID(),
		AccountId:       accId,
		OperationTypeId: req.OperationTypeId,
		Amount:          int(req.Amount * 100),
		EventDate:       time.Now(),
	})
	if err != nil {
		ext.Error.Set(span, true)
		span.SetTag("error", fmt.Sprintf("insert transaction error: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
