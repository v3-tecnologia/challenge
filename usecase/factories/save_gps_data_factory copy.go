package factory

import (
	"github.com/mkafonso/go-cloud-challenge/repository"
	"github.com/mkafonso/go-cloud-challenge/usecase"
)

func SaveGPSDataFactory(repo repository.GPSRepositoryInterface) *usecase.SaveGPSData {
	usecase := usecase.NewSaveGPSData(repo)
	return usecase
}
