package usecase

import (
	"github.com/SuperP2TL/Backend/presentation"
	"mime/multipart"
)

func (uc *Usecase) CreateSingleFindingReports(input presentation.CreateFindingReportsRequest) (lastInsertedID int, err error) {
	return uc.repositories.CreateFindingReports([]presentation.CreateFindingReportsRequest{input})
}

func (uc *Usecase) StoreTemuanPhotoFiles(fpBa, fpLokasi *multipart.FileHeader, id string) error {
	return uc.repositories.StoreTemuanFile(fpBa, fpLokasi, id)
}

func (uc *Usecase) GetTemuanBaseDirectory() string {
	return uc.repositories.GetTemuanBaseDirectory()
}
