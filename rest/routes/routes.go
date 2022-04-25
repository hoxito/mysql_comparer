package routes

import (
	"statsv0/controllers"
	"statsv0/rest/middlewares"

	"github.com/gin-gonic/gin"
)

func PeakHourRoute(router *gin.Engine) {
	router.GET("/v1/deploy/update", middlewares.ValidateAuthentication, controllers.GetYearPeakHourSorted())
	router.DELETE("/v1/deploy/rollback", middlewares.ValidateAuthentication, controllers.DeletePeakHour())
	router.DELETE("/v1/deploy/rollback/:minutes", middlewares.ValidateAuthentication, controllers.DeletePeakHour())

}

func GreatestOrdersRoute(router *gin.Engine) {

}
