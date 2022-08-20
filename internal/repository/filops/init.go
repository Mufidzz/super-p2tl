package filops

import (
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"path"
)

type Filops struct {
	RootDirectory           string
	BankDataDirectory       string
	PenormalanDataDirectory string
	TemuanDataDirectory     string
}

func New(
	rootDir string,
) (*Filops, error) {
	_, err := os.Stat(rootDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(rootDir, os.ModeDir)
		if err != nil {
			return nil, fmt.Errorf("[Filops][Init] Failed init folder, trace %v", err)
		}
		fmt.Println("[Filops][Init] Success Creating New Folder as Root")
	}

	return &Filops{
		RootDirectory:           rootDir,
		BankDataDirectory:       path.Join(rootDir, "bank_data"),
		PenormalanDataDirectory: path.Join(rootDir, "penormalan_data"),
		TemuanDataDirectory:     path.Join(rootDir, "temuan"),
	}, nil
}
