package wiki

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeImportJobStatusReq struct {
	Action        string `json:"Action"`        // DescribeImportJobStatus
	Authorization string `json:"Authorization"` // 8c044154bd5c7cadb5720944c733b018d0984b3
	JobID         string `json:"JobId"`         // asdaqweqsdaddsdsadsa
	ProjectName   string `json:"ProjectName"`   // demo-project
}

type DescribeImportJobStatusResp struct {
	Response struct {
		JobID     string `json:"JobId"`            // 5b3848e2dfa921e183a071df84aa
		RequestID int64  `json:"RequestId,string"` // 1
		Status    string `json:"Status"`           // WAIT_PROCESS
	} `json:"Response"`
}

// DescribeImportJobStatus 任务状态查询
func (w *Wiki) DescribeImportJobStatus(req *DescribeImportJobStatusReq) (resp DescribeImportJobStatusResp, err error) {
	err = poster.Post(w.c, "DescribeImportJobStatus", req, &resp)
	return
}
