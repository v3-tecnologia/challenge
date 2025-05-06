package factory

import (
	"github.com/mkafonso/go-cloud-challenge/repository"
	"github.com/mkafonso/go-cloud-challenge/usecase"
)

func SaveGyroscopeDataFactory(repo repository.GyroscopeRepositoryInterface) *usecase.SaveGyroscopeData {
	usecase := usecase.NewSaveGyroscopeData(repo)
	return usecase
}
