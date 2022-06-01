package test

import (
	"mysqlbinlogparser/tools/custerror"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func TestDiff(t *testing.T) {
	response := ResponseWriter(t)
	context, _ := gin.CreateTestContext(response)
	context.Request, _ = http.NewRequest("GET", "/v1/db/simple/diff", nil)

	GetDiff(context)

	response.Assert(0, "")
	assert.Equal(t, context.Errors.Last().Error(), custerror.Unauthorized.Error())
}
