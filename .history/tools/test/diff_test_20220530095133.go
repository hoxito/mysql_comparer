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
	context.JSON(500, gin.H{"error": "Internal server error"})
	response.Assert(500, "{\"error\":\"Internal server error\"}")
}

func TestDiff2(t *testing.T) {
	response := ResponseWriter(t)
	context, _ := gin.CreateTestContext(response)
	context.Request, _ = http.NewRequest("GET", "/", nil)

	// GetDiff(context)

	response.Assert(0, "")
	assert.Equal(t, context.Errors.Last().Error(), custerror.Unauthorized.Error())
}
