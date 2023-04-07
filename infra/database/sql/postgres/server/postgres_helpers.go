package server

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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

func executeQuery(ctx context.Context, query string, conn *sql.Conn, lim int64) ([]map[string]interface{}, error) {
	// Execute query
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// create slice to hold records
	records := make([]map[string]interface{}, 0)

	// iterate over rows and convert to map
	for rows.Next() {
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

func mapToRecord(record map[string]interface{}) (*pb.RecordResponse, error) {
	pm, err := pkg.MapToProtoAny(record)
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
