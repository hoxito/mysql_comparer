package services

import (
	"github.com/stretchr/testify/assert"
	"mysqlbinlogparser/tools/custerror"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func TestDiffName(t *testing.T) {
	diffName("table1", "table2")

	response.Assert(200, "")
	assert.Equal(t, context.Errors.Last().Error(), custerror.Unauthorized.Error())
}
