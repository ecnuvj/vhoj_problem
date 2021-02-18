package adapter

import (
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	"github.com/ecnuvj/vhoj_problem/pkg/sdk/problempb"
	"github.com/golang/protobuf/ptypes"
)

func ModelProblemToRpcProblem(problem *model.Problem) *problempb.Problem {
	return &problempb.Problem{
		ProblemId:    uint64(problem.ID),
		Title:        problem.RawProblem.Title,
		Description:  problem.RawProblem.Description,
		SampleInput:  problem.RawProblem.SampleInput,
		SampleOutput: problem.RawProblem.SampleOutput,
		Input:        problem.RawProblem.Input,
		Output:       problem.RawProblem.Output,
		TimeLimit:    problem.RawProblem.TimeLimit,
		MemoryLimit:  problem.RawProblem.MemoryLimit,
		Submitted:    problem.Submitted,
		Accepted:     problem.Accepted,
	}
}

func ModelContestToRpcContest(contest *model.Contest) *problempb.Contest {
	startTime, _ := ptypes.TimestampProto(contest.StartTime)
	endTime, _ := ptypes.TimestampProto(contest.EndTime)
	problemIds := make([]uint64, len(contest.ProblemIds))
	for i, p := range contest.ProblemIds {
		problemIds[i] = uint64(p)
	}
	return &problempb.Contest{
		ContestId:   uint64(contest.ID),
		Title:       contest.Title,
		Description: contest.Description,
		StartTime:   startTime,
		EndTime:     endTime,
		Creator:     ModelUserToRpcUser(contest.User),
		ProblemIds:  problemIds,
	}
}

func ModelUserToRpcUser(user *model.User) *problempb.User {
	return &problempb.User{
		UserId:   uint64(user.ID),
		Username: user.Nickname,
	}
}
