package dao

import (
	"context"
	"rfgodata/beans/query/data"
	"rfgodata/definitions"
	"rfgodata/utils"
	"rfgodatamongo/beans"
	"rfgodatamongo/constants"
	"rfgodatamongo/entity"
	"rfgodatamongo/transaction"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// BaseMongoDao estructura base para los daos de mongo (bases de datos no relacionales)
//
// Ver la interfaz IBaseDao para la documentación de cabecera de los métodos
type BaseMongoDao struct {
}

func (daoMongo BaseMongoDao) Edit(mapParams *map[string]interface{}, commonsParametersSendFindData definitions.ICommonsParametersSendFindData, data interface{}) (interface{}, error) {
	var returnData interface{} = nil
	var returnErrorEdit error = nil

	if data != nil {

		// Añadimos campo de actualizado a fecha de
		utils.AddUpdatedAt(data)

		// Buscamos la transacción
		transaction, returnError := utils.GetTransactionInParams(mapParams)

		// En el caso de no tener error a la hora de buscar la transación guardamos los datos
		if returnError == nil {

			transactionFactory := transaction.TransactionFactory()

			// Guardamos los datos con la transación
			collection := transactionFactory.(*mongo.Database).Collection(data.(entity.Tabler).TableName())
			_, err := collection.UpdateOne(context.TODO(), bson.M{constants.NameFieldPkInMongoCollection: utils.GetPkValue(data, constants.NameFieldPKInGolangStruct)}, data)
			returnErrorEdit = err

			// En el caso de que se haya guardado de forma correcta guardamos
			if err == nil {
				returnData = data

			}

		} else {
			returnErrorEdit = returnError
		}

	} else {
		// TODO poner un mensajes de error más especifico
		returnData = constants.ErrInvalidData
	}

	return returnData, returnErrorEdit
}

func (daoMongo BaseMongoDao) Add(mapParams *map[string]interface{}, commonsParametersSendFindData definitions.ICommonsParametersSendFindData, data interface{}) (interface{}, error) {
	var returnData interface{} = nil
	var returnErrorMethod error = nil

	if data != nil {

		// Añadimos fecha de creación
		utils.AddCreatedAt(data)

		// Añadimos campo de actualizado a fecha de
		utils.AddUpdatedAt(data)

		// Buscamos la transacción
		transaction, returnError := utils.GetTransactionInParams(mapParams)

		// En el caso de no tener error a la hora de buscar la transación guardamos los datos
		if returnError == nil {

			transactionFactory := transaction.TransactionFactory()

			// Guardamos los datos con la transación
			collection := transactionFactory.(*mongo.Database).Collection(data.(entity.Tabler).TableName())
			result, err := collection.InsertOne(context.TODO(), data)
			returnErrorMethod = err

			// En el caso de que se haya guardado de forma correcta guardamos
			if err == nil {

				if result.InsertedID != nil {
					utils.AddPkValue(data, result.InsertedID, constants.NameFieldPKInGolangStruct)
				}

				returnData = data
			}

		} else {
			returnErrorMethod = returnError
		}

	} else {
		// TODO poner un mensajes de error más especifico
		returnErrorMethod = constants.ErrInvalidData
	}

	return returnData, returnErrorMethod
}

func (daoMongo BaseMongoDao) Delete(mapParams *map[string]interface{}, commonsParametersSendFindData definitions.ICommonsParametersSendFindData, data interface{}) (bool, error) {
	var returnData bool = false
	var returnErrorMethod error = nil

	if data != nil {

		// Buscamos la transacción
		transaction, returnError := utils.GetTransactionInParams(mapParams)

		// En el caso de no tener error a la hora de buscar la transación guardamos los datos
		if returnError == nil {

			transactionFactory := transaction.TransactionFactory()

			// Guardamos los datos con la transación
			collection := transactionFactory.(*mongo.Database).Collection(data.(entity.Tabler).TableName())
			_, err := collection.DeleteOne(context.TODO(), bson.M{constants.NameFieldPkInMongoCollection: utils.GetPkValue(data, constants.NameFieldPKInGolangStruct)})
			returnErrorMethod = err

			// En el caso de que se haya guardado de forma correcta guardamos
			if err == nil {

				returnData = true
			}

		} else {
			returnErrorMethod = returnError
		}

	} else {
		// TODO poner un mensajes de error más especifico
		returnErrorMethod = constants.ErrInvalidData
	}

	return returnData, returnErrorMethod
}

func (daoMongo BaseMongoDao) Count(mapParams *map[string]interface{}, commonsParametersSendFindData definitions.ICommonsParametersSendFindData, tableName string, filtersClauses []data.FilterClause, joinsClauses []data.JoinClause, groupsClauses []data.GroupClause) (int64, error) {
	var returnData int64 = 0
	var returnErrorMethod error = nil

	transaction, returnError := utils.GetTransactionInParams(mapParams)

	if returnError == nil {

		transactionFactory := transaction.TransactionFactory()
		collection := transactionFactory.(*mongo.Database).Collection(tableName)

		countPipeline := bson.D{{constants.CountPipelineAggregate, constants.NameFieldCountPipelineAggregate}}

		transactionMongo := castTransactionFactoryToMongoTransaction(transaction)

		cursor, errAggregate := executeAggregate(transactionMongo, collection, mongo.Pipeline{countPipeline})

		if errAggregate == nil {

			var countResult []beans.CountResult

			if err := cursor.All(transactionMongo.Ctx, &countResult); err != nil {
				returnErrorMethod = err
			} else if len(countResult) == 1 {
				returnData = countResult[0].QueryCount
			}

		} else {
			returnErrorMethod = errAggregate
		}

	} else {
		returnErrorMethod = returnError
	}

	return returnData, returnErrorMethod

}

// Método para castear la factoria de la transación a una transación mongo
//
// @transactionToCast a castear
//
// @returns la transación de mongo ya casteada a la estrucutra esperada
func castTransactionFactoryToMongoTransaction(transactionToCast interface{}) transaction.TransactionMongo {
	return transactionToCast.(transaction.TransactionMongo)
}

// Método para ejecutar la acción de agregado de mongo
//
// @param transactionMongo donde tenemos la transación de mongo para obtener por ejemplo el contexto
//
// @param collection donde ejecutar el agregado
//
// @pipelie a ejecutar en el agregado. Ejemplo $count
//
// @returns
//
// - En el caso de que todo sea correcto un cursor de mongo con el resultado del agregado
//
// - En el caso de existir un error devolverá el error producido
func executeAggregate(transactionMongo transaction.TransactionMongo, collection *mongo.Collection, pipelie mongo.Pipeline) (*mongo.Cursor, error) {
	cursor, err := collection.Aggregate(transactionMongo.Ctx, pipelie)
	return cursor, err
}
