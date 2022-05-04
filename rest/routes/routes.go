package routes

import (
	"mysqlbinlogparser/controllers"
	"mysqlbinlogparser/rest/middlewares"

	"github.com/gin-gonic/gin"
)

func PeakHourRoute(router *gin.Engine) {
	router.GET("/v1/db/simple/diff", middlewares.ValidateAuthentication, controllers.getDiff())
	router.GET("/v1/db/simple/history", middlewares.ValidateAuthentication, controllers.getHistory())
	router.GET("/v1/db/full/diff", middlewares.ValidateAuthentication, controllers.getHistory())
	router.GET("/v1/db/full/history", middlewares.ValidateAuthentication, controllers.getHistory())
	router.GET("/v1/db/log/diff", middlewares.ValidateAuthentication, controllers.getHistory())
	router.GET("/v1/db/log/history", middlewares.ValidateAuthentication, controllers.getHistory())
	router.GET("/v1/db/update/", middlewares.ValidateAuthentication, controllers.getHistory())
	// router.GET("/v1/db/rollback/", middlewares.ValidateAuthentication, controllers.DeletePeakHour())

}

func GreatestOrdersRoute(router *gin.Engine) {

}
