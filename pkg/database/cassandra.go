package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/mitchellh/mapstructure"
	"github.com/opentracing/opentracing-go"
	"github.com/osodracnai/pismo-challenge/cmd/config"
)

type Cassandra struct {
	Session  *gocql.Session
	keyspace string
}

const (
	AccountsTable     = "accounts"
	TransactionsTable = "transactions"
)

func NewCassandra(cassandra config.Cassandra) (*Cassandra, error) {
	if cassandra.Keyspace == "" {
		return nil, errors.New("cassandra keyspace must be declared")
	}
	cluster := gocql.NewCluster(cassandra.Hosts...)

	queryObserver := QueryTracer{}
	batchObserver := BatchTracer{}
	cluster.Port = cassandra.Port
	cluster.Timeout = cassandra.Timeout
	cluster.ConnectTimeout = cassandra.ConnectTimeout
	cluster.CQLVersion = cassandra.Version
	cluster.Keyspace = cassandra.Keyspace
	cluster.Consistency = gocql.Quorum
	cluster.QueryObserver = &queryObserver
	cluster.BatchObserver = &batchObserver

	if cassandra.Username != "" && cassandra.Password != "" {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: cassandra.Username,
			Password: cassandra.Password,
		}
	}

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return &Cassandra{
		Session:  session,
		keyspace: cassandra.Keyspace,
	}, nil
}

func (db *Cassandra) GetTransactionByID(ctx context.Context, id string) (*Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetTransactionByID")
	defer span.Finish()
	args := []interface{}{id}
	query := fmt.Sprintf("SELECT * FROM %s.%s WHERE transaction_id = ? ", db.keyspace, TransactionsTable)

	var newTransaction Transaction
	mapValues := make(map[string]interface{})
	if err := db.Session.Query(query, args...).WithContext(ctx).Consistency(gocql.Quorum).MapScan(mapValues); err != nil {
		return nil, err
	}
	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &newTransaction,
		TagName:  "json",
	}
	decoder, _ := mapstructure.NewDecoder(cfg)
	err := decoder.Decode(mapValues)
	if err != nil {
		return nil, err
	}

	return &newTransaction, nil
}

func (db *Cassandra) InsertTransaction(ctx context.Context, transaction Transaction) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetTransactionByID")
	defer span.Finish()
	queryString := `insert into pismo.transaction(transaction_id, event_date, account_id, amount, operation_type_id) VALUES (?,?,?,?,?);`

	return db.Session.Query(queryString, transaction.TransactionId, transaction.EventDate, transaction.AccountId, transaction.Amount, transaction.OperationTypeId).WithContext(ctx).Exec()
}

func (db *Cassandra) InsertAccount(ctx context.Context, account Account) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetTransactionByID")
	defer span.Finish()
	queryString := `insert into pismo.accounts(account_id, document_number) VALUES (?,?);`

	return db.Session.Query(queryString, account.AccountId, account.DocumentNumber).WithContext(ctx).Exec()
}
func (db *Cassandra) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetAccountByID")
	defer span.Finish()
	args := []interface{}{id}
	query := fmt.Sprintf("SELECT * FROM %s.%s WHERE account_id = ? ", db.keyspace, AccountsTable)

	var account Account
	mapValues := make(map[string]interface{})
	if err := db.Session.Query(query, args...).WithContext(ctx).Consistency(gocql.Quorum).MapScan(mapValues); err != nil {
		return nil, err
	}
	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &account,
		TagName:  "json",
	}
	decoder, _ := mapstructure.NewDecoder(cfg)
	err := decoder.Decode(mapValues)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
