package usecase

type Repositories struct {
	UserRepo
	RoomEventRepo
	TOSORepo
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
