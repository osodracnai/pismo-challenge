package server

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/osodracnai/pismo-challenge/mocks"
	"github.com/osodracnai/pismo-challenge/pkg/server/accounts"
	"github.com/osodracnai/pismo-challenge/pkg/server/transactions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccounts(t *testing.T) {
	t.Run("Get Account with no error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		accMock := mocks.NewMockDatabase(ctrl)
		server, _ := New(accounts.New(accMock), &transactions.Transactions{})
		router := server.NewEngine(false)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/accounts/1234", nil)

		accMock.EXPECT().GetAccountByID(gomock.Any(), gomock.Any()).Return(nil, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
	t.Run("Get Account with error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		accMock := mocks.NewMockDatabase(ctrl)
		server, _ := New(accounts.New(accMock), &transactions.Transactions{})
		router := server.NewEngine(false)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/accounts/1234", nil)

		accMock.EXPECT().GetAccountByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("test error"))
		router.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
	})

}
