package server

import (
	"context"

	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql/postgres/pb"
)

type PostgresDB struct {
	pb.UnimplementedPostgresServer
}

func (p *PostgresDB) GetById(ctx context.Context, req *pb.GetByIdRequest) (*pb.GetByIdResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresDB) Find(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresDB) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresDB) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresDB) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresDB) Query(ctx context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresDB) Exec(ctx context.Context, req *pb.ExecRequest) (*pb.ExecResponse, error) {
	//TODO implement me
	panic("implement me")
}
