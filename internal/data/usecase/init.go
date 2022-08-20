package usecase

type Repositories struct {
	TOSORepo
	DILRepo
	ReportRepo
	TemuanRepo
}

type Usecase struct {
	repositories *Repositories
}

func New(repositories *Repositories) *Usecase {
	return &Usecase{
		repositories: repositories,
	}
}
