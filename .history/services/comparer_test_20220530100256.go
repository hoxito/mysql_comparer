package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiffName(t *testing.T) {
	var tests = []struct {
		arg1, arg2 []string
		want       []string
	}{
		{["table1"],[ "table2"], "2"},

		{["table1"], ["table1"], ""},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Test diffName(%d,%d)", tt.arg1, tt.arg2)
		t.Run(testname, func(t *testing.T) {
			ans := diffName(tt.arg1, tt.arg2)
			assert.Equal(t,tt.want,ans)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}
