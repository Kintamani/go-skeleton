package app

import (
	"net"
	"os"

	"github.com/Kintamani/go-skeleton/internal/config"
	"github.com/Kintamani/go-skeleton/internal/delivery/grpc/handler"
	"github.com/Kintamani/go-skeleton/internal/delivery/grpc/pb"
	"github.com/Kintamani/go-skeleton/internal/infrastructure/database"
	"github.com/Kintamani/go-skeleton/internal/infrastructure/logger"
	"github.com/Kintamani/go-skeleton/internal/repository"
	"github.com/Kintamani/go-skeleton/internal/usecase"
	"google.golang.org/grpc"
)

func RunGRPC() {
	log := logger.New(config.ENV.App.Environment)
	db := database.New(log)
	defer db.Close()

	exampleRepository := repository.NewExampleRepository(db)
	exampleUseCase := usecase.NewExampleUseCase(exampleRepository)

	lis, err := net.Listen("tcp", ":"+config.ENV.App.PortGRPC)
	if err != nil {
		log.Error("failed to listen tcp for grpc", "error", err)
		os.Exit(1)
	}

	server := grpc.NewServer()
	pb.RegisterExampleServiceServer(server, handler.NewExampleHandler(exampleUseCase))

	log.Info("grpc server started", "port", config.ENV.App.PortGRPC)
	if err := server.Serve(lis); err != nil {
		log.Error("failed to start grpc server", "error", err)
		os.Exit(1)
	}
}
