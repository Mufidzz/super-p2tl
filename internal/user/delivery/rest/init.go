package rest

import (
	"fmt"
	"github.com/SuperP2TL/Backend/presentation"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

type Usecases struct {
	UserUC
	RoomEventUC
	ToSoUC
	TemuanUC
}

type HTTPHandler struct {
	router       *gin.Engine
	socketServer *socketio.Server
	usecases     Usecases
}

func NewHTTP(router *gin.Engine, socketServer *socketio.Server, userUsecase UserUC, roomEventUsecase RoomEventUC, tosoUsecase ToSoUC, temuanUC TemuanUC) *HTTPHandler {
	return &HTTPHandler{
		router:       router,
		socketServer: socketServer,
		usecases: Usecases{
			UserUC:      userUsecase,
			RoomEventUC: roomEventUsecase,
			ToSoUC:      tosoUsecase,
			TemuanUC:    temuanUC,
		},
	}
}

func (handler *HTTPHandler) SetRoutes() {
	router := handler.router

	uaa := router.Group("/user")
	{
		uaa.POST("/login", handler.HandleLogin)
		uaa.GET("/petugas", handler.HandlerGetDataPetugas)
	}

	userWork := router.Group("/user-work")
	{
		userWork.POST("/to-so", handler.HandleAssignUserTOSOWorkload)
		userWork.POST("/temuan", handler.HandleAssignUserTemuanWorkload)

		userWork.GET("/to-so/:user-id", handler.HandleGetDataUserTOSOWorkload)
		userWork.GET("/temuan/:user-id", handler.HandleGetDataUserTemuanWorkload)

		// TODO : Change --> Mock for sending Notification
		userWork.GET("/:id", func(ctx *gin.Context) {
			rooms := handler.socketServer.Rooms("/")
			log.Println(rooms)

			id := ctx.Param("id")

			evtID, err := handler.usecases.CreateSingleLogRoomEvent(presentation.CreateLogRoomEventRequest{
				Namespace: "/",
				Room:      id,
				Args:      "Test",
				Event:     "notify",
				Status:    1,
			})

			if err != nil {
				log.Println(err)
				return
			}

			ok := handler.socketServer.BroadcastToRoom("/", id, "notify", fmt.Sprintf("7#%d", evtID))

			if !ok {
				log.Println("Failed Send Message To Room")
			}

		})
	}
}

func (handler *HTTPHandler) SetSocketRoutes() {
	socketServer := handler.socketServer
	socketServer.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})

	socketServer.OnEvent("/", "register", func(s socketio.Conn, msg string) {
		s.LeaveAll()
		s.Join(msg)
	})

	socketServer.OnError("/", func(conn socketio.Conn, err error) {
		log.Println(err)
	})

	go handler.runSocket()
	handler.router.GET("/socket.io/*any", gin.WrapH(socketServer))
	handler.router.POST("/socket.io/*any", gin.WrapH(socketServer))
}

func (handler *HTTPHandler) runSocket() {
	if err := handler.socketServer.Serve(); err != nil {
		log.Fatalf("socketio listen error: %s\n", err)
	}
}
