package test

import (
	"context"
	"fmt"
	"rfgodatamongo/dao"
	"rfgodatamongo/transaction"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection URI

// TestCount método para prueba de conteo
func TestCount(t *testing.T) {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Uri))

	if err == nil {

		// Ping the primary
		if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
			panic(err)
		}

		db := client.Database(NameDatabaseTest)
		var mapParams *map[string]interface{} = &map[string]interface{}{}
		var tx *transaction.TransactionMongo = &transaction.TransactionMongo{}

		tx.StartTransactionContext(db, context.TODO(), mapParams)

		mongoDao := &dao.BaseMongoDao{}

		responseCount, err := mongoDao.Count(mapParams, nil, CollectionDatabaseTest, nil, nil, nil)

		if err == nil {
			fmt.Printf("Response count %d ", responseCount)
		} else {
			panic(err)
		}

		tx.FinishTransaction(nil)

	} else {
		panic(err)
	}
}
