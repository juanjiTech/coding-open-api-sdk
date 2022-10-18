package cd

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeCdDeployCountByProjectReq struct {
	Action      string `json:"Action"`      // DescribeCdDeployCountByProject
	EndAt       string `json:"EndAt"`       // 2020-12-03 23:59:59
	ProjectName string `json:"ProjectName"` // test
	StartAt     string `json:"StartAt"`     // 2020-12-01 00:00:00
}

type DescribeCdDeployCountByProjectResp struct {
	Response struct {
		Data struct {
			Details []struct {
				Application   string `json:"Application"` // dev
				CdDeployCount struct {
					SucceedCount int64 `json:"SucceedCount"` // 2
					TotalCount   int64 `json:"TotalCount"`   // 2
				} `json:"CdDeployCount"`
			} `json:"Details"`
			EndDate   string `json:"EndDate"`   // 2020-12-03
			StartDate string `json:"StartDate"` // 2020-12-01
			Total     struct {
				SucceedCount int64 `json:"SucceedCount"` // 3
				TotalCount   int64 `json:"TotalCount"`   // 3
			} `json:"Total"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 4cb20019-03de-3ae1-08a6-7b2e0817ca26
	} `json:"Response"`
}

// DescribeCdDeployCountByProject 根据项目名获取关联应用的发布次数
func (c *Cd) DescribeCdDeployCountByProject(req *DescribeCdDeployCountByProjectReq) (resp DescribeCdDeployCountByProjectResp, err error) {
	err = poster.Post(c.c, "DescribeCdDeployCountByProject", req, &resp)
	return
}
