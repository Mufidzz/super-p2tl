package data

import (
	"github.com/SuperP2TL/Backend/internal/user/delivery/rest"
	"github.com/SuperP2TL/Backend/internal/user/usecase"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

type Domain struct {
	Usecase *usecase.Usecase
}

func StartHTTP(router *gin.Engine, socketServer *socketio.Server, userRepo usecase.UserRepo, roomEventRepo usecase.RoomEventRepo, tosoRepo usecase.TOSORepo, temuanRepo usecase.TemuanRepo) *Domain {
	uc := usecase.New(&usecase.Repositories{
		UserRepo:      userRepo,
		RoomEventRepo: roomEventRepo,
		TOSORepo:      tosoRepo,
		TemuanRepo:    temuanRepo,
	})

	httpHandler := rest.NewHTTP(router, socketServer, uc, uc, uc, uc)
	httpHandler.SetRoutes()
	httpHandler.SetSocketRoutes()

	return &Domain{
		Usecase: uc,
	}
}
