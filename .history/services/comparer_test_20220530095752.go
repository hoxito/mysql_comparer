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
        arg1, arg2 string
        want string
    }{
        {"table1", "table2", "2"},
		
        {"table1", "table1", ""},
    }

	  for _, tt := range tests {
        testname := fmt.Sprintf("ge%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
	got:=diffName(arg1, arg2)

	response.Assert(200, "")
	assert.Equal(t, context.Errors.Last().Error(), custerror.Unauthorized.Error())
}
