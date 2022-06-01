package services

import (
	"github.com/stretchr/testify/assert"
	"mysqlbinlogparser/tools/custerror"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func TestDiffNameTableDriven(t *testing.T) {
	   var tests = []struct {
        arg1, arg2 int
        want int
    }{
        {"table1", "table2", 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }
	arg1:="table1"
	arg2:="table2"
	
	expect:=
	got:=diffName(arg1, arg2)

	response.Assert(200, "")
	assert.Equal(t, context.Errors.Last().Error(), custerror.Unauthorized.Error())
}
