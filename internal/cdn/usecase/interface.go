package usecase

import (
	"github.com/SuperP2TL/Backend/presentation"
	"mime/multipart"
)

type BankDataFilopsRepo interface {
	StoreBankDataFile(file *multipart.FileHeader, category string, filename string) error
	GetBankDataCategoryList() (category []string, err error)
	GetBankDataFileList(category string, pagination *presentation.Pagination) (filenames []string, err error)

	GetTemuanBaseDirectory() string
}

type PenormalanFilopsRepo interface {
	StorePenormalanFile(fpSebelum, fpSesudah *multipart.FileHeader, id string) error
	StoreTemuanFile(fpBA, fpLokasi *multipart.FileHeader, id string) error
}

type ReportRepo interface {
	CreateFindingReports(input []presentation.CreateFindingReportsRequest) (lastInsertedID int, err error)
	CreatePenormalanReports(input []presentation.CreatePenormalanReportsRequest) (insertedID []int, err error)
}
