package work

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteIssueWorkHoursReq struct {
	IssueCode              int64  `json:"IssueCode"`              // 1
	ProjectName            string `json:"ProjectName"`            // project-demo
	RollbackRemainingHours bool   `json:"RollbackRemainingHours"` // true
	WorkHourLogID          int64  `json:"WorkHourLogId"`          // 2
}

type DeleteIssueWorkHoursResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // xx
	} `json:"Response"`
}

// DeleteIssueWorkHours 删除工时日志
func (w *Work) DeleteIssueWorkHours(req *DeleteIssueWorkHoursReq) (resp DeleteIssueWorkHoursResp, err error) {
	err = poster.Post(w.c, "DeleteIssueWorkHours", req, &resp)
	return
}
