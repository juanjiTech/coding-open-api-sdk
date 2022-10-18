package test

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeTestCaseSectionListReq struct {
	Action      string `json:"Action"`      // DescribeTestCaseSectionList
	ParentID    int64  `json:"ParentId"`    // 1
	ProjectName string `json:"ProjectName"` // 项目名称
}

type DescribeTestCaseSectionListResp struct {
	Response struct {
		Data struct {
			Sections []struct {
				ID       int64 `json:"Id"`          // 1
				Name     int64 `json:"Name,string"` // 1
				ParentID int64 `json:"ParentId"`    // 0
				Sort     int64 `json:"Sort"`        // 1
			} `json:"Sections"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 62741944-506c-3c11-e77b-47f23fb79086
	} `json:"Response"`
}

// DescribeTestCaseSectionList 测试用例分组列表
func (t *Test) DescribeTestCaseSectionList(req *DescribeTestCaseSectionListReq) (resp DescribeTestCaseSectionListResp, err error) {
	err = poster.Post(t.c, "DescribeTestCaseSectionList", req, &resp)
	return
}
