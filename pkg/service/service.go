package service

import (
	"context"
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/contest_mapper"
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/problem_mapper"
	"github.com/ecnuvj/vhoj_problem/pkg/adapter"
	"github.com/ecnuvj/vhoj_problem/pkg/common"
	"github.com/ecnuvj/vhoj_problem/pkg/sdk/problempb"
	"github.com/ecnuvj/vhoj_rpc/client/rpc_user"
	"github.com/ecnuvj/vhoj_rpc/model/userpb"
)

type ProblemService struct {
}

func (ProblemService) ListProblems(pageNo int32, pageSize int32) ([]*problempb.Problem, *common.PageInfo, error) {
	problems, count, err := problem_mapper.ProblemMapper.FindAllProblems(pageNo, pageSize)
	if err != nil {
		return nil, nil, err
	}
	rpcProblems := make([]*problempb.Problem, len(problems))
	for index, pro := range problems {
		rpcProblems[index] = adapter.ModelProblemToRpcProblem(pro)
	}
	return rpcProblems, &common.PageInfo{
		TotalPages: (count + pageSize - 1) / pageSize,
		TotalCount: count,
	}, nil
}

func (ProblemService) GetProblem(problemId uint) (*problempb.Problem, error) {
	problem, err := problem_mapper.ProblemMapper.FindProblemById(problemId)
	if err != nil {
		return nil, err
	}
	return adapter.ModelProblemToRpcProblem(problem), nil
}

func (ProblemService) SearchProblem(condition *common.ProblemSearchCondition, pageNo int32, pageSize int32) ([]*problempb.Problem, *common.PageInfo, error) {
	problems, count, err := problem_mapper.ProblemMapper.SearchProblemByCondition(&problem_mapper.ProblemSearchParam{
		Title:     condition.Title,
		ProblemId: condition.ProblemId,
	}, pageNo, pageSize)
	if err != nil {
		return nil, &common.PageInfo{}, err
	}
	rpcProblems := make([]*problempb.Problem, len(problems))
	for index, problem := range problems {
		rpcProblems[index] = adapter.ModelProblemToRpcProblem(problem)
	}
	return rpcProblems, &common.PageInfo{
		TotalPages: (count + pageSize - 1) / pageSize,
		TotalCount: count,
	}, nil
}

func (ProblemService) CreateContest(contest *problempb.Contest) (*problempb.Contest, error) {
	modelContest := adapter.RpcContestToModelContest(contest)
	retContest, err := contest_mapper.ContestMapper.CreateContest(modelContest)
	if err != nil {
		return nil, err
	}
	return adapter.ModelContestToRpcContest(retContest), nil
}

func (ProblemService) ListContests(pageNo int32, pageSize int32) ([]*problempb.Contest, *common.PageInfo, error) {
	contests, count, err := contest_mapper.ContestMapper.FindAllContests(pageNo, pageSize)
	if err != nil {
		return nil, &common.PageInfo{}, err
	}
	rpcContests := make([]*problempb.Contest, len(contests))
	for i, c := range contests {
		rpcContests[i] = adapter.ModelContestToRpcContest(c)
	}
	return rpcContests, &common.PageInfo{
		TotalPages: (count + pageSize - 1) / pageSize,
		TotalCount: count,
	}, nil
}

func (ProblemService) SearchContest(condition *common.ContestSearchCondition, pageNo int32, pageSize int32) ([]*problempb.Contest, *common.PageInfo, error) {
	contests, count, err := contest_mapper.ContestMapper.FindContestsByCondition(&contest_mapper.SearchContestCondition{
		Status:      condition.Status,
		Title:       condition.Title,
		CreatorName: condition.CreatorName,
	}, pageNo, pageSize)
	if err != nil {
		return nil, &common.PageInfo{}, err
	}
	rpcContests := make([]*problempb.Contest, len(contests))
	for i, c := range contests {
		rpcContests[i] = adapter.ModelContestToRpcContest(c)
	}
	return rpcContests, &common.PageInfo{
		TotalPages: (count + pageSize - 1) / pageSize,
		TotalCount: count,
	}, nil
}

func (ProblemService) GetContest(contestId uint) (*problempb.Contest, error) {
	contest, err := contest_mapper.ContestMapper.FindContestById(contestId)
	if err != nil {
		return nil, err
	}
	rpcContest := adapter.ModelContestToRpcContest(contest)
	problems, _ := problem_mapper.ProblemMapper.FindProblemsByIds(contest.ProblemIds)
	rpcContest.Problems = adapter.ModelProblemsToRpcProblems(problems)
	return rpcContest, nil
}

func (ProblemService) JoinContest(contestId uint, userIds []uint) error {
	err := contest_mapper.ContestMapper.AddContestParticipants(contestId, userIds)
	if err != nil {
		return err
	}
	return nil
}

func (ProblemService) AddContestAdmins(contestId uint, userIds []uint) error {
	err := contest_mapper.ContestMapper.AddContestAdmins(contestId, userIds)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProblemService) GenerateContestParticipants(contestId uint, count int32) ([]*userpb.User, error) {
	request := &userpb.GenerateUsersRequest{
		GenerateCount: count,
		ContestId:     uint64(contestId),
	}
	resp, err := rpc_user.UserServiceClient.GenerateUsers(context.Background(), request)
	if err != nil {
		return nil, err
	}
	userIds := make([]uint, len(resp.Users))
	for i, user := range resp.Users {
		userIds[i] = uint(user.UserId)
	}
	err = p.JoinContest(contestId, userIds)
	if err != nil {
		return nil, err
	}

	return resp.Users, nil
}

func (ProblemService) GetContestAdmins(contestId uint) ([]*userpb.User, error) {
	userIds, err := contest_mapper.ContestMapper.FindContestAdmins(contestId)
	if err != nil {
		return nil, err
	}

	request := &userpb.GetUsersByIdsRequest{
		UserIds: adapter.UintsToUint64s(userIds),
	}
	resp, err := rpc_user.UserServiceClient.GetUsersByIds(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return resp.Users, nil
}

func (ProblemService) GetContestParticipants(contestId uint) ([]*userpb.User, error) {
	userIds, err := contest_mapper.ContestMapper.FindContestParticipants(contestId)
	if err != nil {
		return nil, err
	}
	request := &userpb.GetUsersByIdsRequest{
		UserIds: adapter.UintsToUint64s(userIds),
	}
	resp, err := rpc_user.UserServiceClient.GetUsersByIds(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return resp.Users, nil
}

func (ProblemService) AddContestProblem(contestId uint, problemId uint) error {
	err := contest_mapper.ContestMapper.AddContestProblem(contestId, problemId)
	if err != nil {
		return err
	}
	return nil
}

func (ProblemService) DeleteContestProblem(contestId uint, problemId uint) error {
	err := contest_mapper.ContestMapper.DeleteContestProblem(contestId, problemId)
	if err != nil {
		return err
	}
	return nil
}

func (ProblemService) DeleteContestAdmin(contestId uint, userId uint) error {
	err := contest_mapper.ContestMapper.DeleteContestAdmin(contestId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (ProblemService) UpdateContest(contest *problempb.Contest) (*problempb.Contest, error) {
	retContest, err := contest_mapper.ContestMapper.UpdateContest(adapter.RpcContestToModelContest(contest))
	if err != nil {
		return nil, err
	}
	return adapter.ModelContestToRpcContest(retContest), nil
}
