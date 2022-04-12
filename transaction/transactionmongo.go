package transaction

import (
	"context"
	"rfgodata/constants/core"

	"go.mongodb.org/mongo-driver/mongo"
)

// TransactionMongo extructura para poder realizar transaciones de mongo
//
// Revisar la clase de definición ITransaction
type TransactionMongo struct {
	Transaction *mongo.Database // Transación de gorm para poder realizar consultas
}

func (transactionMongo TransactionMongo) RollBack() error {
	return nil
}

func (transactionMongo TransactionMongo) Commit() error {
	return nil
}

func (transactionMongo TransactionMongo) FinishTransaction(err error) error {
	var errReturn error = err
	return errReturn
}

func (transactionMongo TransactionMongo) TransactionFactory() interface{} {
	return transactionMongo.Transaction
}

func (transactionMongo TransactionMongo) StartTransactionContext(dbConnection interface{}, ctx context.Context, mapParamsService *map[string]interface{}) {
	var db *mongo.Database = dbConnection.(*mongo.Database)
	transactionMongo.Transaction = db
	(*mapParamsService)[core.ParamTransaction] = transactionMongo
}

func (transactionMongo TransactionMongo) FinishTransactionContext(err error, mapParamsService *map[string]interface{}) error {
	var returnError error
	var transactionError error = transactionMongo.FinishTransaction(err)

	if err != nil {
		returnError = err
	} else {
		returnError = transactionError
	}
	return returnError
}
