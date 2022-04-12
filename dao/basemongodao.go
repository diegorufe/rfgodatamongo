package dao

import (
	"context"
	"rfgodata/definitions"
	"rfgodata/utils"
	"rfgodatamongo/constants"
	"rfgodatamongo/entity"

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

		transactionFactory := transaction.TransactionFactory()

		// En el caso de no tener error a la hora de buscar la transación guardamos los datos
		if returnError == nil {
			// Guardamos los datos con la transación
			collection := transactionFactory.(*mongo.Database).Collection(data.(entity.Tabler).TableName())
			_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": utils.GetPkValue(data, "Id")}, data)
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

		transactionFactory := transaction.TransactionFactory()

		// En el caso de no tener error a la hora de buscar la transación guardamos los datos
		if returnError == nil {
			// Guardamos los datos con la transación
			collection := transactionFactory.(*mongo.Database).Collection(data.(entity.Tabler).TableName())
			result, err := collection.InsertOne(context.TODO(), data)
			returnErrorMethod = err

			// En el caso de que se haya guardado de forma correcta guardamos
			if err == nil {

				if result.InsertedID != nil {
					utils.AddPkValue(data, result.InsertedID, "Id")
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

		transactionFactory := transaction.TransactionFactory()

		// En el caso de no tener error a la hora de buscar la transación guardamos los datos
		if returnError == nil {
			// Guardamos los datos con la transación
			collection := transactionFactory.(*mongo.Database).Collection(data.(entity.Tabler).TableName())
			_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": utils.GetPkValue(data, "Id")})
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
