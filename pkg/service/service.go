package service

import (
	"mas/pkg/er"
	"mas/pkg/repository"
	"mas/pkg/unit"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetProcessedDataServ interface {
	GetProcessedDataS(filter primitive.M, options *options.FindOptions) (units []unit.Unit, err error)
}

type GetProcessedFileServ interface {
	GetProcessedFileS(options *options.FindOptions) (fileNs []unit.ProcessedFile, err error)
}

type GetProcessedErrorServ interface {
	GetProcessedErrorS(options *options.FindOptions) (er []er.ErrorOpenFile, err error)
}

type Service struct {
	repos *repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{repos: repos}
}
