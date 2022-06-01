package services

import (
	"fmt"
	"testing"

	"github.com/kinbiko/jsonassert"
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
func TestGetWalletNotFound(t *testing.T) {
	ja := jsonassert.New(t)

	idWallet := ""
	idUser := "testuser1"

	res := GetWallet(idWallet, idUser)
	ja.Assertf(res, `
    {
        "time": "2019-01-28T21:19:42",
        "uuid": "94ae1a31-63b2-4a55-a478-47764b60c56b"
    }`, `
    {
        "time": "<<PRESENCE>>",
        "uuid": "<<PRESENCE>>"

    }`)
}

func GetWallet(string, string) string {
	return `{
        "error": "Wallet Not Found"
    }`
}
