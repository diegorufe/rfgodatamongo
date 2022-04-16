package constants

const (
	NameFieldPkInMongoCollection    string = "_id"         // Nombre del campo de la pk en mongo
	NameFieldPKInGolangStruct       string = "ID"          // Nombre del campo de la pk en la estrcutura de golang
	CountPipelineAggregate          string = "$count"      // Count pipilene aggregate
	NameFieldCountPipelineAggregate string = "query_count" // Nombre del campo a obtener después de contar
)
