package grpcProduct

import (
	"fmt"
	"log"
	"net"

	"github.com/TIM-DEBUG-ProjectSprintBatch3/go-fiber-template/src/config"
	"github.com/TIM-DEBUG-ProjectSprintBatch3/go-fiber-template/src/di"
	protoProductController "github.com/TIM-DEBUG-ProjectSprintBatch3/go-fiber-template/src/grpc/controller/product/proto"
	"github.com/TIM-DEBUG-ProjectSprintBatch3/go-fiber-template/src/service/proto/product"
	"github.com/samber/do/v2"
	"google.golang.org/grpc"
)

type GrpcServer struct{}

func Listen() {
	_PORT := config.GetPortGrpc()

	// create new grpc server handler
	grpcServer := grpc.NewServer()

	// Registrasikan service gRPC Anda
	ppc := do.MustInvoke[*protoProductController.ProtoProductController](di.Injector)
	product.RegisterProductServiceServer(grpcServer, ppc)

	go func() {
		// create tcp server
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", _PORT))

		if err != nil {
			log.Fatalf("Failed to Listen on Port: %v", err)
		}

		fmt.Printf("> gRPC server listening on :%s\n", _PORT)

		// run grpc server
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
}
