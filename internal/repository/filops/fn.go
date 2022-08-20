package filops

import (
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func (filops *Filops) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	splitFilename := strings.Split(file.Filename, ".")

	ext := splitFilename[len(splitFilename)-1]

	out, err := os.Create(dst + "." + ext)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func (filops *Filops) GetTemuanBaseDirectory() string {
	return filops.TemuanDataDirectory
}
