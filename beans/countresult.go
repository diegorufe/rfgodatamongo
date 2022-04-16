package beans

// CountResult estructura para guardar el resultado de conta
type CountResult struct {
	QueryCount int64 `bson:"query_count"` // Campo donde tendremos el resultado de contar el número de registro
}
