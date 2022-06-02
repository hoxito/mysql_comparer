package services

import (
	"fmt"
	"log"
	"os"
	"strings"

	"database/sql"

	"github.com/hoxito/mysql_comparer/models"

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
func Conn(dataSourceName string) (db *sql.DB, err error) {
	//Opens connection with the given datasourceName and driver=mysql
	fmt.Println("db: " + dataSourceName + " " + driverName)
	db, err = sql.Open(driverName, dataSourceName)
	if err != nil {

		os.Exit(-1)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		panic("ERROR:" + err.Error())
	}
	return db, nil
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

func Main() (diff string, tables string, tablesdiff []*models.TableDiff, err error) {

	driverName = "mysql"
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
	diff = "flags..."
	// parsing toml config file
	if _, err := toml.DecodeFile("config.toml", &dbConfig); err != nil {
		fmt.Println("Error parsing config toml:", err.Error())
		return "error parsing toml", "", nil, err
	}

	dLog.Printf("Loading %s/%s ", dbConfig.Servers["1"].Host, dbConfig.Servers["1"].Name)
	schema1 := dbConfig.SchemaName1
	fmt.Println("source:" + getSource("1"))
	db1, err := Conn(getSource("1"))
	if err != nil {
		return "error connecting to db 1", "", nil, err
	}
	defer db1.Close()

	dLog.Printf("Loading %s/%s", dbConfig.Servers["2"].Host, dbConfig.Servers["2"].Name)
	schema2 := dbConfig.SchemaName2
	db2, err := Conn(getSource("2"))
	if err != nil {
		return "error connecting to db 2", "", nil, err
	}
	defer db2.Close()

	diff = "Connected."
	//diff=("comparing Triggers"
	//TriggerDiff(db1, db2, schema1, schema2)
	// Functions
	diff = diff + "comparing functions..."
	diff = diff + "\n"
	_, res, err := FunctionDiff(db1, db2, schema1, schema2)
	diff = diff + res
	// Tables
	diff = "comparing tables..."
	diff = diff + "comparing tables...\n"
	diff = diff + "\n"
	ts, str, b := TableDiff(db1, db2, schema1, schema2)
	diff = diff + str
	tables = str
	if b {
		diff = diff + "Found differences..."
		diff = diff + "found diffs:" + strings.Join(ts[:], ",")
		diff = diff + "\n"
		//compare columns and indexes
		diff = diff + "comparing columns"
		diff = diff + "comparing columns"
		diff = diff + "\n"
		str, tablesdiff = ColumnDiff(db1, db2, schema1, schema2, ts)
		diff = diff + str
		diff = diff + "comparing indexes"
		diff = diff + "comparing indexes"
		diff = diff + "\n"
		str = IndexDiff(db1, db2, schema1, schema2, ts)
		diff = diff + str
	} else {
		for _, name := range ts {

			tablesdiff = append(tablesdiff,
				&models.TableDiff{
					Name:    name,
					Script1: fmt.Sprintf(`CREATE TABLE %s ( id VARCHAR(50) NOT NULL, PRIMARY KEY (id));`, name),
				})
		}
	}

	diff = diff + "\n"
	diff = diff + "Done!"
	return diff, tables, tablesdiff, nil
}

// TableDiff
func TableDiff(db1, db2 *sql.DB, schema1, schema2 string) (t []string, str string, b bool) {
	tableName1, err := getTableName(db1, schema1)
	if err != nil {
		dLog.Fatalln(err.Error())
	}
	str = fmt.Sprintf("%s / %s tables: %s", (dbConfig.Servers["1"].Host), schema1, tableName1)
	tableName2, err := getTableName(db2, schema2)
	if err != nil {
		dLog.Fatalln(err.Error())
	}

	str = str + fmt.Sprintf("%s / %s tables: %s", (dbConfig.Servers["2"].Host), schema2, tableName2)
	if !isEqual(tableName1, tableName2) {
		t = diffName(tableName1, tableName2)
		str = str + fmt.Sprintf("differences: %d respectively: %s", len(t), t)
		return t, str, false
	}
	t = tableName1
	str = str + fmt.Sprintf("Ambos esquemas contienen las mismas tablas.")
	return t, str, true
}

func getTableName(s *sql.DB, table string) (ts []string, err error) {
	stm, perr := s.Prepare("select table_name from information_schema.tables where table_schema=? order by table_name")
	if perr != nil {
		err = perr
		return nil, err
	}
	defer stm.Close()
	q, qerr := stm.Query(table)
	if qerr != nil {
		err = qerr
		return nil, err
	}
	defer q.Close()

	for q.Next() {
		var name string
		if err := q.Scan(&name); err != nil {
			log.Fatal(err)
			return nil, err
		}
		ts = append(ts, name)
	}
	return ts, nil
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
func FunctionDiff(db1, db2 *sql.DB, schema1, schema2 string) (isDiff bool, str string, err error) {
	functionName1, err := getFunctionName(db1, schema1)
	if err != nil {
		return false, "", err
	}
	functionName2, err := getFunctionName(db2, schema2)
	if err != nil {
		return false, "", err
	}
	str = str + strings.Join(functionName1, " | ") + "\n"
	str = str + strings.Join(functionName2, " | ") + "\n"
	if !isEqual(functionName1, functionName2) {
		dt := diffName(functionName1, functionName2)
		str = str + fmt.Sprintf("differences: %d respectively: %s \n", len(dt), dt)
		return false, str, nil
	}
	return true, str, nil
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
func ColumnDiff(db1, db2 *sql.DB, schema1, schema2 string, table []string) (diff string, tablediff []*models.TableDiff) {
	var td models.TableDiff = models.TableDiff{}
	for _, t := range table {

		columnName1, err := getColumnName(db1, schema1, t)
		if err != nil {
			dLog.Fatalln(err.Error())
			return err.Error(), nil
		}
		columnName2, err := getColumnName(db2, schema2, t)
		if err != nil {
			dLog.Fatalln(err.Error())
			return err.Error(), nil
		}
		if !columnIsEqual(columnName1, columnName2) {
			// dt := diffName(columnName1, columnName2)
			col1, col2 := columnDiff(columnName1, columnName2)

			dLog.Printf("database: %s table: %s different columns: %d", schema1, t, len(col1))
			diff = diff + fmt.Sprintf("database SLAVE table: %s different columns: %d", t, len(col1))
			td.Name = t
			td.Db1 = schema1
			diff = diff + "\n"

			for _, col := range col1 {
				s := genAlterSql(t, col)
				dLog.Printf(s)
				diff = diff + s
				td.Script1 = td.Script1 + s
				diff = diff + "\n"
			}

			dLog.Printf("database: %s table: %s different columns: %d", schema2, t, len(col2))
			diff = diff + fmt.Sprintf("database: MASTER table: %s different columns: %d", t, len(col2))
			diff = diff + "\n"
			td.Db1 = schema2
			for _, col := range col2 {
				s := genAlterSql(t, col)
				dLog.Printf(s)
				diff = diff + s
				td.Script2 = td.Script2 + s
				diff = diff + "\n"
			}
		} else {
			dLog.Printf("both tables %s have same columns", t)
			diff = diff + fmt.Sprintf("both tables %s have same columns", t)
			diff = diff + "\n"
		}
	}
	tablediff = append(tablediff, &td)
	return diff, tablediff
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
			log.Fatal(err)
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
func IndexDiff(db1, db2 *sql.DB, schema1, schema2 string, table []string) (diff string) {
	for _, t := range table {
		indexName1, err := getIndexName(db1, schema1, t)
		if err != nil {
			dLog.Fatalln(err.Error())
			return err.Error()
		}
		indexName2, err := getIndexName(db2, schema2, t)
		if err != nil {
			dLog.Fatalln(err.Error())
			return err.Error()
		}
		if !isEqual(indexName1, indexName2) {
			dt := diffName(indexName1, indexName2)
			dLog.Printf("both databases %s with different indexes with a total of %d respectively: %s", t, len(dt), dt)
			diff = diff + fmt.Sprintf("both databases %s with different indexes with a total of %d respectively: %s", t, len(dt), dt)
		} else {
		}
	}

	return diff
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
			log.Fatal(err)
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
