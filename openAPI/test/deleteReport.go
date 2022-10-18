package test

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteReportReq struct {
	Action      string `json:"Action"`      // DeleteReport
	ProjectName string `json:"ProjectName"` // project name
	ReportID    int64  `json:"ReportId"`    // 1
}

type DeleteReportResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 980cbd88-caac-626a-8c8d-8019acf41bb4
	} `json:"Response"`
}

// DeleteReport 删除测试报告
func (t *Test) DeleteReport(req *DeleteReportReq) (resp DeleteReportResp, err error) {
	err = poster.Post(t.c, "DeleteReport", req, &resp)
	return
}
