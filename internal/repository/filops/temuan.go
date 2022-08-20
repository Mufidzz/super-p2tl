package filops

import (
	"github.com/SuperP2TL/Backend/pkg/response"
	"mime/multipart"
	"os"
	"path"
)

func (filops *Filops) StoreTemuanFile(fpBA, fpLokasi *multipart.FileHeader, id string) error {
	dir := path.Join(filops.TemuanDataDirectory, id)

	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return response.InternalError{
			Type:         "Repo",
			Name:         "Filops",
			FunctionName: "StorePenormalanFile",
			Description:  "failed running create directory",
			Trace:        err,
		}.Error()
	}

	err = filops.SaveUploadedFile(fpBA, path.Join(dir, "foto_ba"))
	if err != nil {
		return response.InternalError{
			Type:         "Repo",
			Name:         "Filops",
			FunctionName: "StorePenormalanFile",
			Description:  "failed saving file",
			Trace:        err,
		}.Error()
	}

	err = filops.SaveUploadedFile(fpLokasi, path.Join(dir, "foto_lokasi"))
	if err != nil {
		return response.InternalError{
			Type:         "Repo",
			Name:         "Filops",
			FunctionName: "StorePenormalanFile",
			Description:  "failed saving file",
			Trace:        err,
		}.Error()
	}

	return nil
}
