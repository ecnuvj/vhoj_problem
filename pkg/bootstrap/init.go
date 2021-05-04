package bootstrap

import (
	"github.com/ecnuvj/vhoj_problem/pkg/bootstrap/database"
	"github.com/ecnuvj/vhoj_problem/pkg/bootstrap/rpc_client"
	"github.com/ecnuvj/vhoj_problem/pkg/bootstrap/rpc_service"
)

func Init() {
	database.Init()
	rpc_service.Init()
	rpc_client.Init()
}
