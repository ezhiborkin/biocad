package repository

import (
	"mas/pkg/er"
	"mas/pkg/unit"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	GetProcessedDataMongo
	GetProcessedFileMongo
	GetProcessedErrorMongo
}

type GetProcessedFileR interface {
	GetProcessedFileM(options *options.FindOptions) (fileNs []unit.ProcessedFile, err error)
}

type GetProcessedDataR interface {
	GetProcessedDataM(filter primitive.M, options *options.FindOptions) (uns []unit.Unit, err error)
}

type GetProcessedErrorR interface {
	GetProcessedErrorM(filter primitive.M, options *options.FindOptions) (er []er.ErrorOpenFile, err error)
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		GetProcessedDataMongo:  *NewGetProcessedDataMongo(client),
		GetProcessedFileMongo:  *NewGetProcessedFileMongo(client),
		GetProcessedErrorMongo: *NewGetProcessedErrorMongo(client),
	}
}
