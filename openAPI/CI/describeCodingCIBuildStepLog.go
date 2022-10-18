package ci

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeCodingCIBuildStepLogReq struct {
	Action  string `json:"Action"`  // DescribeCodingCIBuildStepLog
	BuildID int64  `json:"BuildId"` // 56
	StageID int64  `json:"StageId"` // 24
	Start   int64  `json:"Start"`   // 0
	StepID  int64  `json:"StepId"`  // 25
}

type DescribeCodingCIBuildStepLogResp struct {
	Response struct {
		Data struct {
			Log string `json:"Log"` // 单元测试中...

			MoreData      bool  `json:"MoreData"`      // false
			TextDelivered int64 `json:"TextDelivered"` // 10
			TextSize      int64 `json:"TextSize"`      // 10
		} `json:"Data"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeCodingCIBuildStepLog 获取构建步骤日志
func (c *Ci) DescribeCodingCIBuildStepLog(req *DescribeCodingCIBuildStepLogReq) (resp DescribeCodingCIBuildStepLogResp, err error) {
	err = poster.Post(c.c, "DescribeCodingCIBuildStepLog", req, &resp)
	return
}
