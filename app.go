package main

import (
	"github.com/SuperP2TL/Backend/internal/cdn"
	"github.com/SuperP2TL/Backend/internal/data"
	"github.com/SuperP2TL/Backend/internal/report"
	"github.com/SuperP2TL/Backend/internal/repository/filops"
	"github.com/SuperP2TL/Backend/internal/repository/postgre"
	usr "github.com/SuperP2TL/Backend/internal/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func main() {
	//postgreRepo, err := postgre.New("user=super-p2tl-staging password=super-p2tl-staging dbname=super-p2tl-chief host=139.162.36.125 sslmode=disable")
	postgreRepo, err := postgre.New("user=super-p2tl-staging password=super-p2tl-staging dbname=super-p2tl-chief host=127.0.0.1 sslmode=disable")
	if err != nil {
		panic("[DB Init] error initialize database, trace " + err.Error())
	}

	filopsRepo, err := filops.New("/var/app/services/super-p2tl/cdn-storage")
	//filopsRepo, err := filops.New("E:\\1.ON GOING\\Super P2TL\\Backend-stable\\cdn-storage")
	if err != nil {
		panic("[Filops Init] error initialize, trace " + err.Error())
	}

	StartREST(postgreRepo, filopsRepo)
}

func StartREST(pg *postgre.Postgre, filops *filops.Filops) {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))

	socketServer := socketio.NewServer(nil)

	data.StartHTTP(router, pg, pg, pg, pg)
	usr.StartHTTP(router, socketServer, pg, pg, pg, pg)
	report.StartHTTP(router, pg, pg)
	cdn.StartHTTP(router, filops, filops, pg)

	router.Run(":4455")
}
