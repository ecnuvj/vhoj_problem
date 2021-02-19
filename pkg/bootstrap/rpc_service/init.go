package rpc_service

import (
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/rpc_config"
	problem "github.com/ecnuvj/vhoj_problem/pkg/handler"
	"github.com/ecnuvj/vhoj_problem/pkg/sdk/problempb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Init() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", rpc_config.ProblemRpc.Address.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	handler, err := problem.NewProblemHandler()
	if err != nil {
		log.Fatalf("failed to create handler: %v", err)
	}

	s := grpc.NewServer()
	problempb.RegisterProblemServiceServer(s, handler)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
