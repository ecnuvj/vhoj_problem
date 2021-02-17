package handler

import (
	"context"
	"github.com/ecnuvj/vhoj_problem/pkg/sdk/problempb"
	"github.com/ecnuvj/vhoj_problem/pkg/service"
)

type ProblemHandler struct {
	problempb.UnimplementedProblemServiceServer
	ProblemService *service.ProblemService
}

func NewProblemHandler() (*ProblemHandler,error) {
	return &ProblemHandler{
		ProblemService: &service.ProblemService{},
	},nil
}

func (p *ProblemHandler) ListProblems(ctx context.Context, request *problempb.ListProblemsRequest) (*problempb.ListProblemsResponse, error) {
	panic("implement me")
}

func (p *ProblemHandler) GetProblemById(ctx context.Context, request *problempb.GetProblemByIdRequest) (*problempb.GetProblemByIdResponse, error) {
	panic("implement me")
}

func (p *ProblemHandler) SearchProblemByCondition(ctx context.Context, request *problempb.SearchProblemByConditionRequest) (*problempb.SearchProblemByConditionResponse, error) {
	panic("implement me")
}

func (p *ProblemHandler) CreateContest(ctx context.Context, request *problempb.CreateContestRequest) (*problempb.CreateContestResponse, error) {
	panic("implement me")
}

func (p *ProblemHandler) ListContests(ctx context.Context, request *problempb.ListContestsRequest) (*problempb.ListContestsResponse, error) {
	panic("implement me")
}

func (p *ProblemHandler) GetContestById(ctx context.Context, request *problempb.GetContestByIdRequest) (*problempb.GetContestByIdResponse, error) {
	panic("implement me")
}

func (p *ProblemHandler) AddContestParticipant(ctx context.Context, request *problempb.AddContestParticipantRequest) (*problempb.AddContestParticipantResponse, error) {
	panic("implement me")
}

func (p *ProblemHandler) GenerateContestParticipants(ctx context.Context, request *problempb.GenerateContestParticipantsRequest) (*problempb.GenerateContestParticipantsResponse, error) {
	panic("implement me")
}

func (p *ProblemHandler) AddContestAdmin(ctx context.Context, request *problempb.AddContestAdminRequest) (*problempb.AddContestAdminResponse, error) {
	panic("implement me")
}

func (p *ProblemHandler) GetContestAdmins(ctx context.Context, request *problempb.GetContestAdminsRequest) (*problempb.GetContestAdminsResponse, error) {
	panic("implement me")
}

func (p *ProblemHandler) GetContestParticipants(ctx context.Context, request *problempb.GetContestParticipantsRequest) (*problempb.GetContestParticipantsResponse, error) {
	panic("implement me")
}

