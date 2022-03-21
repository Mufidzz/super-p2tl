package usecase

func (uc *Usecase) AssignUserTOSOWorkload(userID int64, tosoIDs []int64) (id int64, err error) {
	return uc.repositories.AssignToSoToUser(userID, tosoIDs)
}
