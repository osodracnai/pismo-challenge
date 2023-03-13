package accounts

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
	"strconv"
)

type Database interface {
	InsertAccount(ctx context.Context, account database.Account) error
	GetAccountByID(ctx context.Context, id string) (*database.Account, error)
}

type CreateRequest struct {
	DocumentNumber string `json:"document_number"`
}

func New(db Database) *Accounts {
	return &Accounts{db: db}
}

type Accounts struct {
	db Database
}

func (a *Accounts) Create(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "CreateAccount")
	defer span.Finish()
	var req CreateRequest
	logrus.Debugln("Binding create account request")
	if err := c.BindJSON(&req); err != nil {
		ext.Error.Set(span, true)
		span.SetTag("error", fmt.Sprintf("binding create account request: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	docNumber, err := strconv.Atoi(req.DocumentNumber)
	if err != nil {
		ext.Error.Set(span, true)
		span.SetTag("error", fmt.Sprintf("binding create account request: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = a.db.InsertAccount(ctx, database.Account{
		AccountId:      gocql.MustRandomUUID(),
		DocumentNumber: docNumber,
	})
	if err != nil {
		ext.Error.Set(span, true)
		span.SetTag("error", fmt.Sprintf("insert account error: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (a *Accounts) GetById(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "GetAccountById")
	defer span.Finish()
	id := c.Param("accountId")
	account, err := a.db.GetAccountByID(ctx, id)
	if err != nil {
		ext.Error.Set(span, true)
		span.SetTag("error", fmt.Sprintf("insert account error: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}
