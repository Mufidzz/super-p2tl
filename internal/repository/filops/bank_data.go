package filops

import (
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/presentation"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path"
)

func (filops *Filops) StoreBankDataFile(file *multipart.FileHeader, category string, filename string) error {
	dir := path.Join(filops.BankDataDirectory, category)

	err := os.MkdirAll(dir, 0777)
	if err != nil {
		log.Println(err)

		return response.InternalError{
			Type:         "Repo",
			Name:         "Filops",
			FunctionName: "StoreBankDataFile",
			Description:  "failed running create directory",
			Trace:        err,
		}.Error()
	}

	err = filops.SaveUploadedFile(file, path.Join(dir, filename))
	if err != nil {
		log.Println(err)

		return response.InternalError{
			Type:         "Repo",
			Name:         "Filops",
			FunctionName: "StoreBankDataFile",
			Description:  "failed saving file",
			Trace:        err,
		}.Error()
	}

	return nil
}

func (filops *Filops) GetBankDataCategoryList() (category []string, err error) {
	files, err := ioutil.ReadDir(filops.BankDataDirectory)
	if err != nil {
		log.Println(err)

		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Filops",
			FunctionName: "GetBankDataCategoryList",
			Description:  "failed read folder",
			Trace:        err,
		}.Error()
	}

	for _, f := range files {
		if f.IsDir() {
			category = append(category, f.Name())
		}
	}

	return category, nil
}

func (filops *Filops) GetBankDataFileList(category string, pagination *presentation.Pagination) (filenames []string, err error) {
	files, err := ioutil.ReadDir(path.Join(filops.BankDataDirectory, category))
	if err != nil {
		log.Println(err)

		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Filops",
			FunctionName: "GetBankDataCategoryList",
			Description:  "failed read folder",
			Trace:        err,
		}.Error()
	}

	for _, f := range files {
		if !f.IsDir() {
			filenames = append(filenames, f.Name())
		}
	}

	if pagination != nil {
		last := pagination.Offset + pagination.Count
		if last > int64(len(filenames)) {
			last = int64(len(filenames))
		}

		return filenames[pagination.Offset:last], nil
	}

	return filenames, nil
}
