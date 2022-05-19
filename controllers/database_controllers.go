package controllers

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
)

type tomlConfig struct {
	SchemaName1 string
	SchemaName2 string
	Servers     map[string]database
}

type database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type column struct {
	Name       string
	Type       string
	IsNullable string
	Default    interface{}
	After      string
}

var (
	driverName string
	dbConfig   tomlConfig
	dLog       *log.Logger
)

func init() {
	driverName = "mysql"
}

// Conn Stablishes connection to database
func Conn(dataSourceName string) *sql.DB {
	//Opens connection with the given datasourceName and driver=mysql
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		dLog.Println(err.Error())
		os.Exit(-1)
	}
	if err := db.Ping(); err != nil {
		panic("ERROR:" + err.Error())
	}
	return db
}

func getSource(db string) (source string) {
	// dataSourceName = "user:password@tcp(localhost:3306)/database?charset=utf8"
	source = dbConfig.Servers[db].User +
		":" +
		dbConfig.Servers[db].Password +
		"@tcp(" +
		dbConfig.Servers[db].Host +
		":" +
		dbConfig.Servers[db].Port +
		")/" +
		dbConfig.Servers[db].Name +
		"?charset=utf8"
	return
}

func main() {

	// Create diff log
	logFile, err := os.Create("diff.log")
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	// information logs
	dLog = log.New(logFile, "[Info]", log.LstdFlags) //|log.Lshortfile)

	dLog.SetPrefix("[Info]")
	//flags
	dLog.SetFlags(dLog.Flags() | log.LstdFlags)
	dLog.Println("flags...")
	// parsing toml config file
	if _, err := toml.DecodeFile("config.toml", &dbConfig); err != nil {
		fmt.Println("Error parsing config toml:", err.Error())
		return
	}

	dLog.Printf("Loading %s/%s ", dbConfig.Servers["1"].Host, dbConfig.Servers["1"].Name)
	schema1 := dbConfig.SchemaName1
	db1 := Conn(getSource("1"))
	defer db1.Close()

	dLog.Printf("Loading %s/%s", dbConfig.Servers["2"].Host, dbConfig.Servers["2"].Name)
	schema2 := dbConfig.SchemaName2
	db2 := Conn(getSource("2"))
	defer db2.Close()

	dLog.Println("Connected.")
	//dLog.Println("comparing Triggers")
	//TriggerDiff(db1, db2, schema1, schema2)
	// Functions
	dLog.Println("comparing functions...")
	FunctionDiff(db1, db2, schema1, schema2)
	// Tables
	dLog.Println("comparing tables...")
	ts, b := TableDiff(db1, db2, schema1, schema2)
	if b {
		dLog.Println("Found differences...")
		//compare columns and indexes
		dLog.Println("comparing columns")
		ColumnDiff(db1, db2, schema1, schema2, ts)
		dLog.Println("comparing indexes")
		IndexDiff(db1, db2, schema1, schema2, ts)
	}

	dLog.Println("Done!")
}

// TableDiff
func TableDiff(db1, db2 *sql.DB, schema1, schema2 string) (t []string, b bool) {
	tableName1, err := getTableName(db1, schema1)
	if err != nil {
		dLog.Fatalln(err.Error())
	}

	dLog.Println(dbConfig.Servers["1"].Host, "/", schema1, " tabla: ", tableName1)
	tableName2, err := getTableName(db2, schema2)
	if err != nil {
		dLog.Fatalln(err.Error())
	}

	dLog.Println(dbConfig.Servers["2"].Host, "/", schema2, " tabla: ", tableName2)
	if !isEqual(tableName1, tableName2) {
		t = diffName(tableName1, tableName2)
		dLog.Printf("differences: %d respectively: %s", len(t), t)
		return t, false
	}
	t = tableName1
	dLog.Printf("Ambas tablas de base de datos son iguales.")
	return t, true
}

func getTableName(s *sql.DB, table string) (ts []string, err error) {
	stm, perr := s.Prepare("select table_name from information_schema.tables where table_schema=? order by table_name")
	if perr != nil {
		err = perr
		return
	}
	defer stm.Close()
	q, qerr := stm.Query(table)
	if qerr != nil {
		err = qerr
		return
	}
	defer q.Close()

	for q.Next() {
		var name string
		if err := q.Scan(&name); err != nil {
			log.Fatal(err)
		}
		ts = append(ts, name)
	}
	return
}

// TriggerDiff
func TriggerDiff(db1, db2 *sql.DB, schema1, schema2 string) bool {
	triggerName1, err := getTriggerName(db1, schema1)
	if err != nil {
		dLog.Fatalln(err.Error())
	}
	triggerName2, err := getTriggerName(db2, schema2)
	if err != nil {
		dLog.Fatalln(err.Error())
	}
	if !isEqual(triggerName1, triggerName2) {
		dt := diffName(triggerName1, triggerName2)
		dLog.Printf("differences: %d respectively: %s", len(dt), dt)
		return false
	}
	// dLog.Printf("两个数据库触发器相同")
	return true
}

func getTriggerName(s *sql.DB, schema string) (ts []string, err error) {
	stm, perr := s.Prepare("select TRIGGER_NAME from information_schema.triggers where TRIGGER_SCHEMA=? order by TRIGGER_NAME")
	if perr != nil {
		err = perr
		return
	}
	defer stm.Close()
	q, qerr := stm.Query(schema)
	if qerr != nil {
		err = qerr
		return
	}
	defer q.Close()

	for q.Next() {
		var name string
		if err := q.Scan(&name); err != nil {
			log.Fatal(err)
		}
		ts = append(ts, name)
	}
	return
}

// FunctionDiff
func FunctionDiff(db1, db2 *sql.DB, schema1, schema2 string) bool {
	functionName1, err := getFunctionName(db1, schema1)
	if err != nil {
		dLog.Fatalln(err.Error())
	}
	functionName2, err := getFunctionName(db2, schema2)
	if err != nil {
		dLog.Fatalln(err.Error())
	}
	dLog.Println(functionName1)
	dLog.Println(functionName2)
	if !isEqual(functionName1, functionName2) {
		dt := diffName(functionName1, functionName2)
		dLog.Printf("differences: %d respectively: %s", len(dt), dt)
		return false
	}
	// dLog.Printf("两个数据库函数相同")
	return true
}

func getFunctionName(s *sql.DB, schema string) (ts []string, err error) {
	stm, perr := s.Prepare("select ROUTINE_NAME from information_schema.routines where ROUTINE_SCHEMA=? and ROUTINE_TYPE='FUNCTION' order by ROUTINE_NAME")
	if perr != nil {
		err = perr
		return
	}
	defer stm.Close()
	q, qerr := stm.Query(schema)
	if qerr != nil {
		err = qerr
		return
	}
	defer q.Close()

	for q.Next() {
		var name string
		if err := q.Scan(&name); err != nil {
			log.Fatal(err)
		}
		ts = append(ts, name)
	}
	return
}

func genAlterSql(t string, col column) string {
	var after string
	if col.After != "" {
		after = fmt.Sprintf(" AFTER `%s`", col.After)
	}

	var isNull string
	if col.IsNullable == "YES" {
		isNull = " NULL"
	} else {
		isNull = " NOT NULL"
	}

	var defaultValue string
	if col.Default == nil {
		defaultValue = " DEFAULT NULL"
	} else if col.Default != "" {
		defaultValue = fmt.Sprintf(" DEFAULT '%s'", col.Default)
	}

	return fmt.Sprintf("ALTER TABLE `%s` ADD COLUMN `%s` %s%s%s%s;", t, col.Name, col.Type, isNull, defaultValue, after)
}

// ColumnDiff
func ColumnDiff(db1, db2 *sql.DB, schema1, schema2 string, table []string) {
	for _, t := range table {
		columnName1, err := getColumnName(db1, schema1, t)
		if err != nil {
			dLog.Fatalln(err.Error())
		}
		columnName2, err := getColumnName(db2, schema2, t)
		if err != nil {
			dLog.Fatalln(err.Error())
		}
		if !columnIsEqual(columnName1, columnName2) {
			// dt := diffName(columnName1, columnName2)
			col1, col2 := columnDiff(columnName1, columnName2)

			dLog.Printf("database: %s table: %s different columns: %d", schema1, t, len(col1))
			for _, col := range col1 {
				dLog.Printf(genAlterSql(t, col))
			}

			dLog.Printf("database: %s table: %s different columns: %d", schema2, t, len(col2))
			for _, col := range col2 {
				dLog.Printf(genAlterSql(t, col))
			}
		} else {
			dLog.Printf("both tables %s have same columns", t)
		}
	}
}

func getColumnName(s *sql.DB, schema, table string) (ts []column, err error) {
	stm, perr := s.Prepare("select COLUMN_NAME,column_type,column_default,is_nullable from information_schema.columns where TABLE_SCHEMA=? and TABLE_NAME=? order by ordinal_position asc")
	if perr != nil {
		err = perr
		return
	}
	defer stm.Close()
	q, qerr := stm.Query(schema, table)
	if qerr != nil {
		err = qerr
		return
	}
	defer q.Close()

	ts = make([]column, 0)

	var after string
	for q.Next() {
		var column_name string
		var column_type string
		var column_default interface{}
		var is_nullable string

		if err := q.Scan(&column_name, &column_type, &column_default, &is_nullable); err != nil {

			return nil, err
		}

		col := column{}
		col.Name = column_name
		col.Type = column_type
		col.IsNullable = is_nullable
		col.Default = column_default

		if after == "" {
			col.After = ""
		} else {
			col.After = after
		}
		after = col.Name

		ts = append(ts, col)
	}
	return
}

// IndexDiff
func IndexDiff(db1, db2 *sql.DB, schema1, schema2 string, table []string) (str string, err error) {
	for _, t := range table {
		indexName1, err := getIndexName(db1, schema1, t)
		if err != nil {
			return "", err
		}
		indexName2, err := getIndexName(db2, schema2, t)
		if err != nil {
			return "", err
		}
		if !isEqual(indexName1, indexName2) {
			dt := diffName(indexName1, indexName2)
			str := fmt.Sprintf("both databases %s with different indexes with a total of %d respectively: %s", t, len(dt), dt)
			return str, nil
		} else {
			return fmt.Sprintf("dos tablas de base de datos: %s tienen el mismo indice", t), err
		}
	}
}

func getIndexName(s *sql.DB, schema, table string) (ts []string, err error) {
	stm, perr := s.Prepare("select INDEX_NAME from information_schema.STATISTICS where TABLE_SCHEMA=? and TABLE_NAME=? order by INDEX_NAME")
	if perr != nil {
		err = perr
		return
	}
	defer stm.Close()
	q, qerr := stm.Query(schema, table)
	if qerr != nil {
		err = qerr
		return
	}
	defer q.Close()

	for q.Next() {
		var name string
		if err := q.Scan(&name); err != nil {
			return nil, err
		}
		ts = append(ts, name)
	}
	return
}

func columnIsEqual(x, y []column) bool {
	if len(x) != len(y) {
		return false
	}

	for _, col1 := range x {
		isExist := false
		for _, col2 := range y {
			if col2.Name == col1.Name {
				isExist = true
				break
			}
		}

		if isExist == false {
			return false
		}
	}

	return true
}

func columnDiff(x, y []column) (x1, x2 []column) {

	//generate statement
	//First determine the statement that x, y do not have
	for _, col1 := range x {
		isExist := false
		for _, col2 := range y {
			if col1.Name == col2.Name {
				isExist = true
				break
			}
		}

		if isExist == false {

			//generate statement
			x2 = append(x2, col1)
		}
	}

	for _, col1 := range y {
		isExist := false
		for _, col2 := range x {
			if col1.Name == col2.Name {
				isExist = true
				break
			}
		}

		if isExist == false {

			//generate statement
			x1 = append(x1, col1)
		}
	}

	return

}

func isEqual(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func diffName(a, b []string) []string {
	c := a[:0]
	m := make(map[string]int)
	for _, s := range a {
		m[s] = 1
	}
	//intersection
	n := make(map[string]int)
	for _, s := range b {
		if _, ok := m[s]; ok {
			n[s] = 1
		}
	}
	// different from "a"
	for _, s := range a {
		if _, ok := n[s]; !ok {
			c = append(c, s)
		}
	}
	// different from "b"
	for _, s := range b {
		if _, ok := n[s]; !ok {
			c = append(c, s)
		}
	}
	return c
}
