package database

import (
	"clondy/model"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

// UserCollector ...
type UserCollector interface {
	InsertUser(document interface{}) (*mongo.InsertOneResult, error)
	GetUsers(filter interface{}) ([]*model.User, error)
	GetUser(filter interface{}) (*model.User, error)
	UpdateUser(filter interface{}, update interface{}) (*mongo.UpdateResult, error)
	DeleteUser(filter interface{}) (*mongo.DeleteResult, error)
	CountUsers(filter interface{}) (int64, error)
}

// InsertUser - insert one user
func (db *Database) InsertUser(document interface{}) (*mongo.InsertOneResult, error) {
	collection := db.DB.Collection(os.Getenv("C_USERS"))
	return collection.InsertOne(context.Background(), document)
}

// GetUsers - get all user with confdition
func (db *Database) GetUsers(filter interface{}) ([]*model.User, error) {
	collection := db.DB.Collection(os.Getenv("C_USERS"))
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())
	if err != nil {
		return nil, err
	}

	results := []*model.User{}
	for cur.Next(context.Background()) {
		var result *model.User
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

// GetUser - get one user by id
func (db *Database) GetUser(filter interface{}) (*model.User, error) {
	collection := db.DB.Collection(os.Getenv("C_USERS"))

	var result *model.User
	return result, collection.FindOne(context.Background(), filter).Decode(&result)
}

// UpdateUser - update user detail
func (db *Database) UpdateUser(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	collection := db.DB.Collection(os.Getenv("C_USERS"))
	return collection.UpdateOne(context.Background(), filter, filter)

}

// DeleteUser - delete one user
func (db *Database) DeleteUser(filter interface{}) (*mongo.DeleteResult, error) {
	collection := db.DB.Collection(os.Getenv("C_USERS"))
	return collection.DeleteOne(context.Background(), filter)

}

// CountUsers - count user with condition
func (db *Database) CountUsers(filter interface{}) (int64, error) {
	collection := db.DB.Collection(os.Getenv("C_USERS"))
	return collection.CountDocuments(context.Background(), filter)
}
