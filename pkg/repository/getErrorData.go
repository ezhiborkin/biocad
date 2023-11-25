package repository

import (
	"context"
	"log"
	"mas/pkg/er"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetProcessedErrorMongo struct {
	client *mongo.Client
}

func NewGetProcessedErrorMongo(client *mongo.Client) *GetProcessedErrorMongo {
	return &GetProcessedErrorMongo{client: client}
}

func (r *GetProcessedErrorMongo) GetProcessedErrorM(filter primitive.M, options *options.FindOptions) (ers []er.ErrorOpenFile, err error) {
	cursor, err := r.client.Database(viper.GetString("db.dbname")).Collection(viper.GetString("collectionName.errors")).Find(context.Background(), filter, options)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var errors []er.ErrorOpenFile
	for cursor.Next(context.Background()) {
		var error er.ErrorOpenFile
		if err := cursor.Decode(&error); err != nil {
			log.Println(err)
			return nil, err
		}
		errors = append(errors, error)
	}

	return errors, nil
}
