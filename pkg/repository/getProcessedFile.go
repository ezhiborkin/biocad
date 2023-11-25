package repository

import (
	"context"
	"log"
	"mas/pkg/unit"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetProcessedFileMongo struct {
	client *mongo.Client
}

func NewGetProcessedFileMongo(client *mongo.Client) *GetProcessedFileMongo {
	return &GetProcessedFileMongo{client: client}
}

func (r *GetProcessedFileMongo) GetProcessedFileM(options *options.FindOptions) (fileNs []unit.ProcessedFile, err error) {
	cursor, err := r.client.Database(viper.GetString("db.dbname")).Collection(viper.GetString("collectionName.processedFiles")).Find(context.Background(), bson.M{}, options)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var processedFiles []unit.ProcessedFile
	for cursor.Next(context.Background()) {
		var file unit.ProcessedFile
		if err := cursor.Decode(&file); err != nil {
			log.Println(err)
			return nil, err
		}
		processedFiles = append(processedFiles, file)
	}

	return processedFiles, nil
}
