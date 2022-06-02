package services

import (
	"github.com/stretchr/testify/assert"
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

func TestDiffName(t *testing.T) {
	var tests = []struct {
		arg1, arg2 string
		want       string
	}{
		{"table1", "table2", "2"},

		{"table1", "table1", ""},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Test diffName(%d,%d)", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := diffName(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}