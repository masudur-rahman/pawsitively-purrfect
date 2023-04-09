package server

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/masudur-rahman/pawsitively-purrfect/configs"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql/postgres/pb"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"

	_ "github.com/lib/pq"
)

func getPostgresConnection() (*sql.Conn, error) {
	db, err := sql.Open("postgres", configs.PurrfectConfig.Database.Postgres.String())
	if err != nil {
		return nil, err
	}

	conn, err := db.Conn(context.Background())
	if err != nil {
		return nil, err
	}

	if err = conn.PingContext(context.Background()); err != nil {
		return nil, err
	}
	return conn, nil
}

func generateReadQuery(tableName string, queryParams map[string]interface{}) string {
	var conditions []string

	for key, value := range queryParams {
		condition := fmt.Sprintf("%s = ", key)

		switch v := value.(type) {
		case string:
			condition += fmt.Sprintf("'%s'", v)
		default:
			condition += fmt.Sprintf("%v", v)
		}

		conditions = append(conditions, condition)
	}

	conditionString := strings.Join(conditions, " AND ")
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s", tableName, conditionString)

	return query
}

func scanSingleRecord(rows *sql.Rows) (map[string]interface{}, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]string, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range columns {
		valuePtrs[i] = &values[i]
	}

	if err = rows.Scan(valuePtrs...); err != nil {
		return nil, err
	}
	record := make(map[string]interface{})
	for i, col := range columns {
		record[col] = values[i]
	}

	return record, nil
}

func executeReadQuery(ctx context.Context, query string, conn *sql.Conn, lim int64) ([]map[string]interface{}, error) {
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	records := make([]map[string]interface{}, 0)

	for rows.Next() {
		record, err := scanSingleRecord(rows)
		if err != nil {
			return nil, err
		}

		records = append(records, record)
		if lim > 0 && int64(len(records)) >= lim {
			break
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if lim == 1 && len(records) < 1 {
		return nil, sql.ErrNoRows
	}

	return records, nil
}

func generateInsertQuery(tableName string, record map[string]interface{}) string {
	var cols []string
	var vals []string

	for col, val := range record {
		cols = append(cols, col)
		switch v := val.(type) {
		case string:
			vals = append(vals, fmt.Sprintf("'%s'", strings.ReplaceAll(v, "'", "''")))
		case time.Time:
			vals = append(vals, fmt.Sprintf("'%s'", v.Format("2006-01-02 15:04:05")))
		default:
			vals = append(vals, fmt.Sprintf("%v", v))
		}
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(cols, ", "), strings.Join(vals, ", "))

	return query
}

func executeWriteQuery(ctx context.Context, query string, conn *sql.Conn) (sql.Result, error) {
	result, err := conn.ExecContext(ctx, query)

	return result, err
}

func mapToRecord(record map[string]interface{}) (*pb.RecordResponse, error) {
	pm, err := pkg.ToProtoAny(record)
	if err != nil {
		return nil, err
	}

	return &pb.RecordResponse{Record: pm}, nil
}

func mapsToRecords(records []map[string]interface{}) (*pb.RecordsResponse, error) {
	rr := &pb.RecordsResponse{
		Records: make([]*pb.RecordResponse, 0, len(records)),
	}

	for _, record := range records {
		r, err := mapToRecord(record)
		if err != nil {
			return nil, err
		}

		rr.Records = append(rr.Records, r)
	}
	return rr, nil
}

func getTableName(table interface{}) string {
	tableName := reflect.TypeOf(table).Name()
	if method := reflect.ValueOf(table).MethodByName("TableName"); method.IsValid() {
		rs := method.Call([]reflect.Value{})
		tableName = rs[0].String()
	}

	return tableName
}

func getTableInfo(table interface{}) ([]fieldInfo, error) {
	tableType := reflect.TypeOf(table)
	tableValue := reflect.ValueOf(table)

	if tableType.Kind() == reflect.Ptr {
		tableType = tableType.Elem()
		tableValue = tableValue.Elem()
	}

	if tableType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("getTableInfo: table is not a struct")
	}

	var fields []fieldInfo
	for i := 0; i < tableType.NumField(); i++ {
		fieldType := tableType.Field(i)
		fieldValue := tableValue.Field(i)
		// Skip any field that is not exported (starts with a lowercase letter)
		if fieldType.PkgPath != "" {
			fmt.Println("non-exported fields: ", fieldType.Name)
			continue
		}

		field := getFieldInfo(fieldType, fieldValue)

		fields = append(fields, field)
	}

	return fields, nil
}

func getFieldInfo(fieldType reflect.StructField, fieldValue reflect.Value) fieldInfo {
	fieldName := getFieldName(fieldType)
	sqlType := getSQLType(fieldValue.Type())
	return fieldInfo{
		Name: fieldName,
		Type: sqlType,
	}
}

func getFieldName(fieldType reflect.StructField) string {
	fieldName := fieldType.Name
	if jsonTag := fieldType.Tag.Get("json"); jsonTag != "" {
		fieldName = strings.Split(jsonTag, ",")[0]
	}
	return fieldName
}

func getSQLType(fieldType reflect.Type) string {
	switch fieldType.Kind() {
	case reflect.Int, reflect.Int32:
		return "INTEGER"
	case reflect.Int64, reflect.Uint64:
		return "BIGINT"
	case reflect.Float32, reflect.Float64:
		return "FLOAT"
	case reflect.Bool:
		return "BOOLEAN"
	case reflect.String:
		return "VARCHAR(255)"
	case reflect.Struct:
		if fieldType == reflect.TypeOf(time.Time{}) {
			return "TIMESTAMP WITH TIME ZONE"
		}
	}

	return ""
}

type fieldInfo struct {
	Name string
	Type string
}

func tableExists(conn *sql.Conn, tableName string) (bool, error) {
	tableQuery := "" +
		"SELECT EXISTS (" +
		"    SELECT FROM " +
		"        information_schema.tables " +
		"    WHERE " +
		"        table_schema LIKE 'public' AND " +
		"        table_name = $1" +
		");"

	var exists bool
	err := conn.QueryRowContext(context.Background(), tableQuery, tableName).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if table exists: %v", err)
	}

	return exists, nil
}

func createTableQuery(tableName, fields []fieldInfo) string {
	var fieldsStr []string
	for _, field := range fields {
		fieldsStr = append(fieldsStr, fmt.Sprintf("%s %s", field.Name, field.Type))
	}
	return fmt.Sprintf("CREATE TABLE %s (%s);", tableName, strings.Join(fieldsStr, ", "))
}

func createTable(conn *sql.Conn, tableName string, fields []fieldInfo) error {
	panic("implement create table")
}

func addMissingColumns() error {
	panic("implement add missing columns")
}
