package mergeReq

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyCloseMRReq struct {
	Action  string `json:"Action"`  // ModifyCloseMR
	DepotID int64  `json:"DepotId"` // 809883
	MergeID int64  `json:"MergeId"` // 12
}

type ModifyCloseMRResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 5f988d01-e86c-53c6-1170-c2b67eb9d335
	} `json:"Response"`
}

// ModifyCloseMR 关闭合并请求
func (m *MergeReq) ModifyCloseMR(req *ModifyCloseMRReq) (resp ModifyCloseMRResp, err error) {
	err = poster.Post(m.c, "ModifyCloseMR", req, &resp)
	return
}
