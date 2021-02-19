package service

import (
	"encoding/json"
	"fmt"
	"github.com/ecnuvj/vhoj_problem/pkg/bootstrap/database"
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
