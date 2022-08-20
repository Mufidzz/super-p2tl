package usecase

import "github.com/SuperP2TL/Backend/presentation"

func (uc *Usecase) GetDataUserTemuan(userId int) (res []presentation.GetDataUserTemuanResponse, err error) {
	return uc.repositories.GetDataUserTemuan(userId)
}
