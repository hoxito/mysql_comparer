package models

type Difference struct {
	Master           string       `json:"master,omitempty"`
	Slave            string       `json:"slave,omitempty" validate:"required"`
	Differences      string       `json:"differences,omitempty" validate:"required"`
	Tables           string       `json:"tables,omitempty" validate:"required"`
	TableDifferences []*TableDiff `json:"tabledifferences,omitempty" validate:"required"`
	Errors           error        `json:"errors,omitempty" `
}
type TableDiff struct {
	Name    string `json:"name,omitempty" `
	Db1     string `json:"db1,omitempty" `
	Script1 string `json:"script1,omitempty"`
	Db2     string `json:"db2,omitempty" `
	Script2 string `json:"script2,omitempty"`
}
