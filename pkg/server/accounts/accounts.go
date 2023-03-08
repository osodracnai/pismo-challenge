package accounts

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func New() *Accounts {
	return &Accounts{}
}

type Accounts struct {
}

func (a *Accounts) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func (a *Accounts) GetById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
