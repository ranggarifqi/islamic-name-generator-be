package name

import (
	"context"

	"github.com/ranggarifqi/islamic-name-generator-be/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *mongoRepository) FindBy(filter FindByFilter) (*[]Name, error) {
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

func (r *mongoRepository) FindById(id string) (*Name, error) {
	collection := r.getCollection()

	var result *Name

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(r.ctx, bson.M{"_id": objectID}).Decode(&result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *mongoRepository) Create(payload Name) (*Name, error) {
	collection := r.getCollection()

	insertResult, err := collection.InsertOne(r.ctx, payload)
	if err != nil {
		return nil, err
	}

	idStr := insertResult.InsertedID.(primitive.ObjectID).Hex()

	newData, err := r.FindById(idStr)
	if err != nil {
		return nil, err
	}

	return newData, nil
}

func (r *mongoRepository) UpdateById(id string, payload Name) (*Name, error) {
	collection := r.getCollection()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	bsonMap, err := mongodb.StructToBsonM(payload)
	if err != nil {
		return nil, err
	}

	_, err = collection.UpdateByID(r.ctx, objectID, bson.M{
		"$set": *bsonMap,
	})
	if err != nil {
		return nil, err
	}

	updatedData, err := r.FindById(id)
	if err != nil {
		return nil, err
	}

	return updatedData, nil
}
