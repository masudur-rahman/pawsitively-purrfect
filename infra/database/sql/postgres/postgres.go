package postgres

import (
	"context"
	"database/sql"

	isql "github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql/postgres/pb"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"
)

type Database struct {
	ctx    context.Context
	table  string
	id     string
	client pb.PostgresClient
}

func NewDatabase(ctx context.Context, client pb.PostgresClient) Database {
	return Database{
		ctx:    ctx,
		client: client,
	}
}

func (d Database) Table(name string) isql.Database {
	d.table = name
	return d
}

func (d Database) ID(id string) isql.Database {
	d.id = id
	return d
}

func (d Database) FindOne(document interface{}, filter ...interface{}) (bool, error) {
	var err error
	if err = checkIdOrFilterNonEmpty(d.id, filter); err != nil {
		return false, err
	}

	record := new(pb.RecordResponse)

	if filter == nil {
		record, err = d.client.GetById(d.ctx, &pb.IdParams{
			Table: d.table,
			Id:    d.id,
		})
		if err != nil {
			return false, err
		}
	} else {
		af, err := pkg.ToProtoAny(filter[0])
		if err != nil {
			return false, err
		}

		record, err = d.client.Get(d.ctx, &pb.FilterParams{
			Table:  d.table,
			Filter: af,
		})
		if err != nil {
			return false, err
		}
	}

	rmap, err := pkg.ProtoAnyToMap(record.Record)
	if err != nil {
		return false, err
	}

	if err = pkg.ParseInto(rmap, document); err != nil {
		return false, err
	}

	return true, nil
}

func (d Database) FindMany(documents interface{}, filter interface{}) error {
	af, err := pkg.ToProtoAny(filter)
	if err != nil {
		return err
	}

	records, err := d.client.Find(d.ctx, &pb.FilterParams{
		Table:  d.table,
		Filter: af,
	})
	if err != nil {
		return err
	}

	rmaps := make([]map[string]interface{}, 0)
	for _, record := range records.Records {
		rmap, err := pkg.ProtoAnyToMap(record.Record)
		if err != nil {
			return err
		}
		rmaps = append(rmaps, rmap)
	}

	return pkg.ParseInto(rmaps, documents)
}

func (d Database) InsertOne(document interface{}) (id string, err error) {
	df, err := pkg.ToProtoAny(document)
	if err != nil {
		return "", err
	}

	_, err = d.client.Create(d.ctx, &pb.CreateParams{
		Table:  d.table,
		Record: df,
	})
	if err != nil {
		return "", err
	}

	return "", nil
}

func (d Database) InsertMany(documents []interface{}) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (d Database) UpdateOne(document interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (d Database) DeleteOne(filter ...interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (d Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	//TODO implement me
	panic("implement me")
}

func (d Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	//TODO implement me
	panic("implement me")
}
