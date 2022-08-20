package usecase

type Repositories struct {
	BankDataFilopsRepo
	PenormalanFilopsRepo
	ReportRepo
}

type Usecase struct {
	repositories *Repositories
}

func New(repositories *Repositories) *Usecase {
	return &Usecase{
		repositories: repositories,
	}
}
