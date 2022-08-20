package presentation

type CreateLogRoomEventRequest struct {
	Namespace string
	Room      string
	Args      string
	Event     string
	Status    int
}
