package usecase

type Repositories struct {
	TOSORepo
}

type Usecase struct {
	repositories *Repositories
}

func New(repositories *Repositories) *Usecase {
	return &Usecase{
		repositories: repositories,
	}
}
