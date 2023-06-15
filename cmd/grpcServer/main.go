package main

import (
	"database/sql"
	"net"

	"github.com/igor-cotrim/grpc-go/internal/database"
	"github.com/igor-cotrim/grpc-go/internal/pb"
	"github.com/igor-cotrim/grpc-go/internal/service"
	_ "github.com/mattn/go-sqlite3"
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

	gprcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(gprcServer, categoryService)
	reflection.Register(gprcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := gprcServer.Serve(lis); err != nil {
		panic(err)
	}
}
