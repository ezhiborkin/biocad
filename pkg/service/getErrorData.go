package service

import (
	"mas/pkg/er"
	"mas/pkg/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetProcessedErrorService struct {
	repos repository.GetProcessedErrorMongo
}

func NewGetProcessedErrorService(repos repository.GetProcessedErrorMongo) *GetProcessedErrorService {
	return &GetProcessedErrorService{repos: repos}
}

func (s *Service) GetProcessedErrorS(filter primitive.M, options *options.FindOptions) (er []er.ErrorOpenFile, err error) {
	return s.repos.GetProcessedErrorMongo.GetProcessedErrorM(filter, options)
}
