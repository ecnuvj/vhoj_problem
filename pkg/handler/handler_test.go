package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ecnuvj/vhoj_problem/pkg/bootstrap/database"
	"github.com/ecnuvj/vhoj_problem/pkg/sdk/problempb"
	"testing"
)

func TestProblemHandler_ListProblems(t *testing.T) {
	database.Init()
	handler, _ := NewProblemHandler()
	resp, err := handler.ListProblems(context.Background(), &problempb.ListProblemsRequest{
		PageNo:   1,
		PageSize: 5,
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	str, _ := json.Marshal(resp)
	fmt.Println(string(str))
}
