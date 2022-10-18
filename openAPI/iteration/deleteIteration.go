package iteration

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteIterationReq struct {
	Action        string `json:"Action"`        // DeleteIteration
	IterationCode int64  `json:"IterationCode"` // 1
	ProjectName   string `json:"ProjectName"`   // demo-project
}

type DeleteIterationResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // ae8e2d5f-569b-443e-8c61-440ea3a7562a
	} `json:"Response"`
}

// DeleteIteration 删除迭代
func (i *Iteration) DeleteIteration(req *DeleteIterationReq) (resp DeleteIterationResp, err error) {
	err = poster.Post(i.c, "DeleteIteration", req, &resp)
	return
}
