package handler

import (
	"context"
	"time"

	"github.com/Kintamani/go-skeleton/internal/delivery/grpc/pb"
	"github.com/Kintamani/go-skeleton/internal/usecase"
)

type ExampleHandler struct {
	pb.UnimplementedExampleServiceServer
	useCase *usecase.ExampleUseCase
}

func NewExampleHandler(useCase *usecase.ExampleUseCase) *ExampleHandler {
	return &ExampleHandler{useCase: useCase}
}

func (h *ExampleHandler) FindAll(ctx context.Context, _ *pb.FindAllRequest) (*pb.FindAllResponse, error) {
	examples, err := h.useCase.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	data := make([]*pb.Example, 0, len(examples))
	for _, item := range examples {
		data = append(data, &pb.Example{
			Id:        item.ID,
			Name:      item.Name,
			CreatedAt: item.CreatedAt.Format(time.RFC3339),
		})
	}

	return &pb.FindAllResponse{Data: data}, nil
}
