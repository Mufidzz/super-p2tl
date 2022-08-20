package filops

import (
	"github.com/SuperP2TL/Backend/pkg/response"
	"mime/multipart"
	"os"
	"path"
)

func (filops *Filops) StorePenormalanFile(fpSebelum, fpSesudah *multipart.FileHeader, id string) error {
	dir := path.Join(filops.PenormalanDataDirectory, id)

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

	//extSebelum := strings.Split(fpSebelum.Filename, ".")
	err = filops.SaveUploadedFile(fpSebelum, path.Join(dir, "foto_sebelum"))
	if err != nil {
		return response.InternalError{
			Type:         "Repo",
			Name:         "Filops",
			FunctionName: "StorePenormalanFile",
			Description:  "failed saving file",
			Trace:        err,
		}.Error()
	}

	//extSesudah := strings.Split(fpSesudah.Filename, ".")
	err = filops.SaveUploadedFile(fpSesudah, path.Join(dir, "foto_sesudah"))
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
