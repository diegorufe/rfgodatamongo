package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BaseAuditIntPkEntity entidad base de auditoria para todas la entidades de gorm con pk
type BaseAuditIntPkEntity struct {
	BaseAuditEntity
	ID primitive.ObjectID `bson:"_id" json:"id"`
}
