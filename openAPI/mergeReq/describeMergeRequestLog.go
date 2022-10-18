package mergeReq

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeMergeRequestLogReq struct {
	Action  string `json:"Action"`  // DescribeMergeRequestLog
	DepotID int64  `json:"DepotId"` // 46
	MergeID int64  `json:"MergeId"` // 5
}

type DescribeMergeRequestLogResp struct {
	Response struct {
		Logs []struct {
			Action string `json:"Action"` // create
			ID     int64  `json:"Id"`     // 1280
			Name   string `json:"Name"`   // 洋葱猴1
		} `json:"Logs"`
		RequestID string `json:"RequestId"` // e9fb2397-f4c5-e3f7-4f10-5fe8aec3c6ae
	} `json:"Response"`
}

// DescribeMergeRequestLog 获取合并请求操作记录
func (m *MergeReq) DescribeMergeRequestLog(req *DescribeMergeRequestLogReq) (resp DescribeMergeRequestLogResp, err error) {
	err = poster.Post(m.c, "DescribeMergeRequestLog", req, &resp)
	return
}
