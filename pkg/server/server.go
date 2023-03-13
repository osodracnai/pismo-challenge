package server

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/osodracnai/pismo-challenge/pkg/server/accounts"
	"github.com/osodracnai/pismo-challenge/pkg/server/transactions"
	"github.com/sirupsen/logrus"
)

type Server struct {
	validate     *validator.Validate
	accounts     *accounts.Accounts
	transactions *transactions.Transactions
}

// New is method to get a new server instance
func New(accounts *accounts.Accounts, transactions *transactions.Transactions) (*Server, error) {
	s := Server{
		accounts:     accounts,
		transactions: transactions,
		validate:     validator.New(),
	}
	return &s, nil
}

// Create New Engine
func (s *Server) NewEngine(enableMetrics bool) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	if logrus.GetLevel() == logrus.DebugLevel {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	if enableMetrics {
		p := ginprom.New(
			ginprom.Engine(r),
			ginprom.Path("/metrics"),
		)
		r.Use(p.Instrument())
	}

	r.Use(gin.Recovery())

	r.POST("/accounts", s.accounts.Create)
	r.GET("/accounts/:accountId", s.accounts.GetById)
	r.POST("/transactions", s.transactions.Create)

	return r
}
