package service

import (
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/contest_mapper"
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/problem_mapper"
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	"github.com/ecnuvj/vhoj_problem/pkg/adapter"
	"github.com/ecnuvj/vhoj_problem/pkg/common"
	"github.com/ecnuvj/vhoj_problem/pkg/sdk/problempb"
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
	problem, err := problem_mapper.ProblemMapper.GetProblemById(problemId)
	if err != nil {
		return nil, err
	}
	return adapter.ModelProblemToRpcProblem(problem), nil
}

func (ProblemService) SearchProblem(condition *common.SearchCondition, pageNo int32, pageSize int32) ([]*problempb.Problem, *common.PageInfo, error) {
	if (condition == nil) || (condition.ProblemId == 0 && condition.Title == "") {
		return nil, &common.PageInfo{}, nil
	}
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

func (ProblemService) CreateContest(contest *problempb.Contest) (*model.Contest, error) {
	problemIds := make([]uint, len(contest.ProblemIds))
	for i, p := range contest.ProblemIds {
		problemIds[i] = uint(p)
	}
	modelContest := &model.Contest{
		Title:       contest.Title,
		Description: contest.Description,
		UserId:      uint(contest.Creator.UserId),
		ProblemNum:  int64(len(contest.ProblemIds)),
		ProblemIds:  problemIds,
		StartTime:   contest.StartTime.AsTime(),
		EndTime:     contest.EndTime.AsTime(),
	}
	retContest, err := contest_mapper.ContestMapper.CreateContest(modelContest)
	if err != nil {
		return nil, err
	}
	return retContest, nil
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
