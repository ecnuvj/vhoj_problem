package common

import "github.com/ecnuvj/vhoj_common/pkg/common/constants/contest_status"

type ProblemSearchCondition struct {
	Title     string
	ProblemId uint
}

type ContestSearchCondition struct {
	Status      contest_status.ContestStatus
	Title       string
	CreatorName string
}
