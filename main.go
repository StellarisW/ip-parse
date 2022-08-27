package main

import (
	"github.com/gin-gonic/gin"
	"main/app/common/log"
	"main/app/db"
	"main/app/router"
	"net/http"
	"time"
)

const port = ":8083"

func main() {
	log.NewLogger()
	logger := log.GetSugaredLogger()

	logger.Infof("initiating searcher...")
	err := db.InitSearcher()
	if err != nil {
		logger.Fatalf("initialize searcher failed, err: %v", err)
	}
	logger.Infof("initiate searcher successfully")

	gin.SetMode(gin.DebugMode)

	logger.Infof("initiating routers...")
	routers := router.InitRouter()
	logger.Infof("initiate routers successfully")

	server := &http.Server{
		Addr:           port,
		Handler:        routers,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20, // 请求头最大大小 16 MB
	}

	logger.Infof("server running on %s ...\n", port)
	logger.Errorf(server.ListenAndServe().Error())
}
