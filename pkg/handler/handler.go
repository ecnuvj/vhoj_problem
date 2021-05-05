package handler

import (
	"context"
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/contest_status"
	"github.com/ecnuvj/vhoj_problem/pkg/common"
	"github.com/ecnuvj/vhoj_problem/pkg/sdk/problempb"
	"github.com/ecnuvj/vhoj_problem/pkg/service"
	"github.com/ecnuvj/vhoj_problem/pkg/util"
	"github.com/ecnuvj/vhoj_rpc/model/base"
)

type ProblemHandler struct {
	problempb.UnimplementedProblemServiceServer
	ProblemService *service.ProblemService
}

func NewProblemHandler() (*ProblemHandler, error) {
	return &ProblemHandler{
		ProblemService: &service.ProblemService{},
	}, nil
}

func (p *ProblemHandler) ListProblems(ctx context.Context, request *problempb.ListProblemsRequest) (*problempb.ListProblemsResponse, error) {
	if request == nil {
		return &problempb.ListProblemsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "list problem request is nil"),
		}, fmt.Errorf("list problem request is nil")
	}
	rpcProblems, pageInfo, err := p.ProblemService.ListProblems(request.PageNo, request.PageSize)
	if err != nil {
		return &problempb.ListProblemsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, err
	}
	return &problempb.ListProblemsResponse{
		Problems:     rpcProblems,
		TotalPages:   pageInfo.TotalPages,
		TotalCount:   pageInfo.TotalCount,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) GetProblemById(ctx context.Context, request *problempb.GetProblemByIdRequest) (*problempb.GetProblemByIdResponse, error) {
	if request == nil {
		return &problempb.GetProblemByIdResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "get problem request is nil"),
		}, fmt.Errorf("get problem request is nil")
	}
	rpcProblem, err := p.ProblemService.GetProblem(uint(request.ProblemId))
	if err != nil {
		return &problempb.GetProblemByIdResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, err
	}
	return &problempb.GetProblemByIdResponse{
		Problem:      rpcProblem,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) SearchProblemByCondition(ctx context.Context, request *problempb.SearchProblemByConditionRequest) (*problempb.SearchProblemByConditionResponse, error) {
	if request == nil {
		return &problempb.SearchProblemByConditionResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "search problem request is nil"),
		}, fmt.Errorf("search problem request is nil")
	}
	condition := &common.ProblemSearchCondition{
		Title:     request.Title,
		ProblemId: uint(request.ProblemId),
	}
	rpcProblems, pageInfo, err := p.ProblemService.SearchProblem(condition, request.PageNo, request.PageSize)
	if err != nil {
		return &problempb.SearchProblemByConditionResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, err
	}
	return &problempb.SearchProblemByConditionResponse{
		Problems:     rpcProblems,
		TotalPages:   pageInfo.TotalPages,
		TotalCount:   pageInfo.TotalCount,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) CreateContest(ctx context.Context, request *problempb.CreateContestRequest) (*problempb.CreateContestResponse, error) {
	if request == nil {
		return &problempb.CreateContestResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "create contest request is nil"),
		}, fmt.Errorf("create contest request is nil")
	}
	contest, err := p.ProblemService.CreateContest(request.Contest, request.ContestProblems)
	if err != nil {
		return &problempb.CreateContestResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.CreateContestResponse{
		Contest:      contest,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) ListContests(ctx context.Context, request *problempb.ListContestsRequest) (*problempb.ListContestsResponse, error) {
	if request == nil {
		return &problempb.ListContestsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "list contest request is nil"),
		}, fmt.Errorf("list contest request is nil")
	}
	rpcContests, pageInfo, err := p.ProblemService.ListContests(request.PageNo, request.PageSize)
	if err != nil {
		return &problempb.ListContestsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.ListContestsResponse{
		Contests:     rpcContests,
		TotalPages:   pageInfo.TotalPages,
		TotalCount:   pageInfo.TotalCount,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) SearchContestByCondition(ctx context.Context, request *problempb.SearchContestByConditionRequest) (*problempb.SearchContestByConditionResponse, error) {
	if request == nil {
		return &problempb.SearchContestByConditionResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	condition := &common.ContestSearchCondition{
		Status:      contest_status.ContestStatus(request.Status),
		Title:       request.Title,
		CreatorName: request.CreatorName,
	}
	contests, pageInfo, err := p.ProblemService.SearchContest(condition, request.PageNo, request.PageSize)
	if err != nil {
		return &problempb.SearchContestByConditionResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.SearchContestByConditionResponse{
		Contests:     contests,
		TotalPages:   pageInfo.TotalPages,
		TotalCount:   pageInfo.TotalCount,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) GetContestById(ctx context.Context, request *problempb.GetContestByIdRequest) (*problempb.GetContestByIdResponse, error) {
	if request == nil {
		return &problempb.GetContestByIdResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "get contest request is nil"),
		}, fmt.Errorf("get contest request is nil")
	}
	rpcContest, err := p.ProblemService.GetContest(uint(request.ContestId))
	if err != nil {
		return &problempb.GetContestByIdResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.GetContestByIdResponse{
		Contest:      rpcContest,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) AddContestParticipant(ctx context.Context, request *problempb.AddContestParticipantRequest) (*problempb.AddContestParticipantResponse, error) {
	if request == nil {
		return &problempb.AddContestParticipantResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "join contest request is nil"),
		}, fmt.Errorf("join contest request is nil")
	}
	err := p.ProblemService.JoinContest(uint(request.ContestId), []uint{uint(request.UserId)})
	if err != nil {
		return &problempb.AddContestParticipantResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.AddContestParticipantResponse{
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) GenerateContestParticipants(ctx context.Context, request *problempb.GenerateContestParticipantsRequest) (*problempb.GenerateContestParticipantsResponse, error) {
	if request == nil {
		return &problempb.GenerateContestParticipantsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "generate contest user request is nil"),
		}, fmt.Errorf("generate contest user request is nil")
	}
	users, err := p.ProblemService.GenerateContestParticipants(uint(request.ContestId), request.GenerateCount)
	if err != nil {
		return &problempb.GenerateContestParticipantsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.GenerateContestParticipantsResponse{
		Users:        users,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) AddContestAdmin(ctx context.Context, request *problempb.AddContestAdminRequest) (*problempb.AddContestAdminResponse, error) {
	if request == nil {
		return &problempb.AddContestAdminResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "add contest admin request is nil"),
		}, fmt.Errorf("add contest admin request is nil")
	}
	err := p.ProblemService.AddContestAdmins(uint(request.ContestId), []uint{uint(request.UserId)})
	if err != nil {
		return &problempb.AddContestAdminResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.AddContestAdminResponse{
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) GetContestAdmins(ctx context.Context, request *problempb.GetContestAdminsRequest) (*problempb.GetContestAdminsResponse, error) {
	if request == nil {
		return &problempb.GetContestAdminsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "get contest admin request is nil"),
		}, fmt.Errorf("get contest admin request is nil")
	}
	users, err := p.ProblemService.GetContestAdmins(uint(request.ContestId))
	if err != nil {
		return &problempb.GetContestAdminsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.GetContestAdminsResponse{
		Users:        users,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) GetContestParticipants(ctx context.Context, request *problempb.GetContestParticipantsRequest) (*problempb.GetContestParticipantsResponse, error) {
	if request == nil {
		return &problempb.GetContestParticipantsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "get contest participants request is nil"),
		}, fmt.Errorf("get contest participants request is nil")
	}
	users, err := p.ProblemService.GetContestParticipants(uint(request.ContestId))
	if err != nil {
		return &problempb.GetContestParticipantsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.GetContestParticipantsResponse{
		Users:        users,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) AddContestProblem(ctx context.Context, request *problempb.AddContestProblemRequest) (*problempb.AddContestProblemResponse, error) {
	if request == nil {
		return &problempb.AddContestProblemResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	err := p.ProblemService.AddContestProblem(uint(request.ContestId), uint(request.ProblemId))
	if err != nil {
		return &problempb.AddContestProblemResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.AddContestProblemResponse{
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) DeleteContestProblem(ctx context.Context, request *problempb.DeleteContestProblemRequest) (*problempb.DeleteContestProblemResponse, error) {
	if request == nil {
		return &problempb.DeleteContestProblemResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	err := p.ProblemService.DeleteContestProblem(uint(request.ContestId), uint(request.ProblemId))
	if err != nil {
		return &problempb.DeleteContestProblemResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.DeleteContestProblemResponse{
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) DeleteContestAdmin(ctx context.Context, request *problempb.DeleteContestAdminRequest) (*problempb.DeleteContestAdminResponse, error) {
	if request == nil {
		return &problempb.DeleteContestAdminResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	err := p.ProblemService.DeleteContestAdmin(uint(request.ContestId), uint(request.UserId))
	if err != nil {
		return &problempb.DeleteContestAdminResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.DeleteContestAdminResponse{
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) UpdateContest(ctx context.Context, request *problempb.UpdateContestRequest) (*problempb.UpdateContestResponse, error) {
	if request == nil {
		return &problempb.UpdateContestResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	contest, err := p.ProblemService.UpdateContest(request.Contest)
	if err != nil {
		return &problempb.UpdateContestResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.UpdateContestResponse{
		Contest:      contest,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) UpdateContestProblems(ctx context.Context, request *problempb.UpdateContestProblemsRequest) (*problempb.UpdateContestProblemsResponse, error) {
	if request == nil {
		return &problempb.UpdateContestProblemsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	problems, err := p.ProblemService.UpdateContestProblems(uint(request.ContestId), request.ContestProblems)
	if err != nil {
		return &problempb.UpdateContestProblemsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.UpdateContestProblemsResponse{
		ContestProblems: problems,
		BaseResponse:    util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

//获取比赛题目order和id
func (p *ProblemHandler) GetContestProblems(ctx context.Context, request *problempb.GetContestProblemsRequest) (*problempb.GetContestProblemsResponse, error) {
	if request == nil {
		return &problempb.GetContestProblemsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	problems, err := p.ProblemService.GetContestProblems(uint(request.ContestId))
	if err != nil {
		return &problempb.GetContestProblemsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.GetContestProblemsResponse{
		ContestProblems: problems,
		BaseResponse:    util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) GetUserContests(ctx context.Context, request *problempb.GetUserContestsRequest) (*problempb.GetUserContestsResponse, error) {
	if request == nil {
		return &problempb.GetUserContestsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	contests, pageInfo, err := p.ProblemService.GetUserContests(uint(request.UserId), request.PageNo, request.PageSize)
	if err != nil {
		return &problempb.GetUserContestsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.GetUserContestsResponse{
		Contests:     contests,
		TotalPages:   pageInfo.TotalPages,
		TotalCount:   pageInfo.TotalCount,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) RandProblem(ctx context.Context, request *problempb.RandProblemRequest) (*problempb.RandProblemResponse, error) {
	if request == nil {
		return &problempb.RandProblemResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	problemId, err := p.ProblemService.RandProblem()
	if err != nil {
		return &problempb.RandProblemResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.RandProblemResponse{
		ProblemId:    uint64(problemId),
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (p *ProblemHandler) RawProblemList(ctx context.Context, request *problempb.RawProblemListRequest) (*problempb.RawProblemListResponse, error) {
	if request == nil {
		return &problempb.RawProblemListResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	rawProblems, pageInfo, err := p.ProblemService.RawProblemList(request.PageNo, request.PageSize)
	if err != nil {
		return &problempb.RawProblemListResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &problempb.RawProblemListResponse{
		RawProblems:  rawProblems,
		TotalPages:   pageInfo.TotalPages,
		TotalCount:   pageInfo.TotalCount,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}
