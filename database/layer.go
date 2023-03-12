package database

import (
	"clondy/model"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

// LayerCollector ...
type LayerCollector interface {
	InsertLayer(document interface{}) (*mongo.InsertOneResult, error)
	GetLayers(filter interface{}) ([]*model.Layer, error)
	GetLayer(filter interface{}) (*model.Layer, error)
	UpdateLayer(filter interface{}, update interface{}) (*mongo.UpdateResult, error)
	DeleteLayers(filter interface{}) (*mongo.DeleteResult, error)
	DeleteLayer(filter interface{}) (*mongo.DeleteResult, error)
	CountLayers(filter interface{}) (int64, error)
}

// InsertLayer - insert one user
func (db *Database) InsertLayer(document interface{}) (*mongo.InsertOneResult, error) {
	collection := db.DB.Collection(os.Getenv("C_LAYER"))
	return collection.InsertOne(context.Background(), document)
}

// GetLayers - get all user with confdition
func (db *Database) GetLayers(filter interface{}) ([]*model.Layer, error) {
	collection := db.DB.Collection(os.Getenv("C_LAYER"))
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())
	if err != nil {
		return nil, err
	}

	results := []*model.Layer{}
	for cur.Next(context.Background()) {
		var result *model.Layer
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

// GetLayer - get one user by id
func (db *Database) GetLayer(filter interface{}) (*model.Layer, error) {
	collection := db.DB.Collection(os.Getenv("C_LAYER"))

	var result *model.Layer
	return result, collection.FindOne(context.Background(), filter).Decode(&result)
}

// UpdateLayer - update user detail
func (db *Database) UpdateLayer(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	collection := db.DB.Collection(os.Getenv("C_LAYER"))
	return collection.UpdateOne(context.Background(), filter, update)

}

// DeleteLayer - delete one user
func (db *Database) DeleteLayer(filter interface{}) (*mongo.DeleteResult, error) {
	collection := db.DB.Collection(os.Getenv("C_LAYER"))
	return collection.DeleteOne(context.Background(), filter)

}

// DeleteLayers - delete one user
func (db *Database) DeleteLayers(filter interface{}) (*mongo.DeleteResult, error) {
	collection := db.DB.Collection(os.Getenv("C_LAYER"))
	return collection.DeleteMany(context.Background(), filter)

}

// CountLayers - count user with condition
func (db *Database) CountLayers(filter interface{}) (int64, error) {
	collection := db.DB.Collection(os.Getenv("C_LAYER"))
	return collection.CountDocuments(context.Background(), filter)
}
