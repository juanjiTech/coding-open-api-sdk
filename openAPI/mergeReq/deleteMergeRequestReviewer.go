package mergeReq

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteMergeRequestReviewerReq struct {
	Action            string `json:"Action"`            // DeleteMergeRequestReviewer
	DepotID           int64  `json:"DepotId"`           // 1
	MergeID           int64  `json:"MergeId"`           // 1
	ReviewerGlobalKey string `json:"ReviewerGlobalKey"` // coding-coding
}

type DeleteMergeRequestReviewerResp struct {
	Response struct {
		RequestID string        `json:"RequestId"` // 29ec9529-e03a-4a89-b71d-d988676e608e
		Reviewers []interface{} `json:"Reviewers"` // <nil>
	} `json:"Response"`
}

// DeleteMergeRequestReviewer 删除合并请求评论者
func (m *MergeReq) DeleteMergeRequestReviewer(req *DeleteMergeRequestReviewerReq) (resp DeleteMergeRequestReviewerResp, err error) {
	err = poster.Post(m.c, "DeleteMergeRequestReviewer", req, &resp)
	return
}
