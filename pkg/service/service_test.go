package service

import (
	"encoding/json"
	"fmt"
	"github.com/ecnuvj/vhoj_problem/pkg/bootstrap/database"
	"github.com/ecnuvj/vhoj_problem/pkg/bootstrap/rpc_client"
	"testing"
)

func TestProblemService_GetContest(t *testing.T) {
	database.Init()
	service := &ProblemService{}
	contest, err := service.GetContest(4)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	str, _ := json.Marshal(contest)
	fmt.Println(string(str))
}

func TestProblemService_GenerateContestParticipants(t *testing.T) {
	database.Init()
	rpc_client.Init()
	service := &ProblemService{}
	users, err := service.GenerateContestParticipants(5, 5)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	str, _ := json.Marshal(users)
	fmt.Println(string(str))
}
