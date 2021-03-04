package adapter

import (
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	"github.com/ecnuvj/vhoj_problem/pkg/sdk/problempb"
	"github.com/ecnuvj/vhoj_rpc/model/userpb"
	"github.com/golang/protobuf/ptypes"
)

func ModelProblemToRpcProblem(problem *model.Problem) *problempb.Problem {
	if problem.RawProblem == nil {
		return &problempb.Problem{}
	}
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

func ModelProblemsToRpcProblems(problems []*model.Problem) []*problempb.Problem {
	retProblems := make([]*problempb.Problem, len(problems))
	for i, p := range problems {
		retProblems[i] = ModelProblemToRpcProblem(p)
	}
	return retProblems
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

func ModelUserToRpcUser(user *model.User) *userpb.User {
	var userAuthId uint64
	var password string
	if user.UserAuth != nil {
		userAuthId = uint64(user.UserAuth.ID)
		password = user.UserAuth.Password
	}
	return &userpb.User{
		UserId:     uint64(user.ID),
		Username:   user.Nickname,
		UserAuthId: userAuthId,
		Password:   password,
		Email:      user.Email,
		School:     user.School,
		Roles:      ModelRolesToRpcRoles(user.Roles),
		Accepted:   user.Accepted,
		Submitted:  user.Submitted,
	}
}

func ModelRolesToRpcRoles(roles []*model.Role) []*userpb.Role {
	retRoles := make([]*userpb.Role, len(roles))
	for i, r := range roles {
		retRoles[i] = &userpb.Role{
			RoleId:   uint64(r.ID),
			RoleName: r.RoleName,
		}
	}
	return retRoles
}

func UintsToUint64s(ids []uint) []uint64 {
	ret := make([]uint64, len(ids))
	for i, id := range ids {
		ret[i] = uint64(id)
	}
	return ret
}

func Uint64sToUints(ids []int64) []uint {
	ret := make([]uint, len(ids))
	for i, id := range ids {
		ret[i] = uint(id)
	}
	return ret
}
