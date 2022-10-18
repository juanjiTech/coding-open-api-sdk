package test

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateReportReq struct {
	Action      string  `json:"Action"`      // CreateReport
	Name        string  `json:"Name"`        // run demo
	ProjectName string  `json:"ProjectName"` // project name
	RunIDs      []int64 `json:"RunIds"`      // 1
}

type CreateReportResp struct {
	Response struct {
		Data struct {
			Report struct {
				CreatedAt           string      `json:"CreatedAt"`           // 2021-06-28 15:33:42
				ID                  int64       `json:"Id"`                  // 1750
				Name                string      `json:"Name"`                // 报告2
				StatisticsEndTime   string      `json:"StatisticsEndTime"`   // 2021-06-28 15:33:42
				StatisticsStartTime string      `json:"StatisticsStartTime"` // 2021-06-16 10:44:26
				Status              string      `json:"Status"`              // UNAVAILABLE
				Summary             interface{} `json:"Summary"`             // <nil>
			} `json:"Report"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 980cbd88-caac-626a-8c8d-8019acf41bb4
	} `json:"Response"`
}

// CreateReport 创建测试报告
func (t *Test) CreateReport(req *CreateReportReq) (resp CreateReportResp, err error) {
	err = poster.Post(t.c, "CreateReport", req, &resp)
	return
}
