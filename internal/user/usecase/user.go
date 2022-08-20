package usecase

import (
	"github.com/SuperP2TL/Backend/presentation"
)

func (uc *Usecase) GetDataPetugas(filter *presentation.FilterParamUserData, pagination presentation.Pagination) (res []presentation.GetDataPetugasResponse, err error) {
	filter.Role = presentation.USER_ROLE_PETUGAS
	return uc.repositories.GetDataPetugas(filter, pagination)
}

func (uc *Usecase) Login(username, password string) (user *presentation.GetUserPasswordResponse, err error) {
	user, err = uc.repositories.GetUserPassword(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	// TODO : use bcrypt
	if user.Password == password {
		return user, nil
	}

	return nil, nil
}
