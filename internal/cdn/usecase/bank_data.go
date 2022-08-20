package usecase

import (
	"github.com/SuperP2TL/Backend/presentation"
	"mime/multipart"
)

func (uc *Usecase) StoreBankDataFile(file *multipart.FileHeader, category string, filename string) error {
	return uc.repositories.StoreBankDataFile(file, category, filename)
}

func (uc *Usecase) GetBankDataCategoryList() (category []string, err error) {
	return uc.repositories.GetBankDataCategoryList()
}
func (uc *Usecase) GetBankDataFileList(category string, pagination *presentation.Pagination) (filenames []string, err error) {
	return uc.repositories.GetBankDataFileList(category, pagination)
}

func (uc *Usecase) GetBankDataList() (res []presentation.GetBankDataListResponse, err error) {
	categories, err := uc.repositories.GetBankDataCategoryList()
	if err != nil {
		return nil, err
	}

	for _, cat := range categories {
		_t := presentation.GetBankDataListResponse{
			Category: cat,
		}

		filelist, err := uc.GetBankDataFileList(cat, nil)
		if err != nil {
			return nil, err
		}

		_t.Files = []presentation.GetBankDataListResponseFile{}

		for _, file := range filelist {
			_t.Files = append(_t.Files, presentation.GetBankDataListResponseFile{
				Name: file,
			})
		}

		res = append(res, _t)
	}

	return res, nil
}
