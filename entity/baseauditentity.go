package entity

import (
	"gopkg.in/guregu/null.v4"
)

// BaseAuditEntity entidad base de auditoria para todas la entidades de gorm
type BaseAuditEntity struct {
	CreatedAt null.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt null.Time `bson:"updatedAt" json:"updatedAt"`
}

// Tabler implementar para indicar el nombre de la tabla de la entidad
type Tabler interface {
	TableName() string
}
