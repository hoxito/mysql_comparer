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
		{[]string{"table1"}, []string{"table2"}, []string{"table1", "table2"}},
		{[]string{"table1"}, []string{"table1"}, []string{}},
		{[]string{}, []string{}, []string{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Test diffName(%v,%v)", tt.arg1, tt.arg2)
		t.Run(testname, func(t *testing.T) {
			ans := diffName(tt.arg1, tt.arg2)
			assert.Equal(t, tt.want, ans)
			// if ans != tt.want {
			// 	t.Errorf("got %d, want %d", ans, tt.want)
			// }
		})
	}

}
