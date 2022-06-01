package main

import (
	"fmt"
	"mysqlbinlogparser/rest/middlewares"
	"mysqlbinlogparser/rest/routes"
	"mysqlbinlogparser/tools/env"
	"time"

	cors "github.com/itsjamie/gin-cors"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title           Swagger Diff API
// @version         1.0
// @description     This is a database differenciator API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   jose aranciba
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080

// @securityDefinitions.basic  BasicAuth
func main() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:3010/swagger/swagger.json")))
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
