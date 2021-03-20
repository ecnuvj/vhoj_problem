package adapter

import (
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	"github.com/ecnuvj/vhoj_problem/pkg/sdk/problempb"
	"github.com/ecnuvj/vhoj_rpc/model/userpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
)

func ModelProblemToRpcProblem(problem *model.Problem) *problempb.Problem {
	if problem == nil || problem.RawProblem == nil {
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
	if contest == nil {
		return &problempb.Contest{}
	}
	startTime, _ := ptypes.TimestampProto(contest.StartTime)
	endTime, _ := ptypes.TimestampProto(contest.EndTime)
	problemIds := UintsToUint64s(contest.ProblemIds)
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

func RpcContestToModelContest(contest *problempb.Contest) *model.Contest {
	if contest == nil {
		return &model.Contest{}
	}
	return &model.Contest{
		Model:       gorm.Model{ID: uint(contest.ContestId)},
		Title:       contest.Title,
		Description: contest.Description,
		UserId:      uint(contest.Creator.UserId),
		User:        RpcUserToModelUser(contest.Creator),
		ProblemIds:  Uint64sToUints(contest.ProblemIds),
		StartTime:   contest.StartTime.AsTime(),
		EndTime:     contest.EndTime.AsTime(),
	}
}

func ModelUserToRpcUser(user *model.User) *userpb.User {
	if user == nil {
		return &userpb.User{}
	}
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

func RpcUserToModelUser(user *userpb.User) *model.User {
	if user == nil {
		return &model.User{}
	}
	roles := RpcRolesToModelRoles(user.Roles)
	userAuth := &model.UserAuth{
		Model:    gorm.Model{ID: uint(user.UserAuthId)},
		Password: user.Password,
	}
	return &model.User{
		Model: gorm.Model{
			ID: uint(user.UserId),
		},
		UserAuth:  userAuth,
		Email:     user.Email,
		Nickname:  user.Username,
		School:    user.School,
		Roles:     roles,
		Submitted: user.Submitted,
		Accepted:  user.Accepted,
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

func RpcRolesToModelRoles(roles []*userpb.Role) []*model.Role {
	retRoles := make([]*model.Role, len(roles))
	for i, r := range roles {
		retRoles[i] = &model.Role{
			Model:    gorm.Model{ID: uint(r.RoleId)},
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

func Uint64sToUints(ids []uint64) []uint {
	ret := make([]uint, len(ids))
	for i, id := range ids {
		ret[i] = uint(id)
	}
	return ret
}

func RpcContestProblemToModelContestProblem(problem *problempb.ContestProblem) *model.ContestProblem {
	return &model.ContestProblem{
		ContestId:    uint(problem.ContestId),
		ProblemOrder: problem.ProblemOrder,
		ProblemId:    uint(problem.ProblemId),
		Title:        problem.Title,
		Accepted:     uint(problem.Accepted),
		Submitted:    uint(problem.Submitted),
	}
}

func RpcContestProblemsToModelContestProblems(problems []*problempb.ContestProblem) []*model.ContestProblem {
	retProblems := make([]*model.ContestProblem, len(problems))
	for i, p := range problems {
		retProblems[i] = RpcContestProblemToModelContestProblem(p)
	}
	return retProblems
}

func ModelContestProblemToRpcContestProblem(problem *model.ContestProblem) *problempb.ContestProblem {
	return &problempb.ContestProblem{
		ContestId:    uint64(problem.ContestId),
		ProblemId:    uint64(problem.ProblemId),
		ProblemOrder: problem.ProblemOrder,
		Title:        problem.Title,
		Accepted:     uint64(problem.Accepted),
		Submitted:    uint64(problem.Submitted),
	}
}

func ModelContestProblemsToRpcContestProblems(problems []*model.ContestProblem) []*problempb.ContestProblem {
	retProblems := make([]*problempb.ContestProblem, len(problems))
	for i, p := range problems {
		retProblems[i] = ModelContestProblemToRpcContestProblem(p)
	}
	return retProblems
}
