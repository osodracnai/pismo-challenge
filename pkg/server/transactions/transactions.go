package transactions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func New() *Transactions {
	return &Transactions{}
}

type Transactions struct {
}

func (a *Transactions) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
