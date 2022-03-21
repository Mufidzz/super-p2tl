package main

import (
	"github.com/SuperP2TL/Backend/internal/data"
	"github.com/SuperP2TL/Backend/internal/repository/postgre"
	usr "github.com/SuperP2TL/Backend/internal/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	postgreRepo, err := postgre.New("user=super-p2tl-staging password=super-p2tl-staging dbname=super-p2tl-chief host=139.162.36.125 sslmode=disable")
	if err != nil {
		log.Printf("[DB Init] error initialize database, trace %v", err)
	}

	StartREST(postgreRepo)
}

func StartREST(pg *postgre.Postgre) {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))

	data.StartHTTP(router, pg)
	usr.StartHTTP(router, pg)

	router.Run(":4455")
}
