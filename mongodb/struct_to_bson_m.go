package mongodb

import "go.mongodb.org/mongo-driver/bson"

func StructToBsonM(payload interface{}) (*bson.M, error) {
	marshal, err := bson.Marshal(payload)
	if err != nil {
		return nil, err
	}

	var bsonMap bson.M

	err = bson.Unmarshal(marshal, &bsonMap)
	if err != nil {
		return nil, err
	}

	return &bsonMap, nil
}
