package server

import (
	"context"
	"database/sql"
	"fmt"
	"net"

	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql/postgres/pb"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"

	"google.golang.org/grpc"
)

type PostgresDB struct {
	conn *sql.Conn
	pb.UnimplementedPostgresServer
}

func NewPostgresDB(conn *sql.Conn) *PostgresDB {
	return &PostgresDB{conn: conn}
}

func (p *PostgresDB) GetById(ctx context.Context, params *pb.IdParams) (*pb.RecordResponse, error) {
	filter := map[string]interface{}{
		"id": params.GetId(),
	}
	query := generateReadQuery(params.GetTable(), filter)
	records, err := executeQuery(ctx, query, p.conn, 1)
	if err != nil {
		return nil, err
	}

	return mapToRecord(records[0])
}

func (p *PostgresDB) Find(ctx context.Context, params *pb.FilterParams) (*pb.RecordsResponse, error) {
	filter, err := pkg.ProtoAnyToMap(params.GetFilter())
	query := generateReadQuery(params.GetTable(), filter)
	records, err := executeQuery(ctx, query, p.conn, -1)
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

func StartPostgresServer(host, port string) error {
	server := grpc.NewServer()
	pgConn, err := getPostgresConnection()
	if err != nil {
		return err
	}

	postgres := NewPostgresDB(pgConn)
	pb.RegisterPostgresServer(server, postgres)

	address := fmt.Sprintf("%s:%s", host, port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	logr.DefaultLogger.Infow("gRPC for Postgres server started", "address", address)
	return server.Serve(listener)
}
