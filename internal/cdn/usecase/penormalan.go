package usecase

import (
	"github.com/SuperP2TL/Backend/presentation"
	"mime/multipart"
)

func (uc *Usecase) StorePenormalanPhotoFiles(fpSebelum, fpSesudah *multipart.FileHeader, id string) error {
	return uc.repositories.StorePenormalanFile(fpSebelum, fpSesudah, id)
}

func (uc *Usecase) CreateSinglePenormalanReports(input presentation.CreatePenormalanReportsRequest) (insertedID []int, err error) {
	return uc.repositories.CreatePenormalanReports([]presentation.CreatePenormalanReportsRequest{input})
}
