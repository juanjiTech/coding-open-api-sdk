package team

import "github.com/guonaihong/gout"

type DescribeCodingProjectsReq struct {
	Action      string `json:"Action"`
	ProjectName string `json:"ProjectName,omitempty"` // Optional
	PageNumber  int    `json:"PageNumber"`
	PageSize    int    `json:"PageSize"`
}

type DescribeCodingProjectsResp struct {
	Response struct {
		RequestId string `json:"RequestId"`
		Data      struct {
			PageNumber  int `json:"PageNumber"`
			PageSize    int `json:"PageSize"`
			TotalCount  int `json:"TotalCount"`
			ProjectList []struct {
				Id          int    `json:"Id"`
				CreatedAt   int64  `json:"CreatedAt"`
				UpdatedAt   int64  `json:"UpdatedAt"`
				Status      int    `json:"Status"`
				Type        int    `json:"Type"`
				MaxMember   int    `json:"MaxMember"`
				Name        string `json:"Name"`
				DisplayName string `json:"DisplayName"`
				Description string `json:"Description"`
				Icon        string `json:"Icon"`
				TeamOwnerId int    `json:"TeamOwnerId"`
				UserOwnerId int    `json:"UserOwnerId"`
				StartDate   int    `json:"StartDate"`
				EndDate     int    `json:"EndDate"`
				TeamId      int    `json:"TeamId"`
				IsDemo      bool   `json:"IsDemo"`
				Archived    bool   `json:"Archived"`
			} `json:"ProjectList"`
		} `json:"Data"`
	} `json:"Response"`
}

// DescribeCodingProjects 查询团队内所有项目列表
func (Team) DescribeCodingProjects(PageNumber, PageSize int, ProjectName ...string) (resp DescribeCodingProjectsResp, err error) {
	req := DescribeCodingProjectsReq{
		Action:     "DescribeCodingProjects",
		PageNumber: PageNumber,
		PageSize:   PageSize,
	}
	if len(ProjectName) > 0 {
		req.ProjectName = ProjectName[0]
	}
	err = gout.POST("https://e.coding.net/open-api?Action=DescribeCodingProjects").
		SetJSON(req).
		BindJSON(&resp).
		Do()
	return
}
