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
*@description.markdown Gets the difference between 2 databases
*@Param user_id  path string true "User ID"
*@Param user_id  body string true "User ID"
*@Param user_id  path string true "User ID"
// @Param   enumstring  query     string     false  "string enums"       Enums(A, B, C)
// @Param   enumint     body     int        false  "int enums"          Enums(1, 2, 3)
// @Param   enumnumber  body     number     false  "int enums"          Enums(1.1, 1.2, 1.3)
// @Param   string      body     string     false  "string valid"       minlength(5)  maxlength(10)
// @Param   int         body     int        false  "int valid"          minimum(1)    maximum(10)
// @Param   default     body     string     false  "string default"     default(A)
// @Param   example     body     string     false  "string example"     example(string)
// @Param   collection  body     []string   false  "string collection"  collectionFormat(multi)
// @Param   extensions  body     []string   false  "string collection"  extensions(x-example=test,x-nullable)
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
