package server

import (
	"context"

	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql/postgres/pb"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"
)

type PostgresDB struct {
	pb.UnimplementedPostgresServer
}

func (p *PostgresDB) GetById(ctx context.Context, params *pb.IdParams) (*pb.RecordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresDB) Find(ctx context.Context, params *pb.FilterParams) (*pb.RecordsResponse, error) {
	conn := getPostgresConnection()
	filter, err := pkg.ProtoAnyToMap(params.Filter)
	query := generateReadQuery(params.Table, filter)
	records, err := executeQuery(ctx, query, conn)
	if err != nil {
		return nil, err
	}

	return mapsToRecords(records)
}

func (p *PostgresDB) Create(ctx context.Context, params *pb.CreateParams) (*pb.RecordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresDB) Update(ctx context.Context, params *pb.UpdateParams) (*pb.RecordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresDB) Delete(ctx context.Context, params *pb.IdParams) (*pb.DeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresDB) Query(ctx context.Context, params *pb.QueryParams) (*pb.QueryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresDB) Exec(ctx context.Context, params *pb.ExecParams) (*pb.ExecResponse, error) {
	//TODO implement me
	panic("implement me")
}
