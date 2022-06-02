package routes

import (
	"github.com/hoxito/mysql_comparer/controllers"

	"github.com/gin-gonic/gin"
)

func ComprarerRoute(router *gin.Engine) {

	router.GET("/v1/db/diff", controllers.GetDiff)
	router.GET("/v1/db/diff/notables", controllers.GetDiffNoTables)
	// router.GET("/v1/db/simple/history", middlewares.ValidateAuthentication, controllers.getHistory())
	// router.GET("/v1/db/full/diff", middlewares.ValidateAuthentication, controllers.getHistory())
	// router.GET("/v1/db/full/history", middlewares.ValidateAuthentication, controllers.getHistory())
	// router.GET("/v1/db/log/diff", middlewares.ValidateAuthentication, controllers.getHistory())
	// router.GET("/v1/db/log/history", middlewares.ValidateAuthentication, controllers.getHistory())
	// router.GET("/v1/db/update/", middlewares.ValidateAuthentication, controllers.getHistory())
	// router.GET("/v1/db/rollback/", middlewares.ValidateAuthentication, controllers.DeletePeakHour())

}
