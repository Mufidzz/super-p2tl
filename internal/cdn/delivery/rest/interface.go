package rest

import (
	"github.com/SuperP2TL/Backend/presentation"
	"mime/multipart"
)

type BankDataUC interface {
	StoreBankDataFile(file *multipart.FileHeader, category string, filename string) error
	GetBankDataList() (res []presentation.GetBankDataListResponse, err error)

	StorePenormalanPhotoFiles(fpSebelum, fpSesudah *multipart.FileHeader, id string) error
	StoreTemuanPhotoFiles(fpBa, fpLokasi *multipart.FileHeader, id string) error

	GetTemuanBaseDirectory() string
}

type ReportUC interface {
	CreateSinglePenormalanReports(input presentation.CreatePenormalanReportsRequest) (insertedID []int, err error)
	CreateSingleFindingReports(input presentation.CreateFindingReportsRequest) (lastInsertedID int, err error)
}
