package service

import (
	"mas/pkg/repository"
	"mas/pkg/unit"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetProcessedFileService struct {
	repos repository.GetProcessedFileMongo
}

func NewGetProcessedFileService(repos repository.GetProcessedFileMongo) *GetProcessedFileService {
	return &GetProcessedFileService{repos: repos}
}

func (s *Service) GetProcessedFileS(options *options.FindOptions) (fileNs []unit.ProcessedFile, err error) {
	return s.repos.GetProcessedFileMongo.GetProcessedFileM(options)
}
