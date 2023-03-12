package database

import (
	"clondy/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

// ProjectCollector ...
type ProjectCollector interface {
	InsertProjec(document interface{}) (*mongo.InsertOneResult, error)
	GetProjects(filter interface{}) ([]*model.Project, error)
	GetProject(filter interface{}) (*model.Project, error)
	UpdateProject(filter interface{}, update interface{}) (*mongo.UpdateResult, error)
	DeleteProject(filter interface{}) (*mongo.DeleteResult, error)
	CountProject(filter interface{}) (int64, error)
}

// InsertProjec - insert one user
func (db *Database) InsertProjec(document interface{}) (*mongo.InsertOneResult, error) {
	collection := db.DB.Collection(os.Getenv("C_PROJECTS"))
	return collection.InsertOne(context.Background(), document)
}

// GetProjects - get all user with confdition
func (db *Database) GetProjects(filter interface{}) ([]*model.Project, error) {
	collection := db.DB.Collection(os.Getenv("C_PROJECTS"))
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())
	if err != nil {
		return nil, err
	}

	results := []*model.Project{}
	for cur.Next(context.Background()) {
		var result *model.Project
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

// GetProject - get one user by id
func (db *Database) GetProject(filter interface{}) (*model.Project, error) {
	collection := db.DB.Collection(os.Getenv("C_PROJECTS"))

	var result *model.Project

	return result, collection.FindOne(context.Background(), filter).Decode(&result)
}

// UpdateProject - update user detail
func (db *Database) UpdateProject(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {

	collection := db.DB.Collection(os.Getenv("C_PROJECTS"))
	return collection.UpdateOne(context.Background(), filter, update)

}

// DeleteProject - delete one user
func (db *Database) DeleteProject(filter interface{}) (*mongo.DeleteResult, error) {
	collection := db.DB.Collection(os.Getenv("C_PROJECTS"))
	return collection.DeleteOne(context.Background(), filter)

}

// CountProject - count user with condition
func (db *Database) CountProject(filter interface{}) (int64, error) {
	collection := db.DB.Collection(os.Getenv("C_PROJECTS"))
	return collection.CountDocuments(context.Background(), filter)
}
