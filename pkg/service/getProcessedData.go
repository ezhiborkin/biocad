package service

import (
	"mas/pkg/repository"
	"mas/pkg/unit"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetProcessedDataService struct {
	repos repository.GetProcessedDataMongo
}

func NewGetProcessedDataService(repos repository.GetProcessedDataMongo) *GetProcessedDataService {
	return &GetProcessedDataService{repos: repos}
}

func (s *Service) GetProcessedDataS(filter primitive.M, options *options.FindOptions) (units []unit.Unit, err error) {
	return s.repos.GetProcessedDataMongo.GetProcessedDataM(filter, options)
}
