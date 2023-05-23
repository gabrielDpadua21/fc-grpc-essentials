package main

import (
	"database/sql"
	"net"

	"github.com/gabrielDpadua21/fc-grpc-essentialst/internals/database"
	"github.com/gabrielDpadua21/fc-grpc-essentialst/internals/pb"
	"github.com/gabrielDpadua21/fc-grpc-essentialst/internals/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
