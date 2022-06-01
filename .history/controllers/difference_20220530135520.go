package controllers

import (
	"mysqlbinlogparser/models"
	"mysqlbinlogparser/services"
	"mysqlbinlogparser/tools/metrics"

	"github.com/gin-gonic/gin"
)

// var DifferencesCollection *mongo.Collection = configs.GetCollection(configs.DB, "Differeces")
/*
*
*@title Swagger wallet API
*@description This is a Simple wallet API that can manage users, wallets and transactions between these wallets.
* @license.url http://www.apache.org/licenses/LICENSE-2.0.html
*
 */

/*
*
*@Param user_id  path int true "User ID"
*
*
*
* @Success 200 {array} models.Difference
*
*@Failure 400
*
 */
func GetDiff(c *gin.Context) {
	var diffs models.Difference
	diffs.Master = "aws siis"
	diffs.Slave = "aws siistesting"
	diffs.Differences, diffs.Tables, diffs.TableDifferences, diffs.Errors = services.Main()
	metrics.DiffCounter.Inc()
	c.JSON(200, diffs)

}
