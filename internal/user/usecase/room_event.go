package usecase

import "github.com/SuperP2TL/Backend/presentation"

func (uc *Usecase) CreateSingleLogRoomEvent(input presentation.CreateLogRoomEventRequest) (lastInsertedID int, err error) {
	return uc.repositories.CreateLogRoomEvent([]presentation.CreateLogRoomEventRequest{input})
}
