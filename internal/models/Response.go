package models

import (
	"team01/internal/proto/node"
)

func ErrorResponse(text string) *node.NodeResp {
	return &node.NodeResp{
		Value: text,
		Code:  2,
		//Result:  map[string]bool{cfg.GetAddress(): false},

	}
}

func OkResponse(text string) *node.NodeResp {
	return &node.NodeResp{
		Value: text,
		Code:  1,
	}
}
