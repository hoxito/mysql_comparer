package main

import (
	"fmt"
	"mysqlbinlogparser/rest/middlewares"
	"mysqlbinlogparser/rest/routes"
	"mysqlbinlogparser/tools/env"
	"time"

	docs "mysqlbinlogparser/docs"

	cors "github.com/itsjamie/gin-cors"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.ComprarerRoute(router)
	router.Use(middlewares.ErrorHandler)

	router.Use(cors.Middleware(cors.Config{
		Origins:         "http://localhost:80,http://localhost:8080",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, Size",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	// configs.ConnectDB() implement db to save differences history

	// prometheus.MustRegister(metrics.DiffCounter)
	// http.Handle("/metrics", promhttp.Handler())
	// http.ListenAndServe(":2000", nil)
	router.Run(fmt.Sprintf(":%d", env.Get().Port))
}
