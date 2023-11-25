package repository

import (
	"context"
	"log"
	"mas/pkg/unit"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetProcessedDataMongo struct {
	client *mongo.Client
}

func NewGetProcessedDataMongo(client *mongo.Client) *GetProcessedDataMongo {
	return &GetProcessedDataMongo{client: client}
}

func (r *GetProcessedDataMongo) GetProcessedDataM(filter primitive.M, options *options.FindOptions) (uns []unit.Unit, err error) {
	cursor, err := r.client.Database(viper.GetString("db.dbname")).Collection(viper.GetString("collectionName.processedData")).Find(context.Background(), filter, options)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var units []unit.Unit
	for cursor.Next(context.Background()) {
		var unit unit.Unit
		if err := cursor.Decode(&unit); err != nil {
			log.Println(err)
			return nil, err
		}
		units = append(units, unit)
	}

	return units, nil
}
