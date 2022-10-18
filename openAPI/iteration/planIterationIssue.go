package iteration

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type PlanIterationIssueReq struct {
	Action        string  `json:"Action"`        // PlanIterationIssue
	IssueCode     []int64 `json:"IssueCode"`     // 1
	IterationCode int64   `json:"IterationCode"` // 1
	ProjectName   string  `json:"ProjectName"`   // demo-project
}

type PlanIterationIssueResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // ae8e2d5f-569b-443e-8c61-440ea3a7562a
	} `json:"Response"`
}

// PlanIterationIssue 批量规划迭代
func (i *Iteration) PlanIterationIssue(req *PlanIterationIssueReq) (resp PlanIterationIssueResp, err error) {
	err = poster.Post(i.c, "PlanIterationIssue", req, &resp)
	return
}
