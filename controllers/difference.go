package controllers

import (
	"mysqlbinlogparser/models"
	"mysqlbinlogparser/services"

	"github.com/gin-gonic/gin"
)

// var DifferencesCollection *mongo.Collection = configs.GetCollection(configs.DB, "Differeces")

func GetDiff(c *gin.Context) {
	var diffs models.Difference
	diffs.Master = "aws siis"
	diffs.Slave = "aws siistesting"
	diffs.Differences, diffs.Tables, diffs.TableDifferences, diffs.Errors = services.Main()
	c.JSON(200, diffs)

}
