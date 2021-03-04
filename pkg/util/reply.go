package util

import (
	"fmt"
	"github.com/ecnuvj/vhoj_rpc/model/base"
)

func PbReplyf(status base.REPLY_STATUS, format string, args ...interface{}) *base.BaseResponse {
	return &base.BaseResponse{
		Status:  status,
		Message: fmt.Sprintf(format, args...),
	}
}
