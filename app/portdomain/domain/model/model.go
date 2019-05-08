package model

import (
	"github.com/bkatrenko/ports/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MongoDoc represents internal mongo DB document for storing ports map
type MongoDoc struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	PortID string             `json:"port_id" bson:"_port_id"`
	Doc    *proto.Port        `json:"port_doc" bson:"port_doc"`
}

// GenerateID should generate new bson ID to mongo
func (md *MongoDoc) GenerateID() {
	md.ID = primitive.NewObjectID()
}
