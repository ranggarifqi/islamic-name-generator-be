package name

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_NAME = "names"

type mongoRepository struct {
	ctx context.Context
	db  *mongo.Database
}

func NewMongoRepository(ctx context.Context, db *mongo.Database) INameRepository {
	return &mongoRepository{
		ctx,
		db,
	}
}

// TODO: Integration Test

func (r *mongoRepository) getCollection() *mongo.Collection {
	return r.db.Collection((COLLECTION_NAME))
}

func (r *mongoRepository) findBy(filter FindByFilter) (*[]Name, error) {
	collection := r.getCollection()

	whereOpt := bson.M{}

	if filter.Name != "" {
		whereOpt["name"] = filter.Name
	}

	if filter.Gender != "" {
		whereOpt["gender"] = filter.Gender
	}

	if len(filter.NameTypes) > 0 {
		whereOpt["nameTypes"] = bson.M{
			"$all": filter.NameTypes,
		}
	}

	cursor, err := collection.Find(r.ctx, whereOpt)
	if err != nil {
		return nil, err
	}

	var result []Name

	err = cursor.All(r.ctx, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *mongoRepository) findById(id string) (*Name, error) {
	collection := r.getCollection()

	var result *Name

	err := collection.FindOne(r.ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *mongoRepository) create(payload Name) (*Name, error) {
	collection := r.getCollection()

	insertResult, err := collection.InsertOne(r.ctx, payload)
	if err != nil {
		return nil, err
	}

	newData, err := r.findById(insertResult.InsertedID.(string))
	if err != nil {
		return nil, err
	}

	return newData, nil
}

func (r *mongoRepository) updateById(id string, payload Name) (*Name, error) {
	collection := r.getCollection()

	_, err := collection.UpdateByID(r.ctx, id, payload)
	if err != nil {
		return nil, err
	}

	updatedData, err := r.findById(id)
	if err != nil {
		return nil, err
	}

	return updatedData, nil
}
