package issue

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeBlockedByIssueListReq struct {
	IssueCode   int64  `json:"IssueCode"`   // 1
	ProjectName string `json:"ProjectName"` // project-demo
}

type DescribeBlockedByIssueListResp struct {
	Response struct {
		Issues []struct {
			Assignee struct {
				Avatar        string `json:"Avatar"`        // https://dn-coding-net-production-static.codehub.cn/512b2a62-956b-4ef8-8e84-b3c66e71468f.png?imageMogr2/auto-orient/format/png/crop/!300x300a0a0
				Email         string `json:"Email"`         // coding.com
				GlobalKey     string `json:"GlobalKey"`     // coding
				ID            int64  `json:"Id"`            // 1
				Name          string `json:"Name"`          // xxx
				Phone         int64  `json:"Phone,string"`  // 111
				Status        int64  `json:"Status"`        // 1
				TeamGlobalKey string `json:"TeamGlobalKey"` // coding
				TeamID        int64  `json:"TeamId"`        // 1
			} `json:"Assignee"`
			Code            int64  `json:"Code"`            // 27
			IssueStatusID   int64  `json:"IssueStatusId"`   // 4
			IssueStatusName string `json:"IssueStatusName"` // 未开始
			IssueStatusType string `json:"IssueStatusType"` // TODO
			Name            string `json:"Name"`            // xxx
			Priority        int64  `json:"Priority,string"` // 3
			Type            string `json:"Type"`            // REQUIREMENT
		} `json:"Issues"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeBlockedByIssueList 查询前置事项
func (i *Issue) DescribeBlockedByIssueList(req *DescribeBlockedByIssueListReq) (resp DescribeBlockedByIssueListResp, err error) {
	err = poster.Post(i.c, "DescribeBlockedByIssueList", req, &resp)
	return
}
