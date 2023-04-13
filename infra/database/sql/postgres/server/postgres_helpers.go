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
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"

	"github.com/iancoleman/strcase"
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

func isDefaultValue(value interface{}) bool {
	typ := reflect.TypeOf(value)
	zero := reflect.Zero(typ).Interface()
	return reflect.DeepEqual(value, zero)
}

func toDBCase(fieldName string) string {
	return strcase.ToSnake(fieldName)
}

func fromDBCase(fieldName string) string {
	return strcase.ToLowerCamel(fieldName)
}

func generateReadQuery(tableName string, filter map[string]interface{}) string {
	var conditions []string

	for key, value := range filter {
		if isDefaultValue(value) {
			// don't insert the default value checks into the condition array
			continue
		}

		key = toDBCase(key)
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
	query := fmt.Sprintf("SELECT * FROM \"%s\" WHERE %s", tableName, conditionString)

	return query
}

func scanSingleRecord(rows *sql.Rows) (map[string]interface{}, error) {
	fields, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	scans := make([]interface{}, len(fields))

	for i := range scans {
		scans[i] = &scans[i]
	}
	if err = rows.Scan(scans...); err != nil {
		return nil, err
	}

	record := make(map[string]interface{})
	for i := range scans {
		fieldName := fromDBCase(fields[i])
		record[fieldName] = scans[i]
	}

	return record, nil
}

func executeReadQuery(ctx context.Context, query string, conn *sql.Conn, lim int64) ([]map[string]interface{}, error) {
	logr.DefaultLogger.Infow("Read Query", "query", query)
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
	var values []string

	for col, val := range record {
		//if isDefaultValue(val) {
		//	// don't need to insert the default values into the table
		//	continue
		//}

		col = toDBCase(col)
		cols = append(cols, col)
		switch v := val.(type) {
		case string:
			values = append(values, fmt.Sprintf("'%s'", strings.ReplaceAll(v, "'", "''")))
		case time.Time:
			values = append(values, fmt.Sprintf("'%s'", v.Format("2006-01-02 15:04:05")))
		default:
			values = append(values, fmt.Sprintf("%v", v))
		}
	}

	query := fmt.Sprintf("INSERT INTO \"%s\" (%s) VALUES (%s)", tableName, strings.Join(cols, ", "), strings.Join(values, ", "))

	return query
}

func executeWriteQuery(ctx context.Context, query string, conn *sql.Conn) (sql.Result, error) {
	logr.DefaultLogger.Infow("Write Query", "query", query)
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
