package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Layer ..
type Layer struct {
	ID        primitive.ObjectID     `bson:"_id" json:"_id"`
	ProjectID string                 `json:"project_id" bson:"project_id" `
	Name      string                 `json:"name" bson:"name"`
	Data      map[string]interface{} `json:"data" bson:"data"`
	Type      string                 `json:"type" bson:"type"`
	Style     map[string]interface{} `json:"style" bson:"style"`
}

// Data ..
type Data struct {
	Features Features `json:"features" bson:"features"`
}

// Features ..
type Features []struct {
	Type       string     `json:"type_features" bson:"type_features"`
	Properties properties `json:"properties" bson:"properties"`
	Geometry   geometry   `json:"geometry" bson:"geometry"`
}

type properties interface {
}

type geometry interface {
}

type style interface {
}

// // L ...
// type L struct {
// 	ID        primitive.ObjectID
// 	Name      string
// 	Geometry  map[string]interface{}
// 	ProjectID primitive.ObjectID
// }
